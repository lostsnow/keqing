package cmd

import (
	"fmt"
	"github.com/litsea/logger"
	"github.com/lostsnow/keqing/internal/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"net/http"
	"net/url"
	"strconv"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Bot",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := newBot()
		if err != nil {
			logger.Error(err)
			return
		}

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
	b.Handle("/help", handler.Help)

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

func init() {
	rootCmd.AddCommand(botCmd)
}
