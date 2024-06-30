package material

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
)

func Weapon(ctx telebot.Context) error {
	h := handler.PhotoResponseHandler{
		Buttons: []handler.PhotoButton{
			{
				Dir:  "assets/material/weapon",
				Name: "weapon.png",
			},
		},
		NoPhotoMessage: i18n.T(ctx, "What are you looking for, it does not exist"),
	}

	err := h.Handle(ctx)
	if err != nil {
		return fmt.Errorf("material.Weapon: %w", err)
	}

	return nil
}
