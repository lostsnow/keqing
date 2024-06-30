package cmd

import (
	"github.com/litsea/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3/middleware"

	"github.com/lostsnow/keqing/internal/handler"
	"github.com/lostsnow/keqing/internal/handler/character"
	"github.com/lostsnow/keqing/internal/handler/game"
	"github.com/lostsnow/keqing/internal/handler/lab"
	"github.com/lostsnow/keqing/internal/handler/material"
	"github.com/lostsnow/keqing/internal/handler/weapon"
	"github.com/lostsnow/keqing/pkg/i18n"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "serve bot",
	Run: func(_ *cobra.Command, _ []string) {
		b, err := handler.NewBot()
		if err != nil {
			logger.Error(err)

			return
		}

		lab.RunQrcodeCheckWorker()

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
	b.Bot.Handle("/char_curve", character.Curve)
	b.Bot.Handle("/char_graduation", character.Graduation)
	b.Bot.Handle("/weapon_guide", weapon.Guide)
	b.Bot.Handle("/material_weekly", material.Weekly)
	b.Bot.Handle("/material_daily", material.Daily)
	b.Bot.Handle("/material_boss", material.Boss)
	b.Bot.Handle("/material_char", material.Character)
	b.Bot.Handle("/material_weapon", material.Weapon)

	b.Bot.Handle("/auth_qrcode", lab.AuthQrcode)
	b.Bot.Handle("/game_role", game.Role)
	b.Bot.Handle("/game_refresh_role", game.RefreshRole)

	ag := b.Bot.Group()
	ag.Use(middleware.Whitelist(b.AdminIDs...))

	ag.Handle("/clear_assets_cache", handler.CacheAssetsClear)
	ag.Handle("/trace", handler.Trace)
}

func startBot(b *handler.Bot) {
	b.Bot.Start()
}
