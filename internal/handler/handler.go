package handler

import (
	"fmt"

	"github.com/litsea/logger"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

type MessageType int

const (
	MessageText MessageType = iota + 1
	MessageEmbed
	MessagePhoto
	MessageCachedPhoto
	MessagePhotoId
)

type PhotoButton struct {
	Title       string
	Dir         string
	Name        string
	Message     any
	MessageType MessageType
}

func Help(ctx telebot.Context) error {
	return Send(ctx, "Cut to the chase")
}

func Trace(ctx telebot.Context) error {
	return ctx.Send(fmt.Sprintf("*Bot*: `%d`\n*Chat*: `%d` <%s>\n*From*: `%d`\n*Message*: `%s`",
		ctx.Bot().Me.ID, ctx.Chat().ID, ctx.Chat().Type, ctx.Sender().ID, ctx.Text()),
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
}

func Send(ctx telebot.Context, message string) error {
	return ctx.Send(i18n.T(ctx, message))
}

func Reply(ctx telebot.Context, message string) error {
	return ctx.Reply(i18n.T(ctx, message))
}

func ReportError(ctx telebot.Context, msg string) {
	logger.Error(msg)
	if handlerBot == nil {
		return
	}
	for _, u := range handlerBot.AdminUsers {
		_, _ = ctx.Bot().Send(&u, msg, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	}
}
