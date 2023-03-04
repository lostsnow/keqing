package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
	"github.com/lostsnow/keqing/pkg/object"
)

var (
	ErrDownloadPhotoNotFound = errors.New("download photo not found")
)

var (
	currentInlineKeywordMark = "✅"
)

type PhotoResponseHandler struct {
	Buttons        []PhotoButton
	NoPhotoMessage any
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
				title = title + currentInlineKeywordMark
			}
			botBtn := sel.Data(title, toPhotoUniqueId(btn.Dir+"/"+btn.Name))
			botBtns = append(botBtns, botBtn)
		}
		chunks := object.ChunkBy(botBtns, 3)
		rows := make([]telebot.Row, 0, len(chunks))
		for _, chunk := range chunks {
			rows = append(rows, chunk)
		}
		sel.Inline(rows...)

		for _, bb := range botBtns {
			ctx.Bot().Handle(&bb, func(c telebot.Context) error {
				rFilePath := fromPhotoUniqueId(c.Callback().Unique)
				mt, m := h.Get(c, rFilePath)
				if _, ok := m.(telebot.Inputtable); !ok {
					return c.Respond()
				}
				err := updateCurrentInlineKeyboard(sel, c.Callback().Unique)
				if err != nil {
					return c.Respond()
				}

				_, err = c.Bot().EditMedia(c.Message(), m.(telebot.Inputtable), sel)
				if err != nil {
					return fmt.Errorf("edit photo %s/%s failed: %s", h.Buttons[0].Dir, h.Buttons[0].Name, err)
				}

				cacheFileIdPath := getCacheFileIdPath(rFilePath)
				cachePhotoId(m, mt, cacheFileIdPath)
				return nil
			})
		}
	}

	mt, m := h.GetByButton(ctx, h.Buttons[0])
	err := ctx.Reply(m, sel)
	if err != nil {
		return fmt.Errorf("send photo %s/%s failed: %s", h.Buttons[0].Dir, h.Buttons[0].Name, err)
	}

	cacheFileIdPath := getCacheFileIdPath(fmt.Sprintf("%s/%s", h.Buttons[0].Dir, h.Buttons[0].Name))
	cachePhotoId(m, mt, cacheFileIdPath)
	return nil
}

func (h PhotoResponseHandler) GetByButton(ctx telebot.Context, fb PhotoButton) (MessageType, any) {
	return h.Get(ctx, fmt.Sprintf("%s/%s", fb.Dir, fb.Name))
}

func (h PhotoResponseHandler) Get(ctx telebot.Context, relFilePath string) (MessageType, any) {
	cacheFilePath := getCacheFilePath(relFilePath)
	cacheFileIdPath := getCacheFileIdPath(relFilePath)

	if _, err := os.Stat(cacheFileIdPath); err == nil {
		fb, err := os.ReadFile(cacheFileIdPath)
		if err == nil && len(fb) > 0 {
			return MessagePhotoId, &telebot.Photo{File: telebot.File{FileID: string(fb)}}
		}
	}

	f, err := os.Open(cacheFilePath)
	if err == nil {
		return MessageCachedPhoto, &telebot.Photo{File: telebot.FromReader(f)}
	}

	url := getFileUrl(relFilePath)
	err = downloadPhoto(cacheFilePath, url)
	if err != nil {
		if err != ErrDownloadPhotoNotFound {
			ReportError(ctx, err.Error())
		}
		return MessageEmbed, h.NoPhotoMessage
	}

	p := &telebot.Photo{File: telebot.FromDisk(cacheFilePath)}
	return MessagePhoto, p
}

func downloadPhoto(filePath string, url string) error {
	fileDir := filepath.Dir(filePath)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.MkdirAll(fileDir, os.ModeDir); err != nil {
			return fmt.Errorf("create download photo dir failed: %s", err)
		}
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download photo %s failed: %s", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ErrDownloadPhotoNotFound
	}

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		_ = os.Remove(filePath)
		return fmt.Errorf("create download photo %s failed: %s", filePath, err)
	}
	defer out.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		_ = os.Remove(filePath)
		return fmt.Errorf("write download photo %s failed: %s", filePath, err)
	}

	return nil
}

func getCacheFilePath(relFilePath string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("cache.dir"), relFilePath)
}

func getCacheFileIdPath(relFilePath string) string {
	cacheFilePath := getCacheFilePath(relFilePath)
	return fmt.Sprintf("%s.id", cacheFilePath)
}

func getFileUrl(relFilePath string) string {
	return fmt.Sprintf("%s/%s", viper.GetString("app.assets.base-url"), relFilePath)
}

func cachePhotoId(msg any, mt MessageType, fileIdPath string) {
	if _, ok := msg.(*telebot.Photo); !ok {
		return
	}
	if mt != MessagePhoto {
		return
	}

	fileId := msg.(*telebot.Photo).FileID
	if fileId == "" {
		return
	}
	if _, err := os.Stat(fileIdPath); os.IsNotExist(err) {
		_ = os.WriteFile(fileIdPath, []byte(fileId), 0644)
	}
	return
}

func toPhotoUniqueId(path string) string {
	replacer := strings.NewReplacer(
		"/", "--",
		".", "__",
		" ", "00",
	)
	return replacer.Replace(path)
}

func fromPhotoUniqueId(path string) string {
	replacer := strings.NewReplacer(
		"--", "/",
		"__", ".",
		"00", " ",
	)
	return replacer.Replace(path)
}

func updateCurrentInlineKeyboard(sel *telebot.ReplyMarkup, uniq string) error {
	if sel == nil {
		return errors.New("no keyboard")
	}
	for i, row := range sel.InlineKeyboard {
		for j, col := range row {
			if col.Unique != uniq && strings.HasSuffix(col.Text, currentInlineKeywordMark) {
				sel.InlineKeyboard[i][j].Text = strings.ReplaceAll(sel.InlineKeyboard[i][j].Text, currentInlineKeywordMark, "")
			} else if col.Unique == uniq {
				if !strings.HasSuffix(col.Text, currentInlineKeywordMark) {
					sel.InlineKeyboard[i][j].Text = col.Text + currentInlineKeywordMark
				} else {
					return errors.New("keyboard has already active")
				}
			}
		}
	}
	return nil
}
