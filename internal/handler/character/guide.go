package character

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/character"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

func Guide(ctx telebot.Context) error {
	if len(ctx.Args()) == 0 {
		return nil
	}

	name := strings.Join(ctx.Args(), " ")
	chars := character.Search(name)
	if len(chars) == 0 {
		return ctx.Reply(handler.UnknownPhoto)
	}
	max := len(chars)
	if max > 6 {
		max = 6
	}

	buttons := make([]handler.PhotoButton, 0, max)
	for idx, char := range chars {
		if idx == max {
			break
		}
		buttons = append(buttons, handler.PhotoButton{
			Title: i18n.T(ctx, char.Id),
			Dir:   fmt.Sprintf("assets/character/guide/%s", char.Elemental.Id),
			Name:  fmt.Sprintf("%s.png", char.Id),
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	return h.Handle(ctx)
}
