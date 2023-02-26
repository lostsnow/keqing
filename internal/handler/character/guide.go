package character

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/litsea/logger"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/character"
)

var (
	ErrDownloadFileNotFound = errors.New("download file not found")
)

func Guide(ctx telebot.Context) error {
	if len(ctx.Args()) == 0 {
		return nil
	}

	name := strings.Join(ctx.Args(), " ")
	chars := character.Search(name)
	if len(chars) == 0 {
		return handler.Send(ctx, "No such character")
	}

	fileDir := fmt.Sprintf("assets/character/guide/%s", chars[0].Elemental.Id)
	fileName := fmt.Sprintf("%s.png", chars[0].Id)
	cacheFileDir := fmt.Sprintf("%s/%s", viper.GetString("cache.dir"), fileDir)
	cacheFilePath := fmt.Sprintf("%s/%s", cacheFileDir, fileName)
	fileIdPath := fmt.Sprintf("%s/%s.id", cacheFileDir, fileName)

	if _, err := os.Stat(fileIdPath); err == nil {
		fb, err := os.ReadFile(fileIdPath)
		if err == nil && len(fb) > 0 {
			return ctx.Send(&telebot.Photo{File: telebot.File{FileID: string(fb)}})
		}
	}

	f, err := os.Open(cacheFilePath)
	if err == nil {
		return ctx.Send(&telebot.Photo{File: telebot.FromReader(f)})
	}

	url := fmt.Sprintf("%s/%s/%s", viper.GetString("app.assets.base-url"), fileDir, fileName)

	err = downloadFile(cacheFileDir, fileName, url)
	if err != nil {
		if err != ErrDownloadFileNotFound {
			logger.Error(err)
		}
		return handler.Send(ctx, "No such character guide")
	}

	p := &telebot.Photo{File: telebot.FromDisk(cacheFilePath)}
	err = ctx.Send(p)
	if err != nil {
		return err
	}

	if p.FileID == "" {
		return nil
	}
	if _, err = os.Stat(fileIdPath); os.IsNotExist(err) {
		_ = os.WriteFile(fileIdPath, []byte(p.FileID), 0644)
	}
	return nil
}

func downloadFile(fileDir, fileName string, url string) error {
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.MkdirAll(fileDir, os.ModeDir); err != nil {
			return fmt.Errorf("create download dir failed: %s", err)
		}
	}

	// Create the file
	filePath := fmt.Sprintf("%s/%s", fileDir, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create download file %s failed: %s", filePath, err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download file %s failed: %s", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return ErrDownloadFileNotFound
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("write download file %s failed: %s", filePath, err)
	}

	return nil
}
