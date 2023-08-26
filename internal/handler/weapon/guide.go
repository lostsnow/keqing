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
	if len(weapons) == 0 {
		return ctx.Reply(handler.UnknownPhoto)
	}
	max := len(weapons)
	if max > 9 {
		max = 9
	}

	buttons := make([]handler.PhotoButton, 0, max)
	for idx, weapon := range weapons {
		if idx == max {
			break
		}
		buttons = append(buttons, handler.PhotoButton{
			Title: i18n.T(ctx, weapon.ID),
			Dir:   fmt.Sprintf("assets/weapon/guide/%s", weapon.Type.ID),
			Name:  fmt.Sprintf("%s.png", weapon.ID),
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	return h.Handle(ctx)
}
