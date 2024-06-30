package character

import (
	"gopkg.in/telebot.v3"
)

func Curve(ctx telebot.Context) error {
	return Handler(ctx, "curve")
}
