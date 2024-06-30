package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/character"
	"github.com/lostsnow/keqing/pkg/weapon"
)

var handlerBot *Bot

var ErrInvalidTelegramProxyURL = errors.New("invalid telegram proxy url")

type Bot struct {
	Bot        *telebot.Bot
	AdminIDs   []int64
	AdminUsers []telebot.User
}

func NewBot() (*Bot, error) {
	if err := initModel(); err != nil {
		return nil, err
	}

	admins := viper.GetStringSlice("admins")
	adminIDs := make([]int64, 0, len(admins))
	adminUsers := make([]telebot.User, 0, len(admins))

	for _, admin := range admins {
		id, err := strconv.ParseInt(admin, 10, 64)
		if err == nil {
			adminIDs = append(adminIDs, id)
			adminUsers = append(adminUsers, telebot.User{ID: id})
		}
	}

	b := &Bot{
		AdminIDs:   adminIDs,
		AdminUsers: adminUsers,
	}

	pref := telebot.Settings{
		Token:   viper.GetString("telegram.token"),
		Poller:  &telebot.LongPoller{Timeout: viper.GetDuration("telegram.poll-duration")},
		Verbose: viper.GetBool("telegram.verbose"),
		OnError: func(err error, ctx telebot.Context) {
			msgBytes, e := json.MarshalIndent(ctx.Update(), "", "  ")
			if e != nil {
				ReportError(ctx, "*%s\n[ERROR] %s*\n```\n%v\n```", time.Now(), e, ctx.Update())

				return
			}
			ReportError(ctx, "*%s\n[ERROR] %s*\n```\n%s\n```", time.Now(), err, msgBytes)
		},
	}

	proxy := viper.GetString("proxy.telegram")
	if proxy != "" {
		p, err := url.Parse(proxy)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidTelegramProxyURL, proxy)
		}

		pref.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(p),
			},
		}
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("initialize bot failed: %w", err)
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
	initFn(weapon.Init)

	return err
}
