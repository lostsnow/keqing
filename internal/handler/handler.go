package handler

import (
	"fmt"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
)

func Help(ctx telebot.Context) error {
	return ctx.Send(i18n.T(ctx, "Cut to the chase"))
}

func Trace(ctx telebot.Context) error {
	return ctx.Send(fmt.Sprintf("*Bot*: `%d`\n*Chat*: `%d` <%s>\n*From*: `%d`\n*Message*: `%s`",
		ctx.Bot().Me.ID, ctx.Chat().ID, ctx.Chat().Type, ctx.Sender().ID, ctx.Text()),
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
}
