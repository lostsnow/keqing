package material

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
)

func Weekly(ctx telebot.Context) error {
	h := handler.PhotoResponseHandler{
		Buttons: []handler.PhotoButton{
			{
				Dir:  "assets/material/weekly",
				Name: "weekly.png",
			},
		},
		NoPhotoMessage: i18n.T(ctx, "What are you looking for, it does not exist"),
	}

	err := h.Handle(ctx)
	if err != nil {
		return fmt.Errorf("material.Weekly: %w", err)
	}

	return nil
}
