package handler

import (
	"errors"
	"fmt"
	"strings"

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
	MessagePhotoID
)

var CurrentInlineKeywordMark = "✅"

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

	for idx := range handlerBot.AdminUsers {
		_, _ = ctx.Bot().Send(&handlerBot.AdminUsers[idx], msg, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	}
}

func UpdateCurrentInlineKeyboard(sel *telebot.ReplyMarkup, uniq string) error {
	if sel == nil {
		return errors.New("no keyboard")
	}

	for i, row := range sel.InlineKeyboard {
		for j, col := range row {
			if col.Unique != uniq && strings.HasSuffix(col.Text, CurrentInlineKeywordMark) {
				sel.InlineKeyboard[i][j].Text = strings.ReplaceAll(sel.InlineKeyboard[i][j].Text, CurrentInlineKeywordMark, "")
			} else if col.Unique == uniq {
				if !strings.HasSuffix(col.Text, CurrentInlineKeywordMark) {
					sel.InlineKeyboard[i][j].Text = col.Text + CurrentInlineKeywordMark
				} else {
					return errors.New("keyboard has already active")
				}
			}
		}
	}

	return nil
}
