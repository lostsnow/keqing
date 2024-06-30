package game

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
)

func RefreshRole(ctx telebot.Context) error {
	err := handler.Reply(ctx, i18n.T(ctx, "Not yet implemented, send command %s instead", "/auth_qrcode"))
	if err != nil {
		return fmt.Errorf("game.RefreshRole: %w", err)
	}

	return nil
}
