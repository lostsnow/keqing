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

func Send(ctx telebot.Context, message string) error {
	return ctx.Send(i18n.T(ctx, message))
}

func Reply(ctx telebot.Context, message string) error {
	return ctx.Reply(i18n.T(ctx, message))
}

func ReportError(ctx telebot.Context, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logger.Error(msg)
	if handlerBot == nil {
		return
	}
	for _, u := range handlerBot.AdminUsers {
		_, _ = ctx.Bot().Send(&u, msg, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	}
}
