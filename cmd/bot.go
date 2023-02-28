package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/litsea/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/internal/handler/character"
	"github.com/lostsnow/keqing/internal/handler/material"
	"github.com/lostsnow/keqing/pkg/i18n"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "serve bot",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := newBot()
		if err != nil {
			logger.Error(err)
			return
		}

		b.Use(i18n.SetLanguage())
		setHandler(b)

		b.Start()
	},
}

func newBot() (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:   viper.GetString("telegram.token"),
		Poller:  &telebot.LongPoller{Timeout: viper.GetDuration("telegram.poll-duration")},
		Verbose: viper.GetBool("telegram.verbose"),
		OnError: func(err error, context telebot.Context) {
			logger.Errorf("bot error: %s", err)
		},
	}

	proxy := viper.GetString("proxy")
	if proxy != "" {
		p, err := url.Parse(proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid p url: %s", proxy)
		}

		pref.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(p),
			},
		}
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("initialize bot failed: %s", err)
	}

	return b, nil
}

func setHandler(b *telebot.Bot) {
	if viper.GetBool("telegram.debug") {
		b.Use(middleware.Logger())
	}

	b.Handle("/help", handler.Help)
	b.Handle("/char_guide", character.Guide)
	b.Handle("/char_material", character.Material)
	b.Handle("/material_weekly", material.Weekly)

	ag := b.Group()
	admins := viper.GetStringSlice("admins")
	adminIds := make([]int64, 0, len(admins))
	for _, admin := range admins {
		id, err := strconv.ParseInt(admin, 10, 64)
		if err == nil {
			adminIds = append(adminIds, id)
		}
	}
	ag.Use(middleware.Whitelist(adminIds...))

	ag.Handle("/trace", handler.Trace)
}
