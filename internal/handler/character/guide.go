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
		return handler.Send(ctx, "No such character")
	}

	h := handler.PhotoResponseHandler{
		FileDir:        fmt.Sprintf("assets/character/guide/%s", chars[0].Elemental.Id),
		FileName:       fmt.Sprintf("%s.png", chars[0].Id),
		NoPhotoMessage: i18n.T(ctx, "No such character guide"),
	}

	return h.Send(ctx)
}
