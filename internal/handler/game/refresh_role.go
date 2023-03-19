package game

import (
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
)

func RefreshRole(ctx telebot.Context) error {
	return ctx.Reply(i18n.T(ctx, "Not yet implemented, send command %s instead", "/auth_qrcode"))
}
