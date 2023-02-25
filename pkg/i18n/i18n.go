package i18n

//go:generate gotext -srclang=en-US update -out=catalog/catalog.go -lang=en-US,zh-Hans github.com/lostsnow/keqing

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/telebot.v3"
	"strings"
)

const (
	MessagePrinter = "MessagePrinter"
)

var langMap = map[string]language.Tag{
	"en-us":   language.AmericanEnglish,
	"zh-hans": language.SimplifiedChinese,
}

var printers = make(map[language.Tag]*message.Printer, len(langMap))

func findLanguage(lang string) language.Tag {
	if _, ok := langMap[strings.ToLower(lang)]; ok {
		return langMap[lang]
	}
	return language.AmericanEnglish
}

func getPrinter(lang language.Tag) *message.Printer {
	if p, ok := printers[lang]; ok {
		return p
	}
	printers[lang] = message.NewPrinter(lang)
	return printers[lang]
}

func SetLanguage() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			// @TODO: change by chat settings
			ctx.Set(MessagePrinter, getPrinter(findLanguage(ctx.Sender().LanguageCode)))
			return next(ctx)
		}
	}
}

func T(ctx telebot.Context, format string, args ...any) string {
	p, ok := ctx.Get(MessagePrinter).(*message.Printer)
	if !ok {
		return fmt.Sprintf(format, args...)
	}
	return p.Sprintf(format, args...)
}
