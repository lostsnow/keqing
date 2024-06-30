package character

import (
	"gopkg.in/telebot.v3"
)

func Guide(ctx telebot.Context) error {
	return Handler(ctx, "guide")
}
