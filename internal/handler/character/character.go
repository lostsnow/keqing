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

func Handler(ctx telebot.Context, typ string) error {
	if len(ctx.Args()) == 0 {
		return nil
	}

	name := strings.Join(ctx.Args(), " ")
	chars := character.Search(name)

	maxLen := len(chars)
	if maxLen == 0 {
		return ctx.Reply(handler.UnknownPhoto)
	}

	if maxLen > 9 {
		maxLen = 9
	}

	buttons := make([]handler.PhotoButton, 0, maxLen)

	for idx, char := range chars {
		if idx == maxLen {
			break
		}

		buttons = append(buttons, handler.PhotoButton{
			Title: i18n.T(ctx, char.ID),
			Dir:   fmt.Sprintf("assets/character/%s/%s", typ, char.Elemental.ID),
			Name:  fmt.Sprintf("%s.png", char.ID),
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	return h.Handle(ctx)
}
