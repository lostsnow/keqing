package character

import (
	"gopkg.in/telebot.v3"

	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

func Material(ctx telebot.Context) error {
	return Handler(ctx, "material")
}
