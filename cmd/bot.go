package cmd

import (
	"github.com/litsea/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3/middleware"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/internal/handler/character"
	"github.com/lostsnow/keqing/internal/handler/material"
	"github.com/lostsnow/keqing/pkg/i18n"
	_ "github.com/lostsnow/keqing/pkg/i18n/catalog"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "serve bot",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := handler.NewBot()
		if err != nil {
			logger.Error(err)
			return
		}

		setBotHandler(b)
		startBot(b)
	},
}

func setBotHandler(b *handler.Bot) {
	b.Bot.Use(i18n.SetLanguage())
	if viper.GetBool("telegram.debug") {
		b.Bot.Use(middleware.Logger())
	}
	b.Bot.Use(middleware.Recover())

	b.Bot.Handle("/help", handler.Help)
	b.Bot.Handle("/char_guide", character.Guide)
	b.Bot.Handle("/char_material", character.Material)
	b.Bot.Handle("/material_weekly", material.Weekly)

	ag := b.Bot.Group()
	ag.Use(middleware.Whitelist(b.AdminIds...))

	ag.Handle("/clear_assets_cache", handler.CacheAssetsClear)
	ag.Handle("/trace", handler.Trace)
}

func startBot(b *handler.Bot) {
	b.Bot.Start()
}
