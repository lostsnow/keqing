package material

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
)

func Character(ctx telebot.Context) error {
	h := handler.PhotoResponseHandler{
		Buttons: []handler.PhotoButton{
			{
				Dir:  "assets/material/character",
				Name: "character.png",
			},
		},
		NoPhotoMessage: i18n.T(ctx, "What are you looking for, it does not exist"),
	}

	err := h.Handle(ctx)
	if err != nil {
		return fmt.Errorf("material.Character: %w", err)
	}

	return nil
}
