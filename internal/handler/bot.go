package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/character"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

var (
	handlerBot *Bot
)

type Bot struct {
	Bot        *telebot.Bot
	AdminIds   []int64
	AdminUsers []telebot.User
}

func NewBot() (*Bot, error) {
	if err := initModel(); err != nil {
		return nil, err
	}

	admins := viper.GetStringSlice("admins")
	adminIds := make([]int64, 0, len(admins))
	adminUsers := make([]telebot.User, 0, len(admins))
	for _, admin := range admins {
		id, err := strconv.ParseInt(admin, 10, 64)
		if err == nil {
			adminIds = append(adminIds, id)
			adminUsers = append(adminUsers, telebot.User{ID: id})
		}
	}
	b := &Bot{
		AdminIds:   adminIds,
		AdminUsers: adminUsers,
	}

	pref := telebot.Settings{
		Token:   viper.GetString("telegram.token"),
		Poller:  &telebot.LongPoller{Timeout: viper.GetDuration("telegram.poll-duration")},
		Verbose: viper.GetBool("telegram.verbose"),
		OnError: func(err error, ctx telebot.Context) {
			msgBytes, _ := json.MarshalIndent(ctx.Update(), "", "  ")
			msg := fmt.Sprintf("*%s\n[ERROR] %s*\n```\n%s\n```", time.Now(), err, msgBytes)
			ReportError(ctx, msg)
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

	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("initialize bot failed: %s", err)
	}
	b.Bot = bot
	handlerBot = b

	return b, nil
}

func initModel() error {
	var err error
	initFn := func(fn func() error) {
		if err != nil {
			return
		}
		err = fn()
	}

	initFn(character.Init)
	return err
}
