package character

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/pkg/character"
	"github.com/lostsnow/keqing/pkg/i18n"
)

func Handler(ctx telebot.Context, typ string) error {
	if len(ctx.Args()) == 0 {
		return nil
	}

	name := strings.Join(ctx.Args(), " ")
	chars := character.Search(name)

	maxLen := len(chars)
	if maxLen == 0 {
		err := handler.ReplyPhoto(ctx, handler.UnknownPhoto)
		if err != nil {
			return fmt.Errorf("character.Handler: %w", err)
		}
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
			Name:  char.ID + ".png",
		})
	}

	h := handler.PhotoResponseHandler{
		Buttons:        buttons,
		NoPhotoMessage: handler.NoPhoto,
	}

	err := h.Handle(ctx)
	if err != nil {
		return fmt.Errorf("character.Handler: %w", err)
	}

	return nil
}
