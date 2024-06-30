package weapon

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/i18n"
	"github.com/lostsnow/keqing/pkg/weapon"
)

func Guide(ctx telebot.Context) error {
	if len(ctx.Args()) == 0 {
		return nil
	}

	name := strings.Join(ctx.Args(), " ")
	weapons := weapon.Search(name)

	mx := len(weapons)
	if mx == 0 {
		err := handler.ReplyPhoto(ctx, handler.UnknownPhoto)
		if err != nil {
			return fmt.Errorf("weapon.Guide: %w", err)
		}
	}

	if mx > 9 {
		mx = 9
	}

	buttons := make([]handler.PhotoButton, 0, mx)

	for idx, w := range weapons {
		if idx == mx {
			break
		}

		buttons = append(buttons, handler.PhotoButton{
			Title: i18n.T(ctx, w.ID),
			Dir:   "assets/weapon/guide/" + w.Type.ID,
			Name:  w.ID + ".png",
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	err := h.Handle(ctx)
	if err != nil {
		return fmt.Errorf("weapon.Guide: %w", err)
	}

	return nil
}
