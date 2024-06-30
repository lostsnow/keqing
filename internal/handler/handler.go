package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/litsea/logger"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/i18n"
)

type MessageType int

var (
	ErrNoKeyboard               = errors.New("no keyboard")
	ErrKeyboardHasAlreadyActive = errors.New("keyboard has already active")
)

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
	err := ctx.Send(i18n.T(ctx, message))
	if err != nil {
		return fmt.Errorf("handler.Send: %w", err)
	}

	return nil
}

func Reply(ctx telebot.Context, message string) error {
	err := ctx.Reply(i18n.T(ctx, message))
	if err != nil {
		return fmt.Errorf("handler.Reply: %w", err)
	}

	return nil
}

func SendPhoto(ctx telebot.Context, message *telebot.Photo) error {
	err := ctx.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("handler.SendPhoto: %w", err)
	}

	return nil
}

func ReplyPhoto(ctx telebot.Context, message *telebot.Photo) error {
	err := ctx.Reply(ctx, message)
	if err != nil {
		return fmt.Errorf("handler.ReplyPhoto: %w", err)
	}

	return nil
}

func Respond(ctx telebot.Context, resp ...*telebot.CallbackResponse) error {
	err := ctx.Respond(resp...)
	if err != nil {
		return fmt.Errorf("handler.Respond: %w", err)
	}

	return nil
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
		return ErrNoKeyboard
	}

	for i, row := range sel.InlineKeyboard {
		for j, col := range row {
			if col.Unique != uniq && strings.HasSuffix(col.Text, CurrentInlineKeywordMark) {
				sel.InlineKeyboard[i][j].Text = strings.ReplaceAll(sel.InlineKeyboard[i][j].Text, CurrentInlineKeywordMark, "")
			} else if col.Unique == uniq {
				if !strings.HasSuffix(col.Text, CurrentInlineKeywordMark) {
					sel.InlineKeyboard[i][j].Text = col.Text + CurrentInlineKeywordMark
				} else {
					return ErrKeyboardHasAlreadyActive
				}
			}
		}
	}

	return nil
}
