package character

import (
	"gopkg.in/telebot.v3"
)

func Graduation(ctx telebot.Context) error {
	return Handler(ctx, "graduation")
}
