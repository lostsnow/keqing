package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

var (
	ErrDownloadPhotoNotFound = errors.New("download photo not found")
)

type PhotoResponseHandler struct {
	FileDir        string
	FileName       string
	NoPhotoMessage string
}

func (h PhotoResponseHandler) Handle(ctx telebot.Context) error {
	cacheFileDir := fmt.Sprintf("%s/%s", viper.GetString("cache.dir"), h.FileDir)
	cacheFilePath := fmt.Sprintf("%s/%s", cacheFileDir, h.FileName)
	fileIdPath := fmt.Sprintf("%s/%s.id", cacheFileDir, h.FileName)

	if _, err := os.Stat(fileIdPath); err == nil {
		fb, err := os.ReadFile(fileIdPath)
		if err == nil && len(fb) > 0 {
			return ctx.Reply(&telebot.Photo{File: telebot.File{FileID: string(fb)}})
		}
	}

	f, err := os.Open(cacheFilePath)
	if err == nil {
		return ctx.Reply(&telebot.Photo{File: telebot.FromReader(f)})
	}

	url := fmt.Sprintf("%s/%s/%s", viper.GetString("app.assets.base-url"), h.FileDir, h.FileName)

	err = downloadPhoto(cacheFileDir, h.FileName, url)
	if err != nil {
		if err != ErrDownloadPhotoNotFound {
			ReportError(ctx, err.Error())
		}
		return Reply(ctx, h.NoPhotoMessage)
	}

	p := &telebot.Photo{File: telebot.FromDisk(cacheFilePath)}
	err = ctx.Reply(p)
	if err != nil {
		return fmt.Errorf("send photo %s/%s failed: %s", h.FileDir, h.FileName, err)
	}

	if p.FileID == "" {
		return nil
	}
	if _, err = os.Stat(fileIdPath); os.IsNotExist(err) {
		_ = os.WriteFile(fileIdPath, []byte(p.FileID), 0644)
	}
	return nil
}

func downloadPhoto(fileDir, fileName string, url string) error {
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
	filePath := fmt.Sprintf("%s/%s", fileDir, fileName)
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
