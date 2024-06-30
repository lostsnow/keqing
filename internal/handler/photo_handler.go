package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/data"
	"github.com/lostsnow/keqing/pkg/object"
)

var ErrDownloadPhotoNotFound = errors.New("download photo not found")

type PhotoResponseHandler struct {
	Buttons        []PhotoButton
	NoPhotoMessage any
}

var (
	UnknownPhoto *telebot.Photo
	NoPhoto      *telebot.Photo
)

//nolint:gochecknoinits
func init() {
	f, _ := data.Embed.Open("embed/400.png")
	UnknownPhoto = &telebot.Photo{File: telebot.File{FileReader: f}}
	f, _ = data.Embed.Open("embed/404.png")
	NoPhoto = &telebot.Photo{File: telebot.File{FileReader: f}}
}

func (h PhotoResponseHandler) Handle(ctx telebot.Context) error {
	if len(h.Buttons) == 0 {
		return nil
	}

	sel := &telebot.ReplyMarkup{}

	if len(h.Buttons) > 1 {
		botBtns := make([]telebot.Btn, 0, len(h.Buttons))

		for idx, btn := range h.Buttons {
			title := btn.Title
			if idx == 0 {
				title += CurrentInlineKeywordMark
			}

			botBtn := sel.Data(title, toPhotoUniqueID(btn.Dir+"/"+btn.Name))
			botBtns = append(botBtns, botBtn)
		}

		chunks := object.ChunkBy(botBtns, 3)
		rows := make([]telebot.Row, 0, len(chunks))

		for _, chunk := range chunks {
			rows = append(rows, chunk)
		}

		sel.Inline(rows...)

		for idx := range botBtns {
			ctx.Bot().Handle(&botBtns[idx], func(c telebot.Context) error {
				rFilePath := fromPhotoUniqueID(c.Callback().Unique)

				mt, m := h.Get(c, rFilePath)
				if _, ok := m.(telebot.Inputtable); !ok {
					return c.Respond()
				}

				err := UpdateCurrentInlineKeyboard(sel, c.Callback().Unique)
				if err != nil {
					return c.Respond()
				}

				_, err = c.Bot().EditMedia(c.Message(), m.(telebot.Inputtable), sel)
				if err != nil {
					return fmt.Errorf("edit photo %s/%s failed: %w", h.Buttons[0].Dir, h.Buttons[0].Name, err)
				}

				cacheFileIDPath := getCacheFileIDPath(rFilePath)
				cachePhotoID(m, mt, cacheFileIDPath)

				return nil
			})
		}
	}

	mt, m := h.GetByButton(ctx, h.Buttons[0])

	err := ctx.Reply(m, sel)
	if err != nil {
		return fmt.Errorf("send photo %s/%s failed: %w", h.Buttons[0].Dir, h.Buttons[0].Name, err)
	}

	cacheFileIDPath := getCacheFileIDPath(fmt.Sprintf("%s/%s", h.Buttons[0].Dir, h.Buttons[0].Name))
	cachePhotoID(m, mt, cacheFileIDPath)

	return nil
}

func (h PhotoResponseHandler) GetByButton(ctx telebot.Context, fb PhotoButton) (MessageType, any) {
	return h.Get(ctx, fmt.Sprintf("%s/%s", fb.Dir, fb.Name))
}

func (h PhotoResponseHandler) Get(ctx telebot.Context, relFilePath string) (MessageType, any) {
	cacheFilePath := getCacheFilePath(relFilePath)
	cacheFileIDPath := getCacheFileIDPath(relFilePath)
	cacheExpired := false

	fi, err := os.Stat(cacheFileIDPath)
	if err == nil && fi.ModTime().Before(time.Now().AddDate(0, -1, 0)) {
		cacheExpired = true
	}

	if err == nil && !cacheExpired {
		fb, err := os.ReadFile(cacheFileIDPath)
		if err == nil && len(fb) > 0 {
			return MessagePhotoID, &telebot.Photo{File: telebot.File{FileID: string(fb)}}
		}
	}

	f, err := os.Open(cacheFilePath)
	if err == nil && !cacheExpired {
		return MessageCachedPhoto, &telebot.Photo{File: telebot.FromReader(f)}
	}

	fileURL := getFileURL(relFilePath)

	err = downloadPhoto(cacheFilePath, fileURL)
	if err != nil {
		if !errors.Is(err, ErrDownloadPhotoNotFound) {
			ReportError(ctx, err.Error())
		}

		return MessageEmbed, h.NoPhotoMessage
	}

	p := &telebot.Photo{File: telebot.FromDisk(cacheFilePath)}

	return MessagePhoto, p
}

func downloadPhoto(filePath string, rawURL string) error {
	fileDir := filepath.Dir(filePath)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.MkdirAll(fileDir, os.ModeDir); err != nil {
			return fmt.Errorf("create download photo dir failed: %w", err)
		}
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("download photo url %s invalid: %w", rawURL, err)
	}

	// Get the data
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("download photo %s failed: %w", rawURL, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("download photo %s failed: %w", rawURL, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrDownloadPhotoNotFound
	}

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		_ = os.Remove(filePath)

		return fmt.Errorf("create download photo %s failed: %w", filePath, err)
	}
	defer out.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		_ = os.Remove(filePath)

		return fmt.Errorf("write download photo %s failed: %w", filePath, err)
	}

	return nil
}

func getCacheFilePath(relFilePath string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("cache.dir"), relFilePath)
}

func getCacheFileIDPath(relFilePath string) string {
	cacheFilePath := getCacheFilePath(relFilePath)

	return cacheFilePath + ".id"
}

func getFileURL(relFilePath string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("app.assets.base-url"), relFilePath)
}

func cachePhotoID(msg any, mt MessageType, fileIDPath string) {
	if _, ok := msg.(*telebot.Photo); !ok {
		return
	}

	if mt != MessagePhoto {
		return
	}

	tPhoto, ok := msg.(*telebot.Photo)
	if !ok {
		return
	}

	fileID := tPhoto.FileID
	if fileID == "" {
		return
	}

	if _, err := os.Stat(fileIDPath); os.IsNotExist(err) {
		_ = os.WriteFile(fileIDPath, []byte(fileID), 0o600)
	}
}

func toPhotoUniqueID(path string) string {
	replacer := strings.NewReplacer(
		"/", "--",
		".", "__",
		" ", "00",
	)

	return replacer.Replace(path)
}

func fromPhotoUniqueID(path string) string {
	replacer := strings.NewReplacer(
		"--", "/",
		"__", ".",
		"00", " ",
	)

	return replacer.Replace(path)
}
