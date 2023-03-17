package character

import (
	"gopkg.in/telebot.v3"

	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

func Graduation(ctx telebot.Context) error {
	return Handler(ctx, "graduation")
}
