package character

import (
	"gopkg.in/telebot.v3"
)

func Material(ctx telebot.Context) error {
	return Handler(ctx, "material")
}
