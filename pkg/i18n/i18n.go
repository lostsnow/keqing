package i18n

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/telebot.v3"
)

//go:generate go run -mod=mod github.com/lostsnow/keqing/cmd/i18n/generate
//go:generate gotext -srclang=en-US update -out=./catalog/catalog.go -lang=en-US,zh-Hans github.com/lostsnow/keqing

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

func TS(format string, args ...any) []string {
	ns := make([]string, 0, len(langMap))
	v := make(map[string]struct{})

	for _, l := range langMap {
		p := getPrinter(l)
		s := p.Sprintf(format, args...)

		if _, ok := v[s]; ok {
			continue
		}

		v[s] = struct{}{}

		ns = append(ns, s)
	}

	return ns
}
