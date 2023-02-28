package material

import (
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

func Weekly(ctx telebot.Context) error {
	h := handler.PhotoResponseHandler{
		FileDir:        "assets/material/weekly",
		FileName:       "weekly.png",
		NoPhotoMessage: i18n.T(ctx, "What are you looking for, it does not exist"),
	}

	return h.Send(ctx)
}
