package lab

import (
	"bytes"
	"context"
	"fmt"

	"github.com/litsea/logger"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/db"
	"github.com/lostsnow/keqing/internal/lab/api"
	"github.com/lostsnow/keqing/pkg/entity/user"
	"github.com/lostsnow/keqing/pkg/i18n"
)

var qrcodeCheckPool *api.QrcodeCheckPool

func RunQrcodeCheckWorker() {
	logger.Info("running QR code check worker")

	qrcodeCheckPool = api.NewQrcodeCheckPool()

	go qrcodeCheckPool.Worker()
}

func AuthQrcode(ctx telebot.Context) error {
	if ctx.Chat().Type != telebot.ChatPrivate {
		return ctx.Reply(i18n.T(ctx, "Please send me a private chat message"))
	}

	userID := ctx.Sender().ID
	if qrcodeCheckPool.IsRunning(userID) {
		return ctx.Reply(i18n.T(ctx, "QR code login is being checked"))
	}

	err := db.DB.Client.User.
		Create().
		SetUserID(userID).
		SetIsBot(ctx.Sender().IsBot).
		SetUserName(ctx.Sender().Username).
		SetFirstName(ctx.Sender().FirstName).
		SetLastName(ctx.Sender().LastName).
		OnConflictColumns(user.FieldUserID).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		_ = ctx.Reply(i18n.T(ctx, "Upsert user failed"))

		return fmt.Errorf("upsert user %d failed: %w", userID, err)
	}

	req := api.NewQrcodeFetch()

	resp, err := req.Do()
	if err != nil {
		_ = ctx.Reply(i18n.T(ctx, "Fetch QR code failed"))

		return err
	}

	qr, err := req.ToImage(resp.URL)
	if err != nil {
		_ = ctx.Reply(i18n.T(ctx, "Generate QR code failed"))

		return err
	}

	qrcodeCheckPool.Add(ctx, userID, api.NewQrcodeQuery(resp.AppID, resp.Device, resp.Ticket))

	return ctx.Reply(&telebot.Photo{
		File:    telebot.FromReader(bytes.NewReader(qr)),
		Caption: i18n.T(ctx, "Please scan the QR code with the miyoushe App"),
	})
}
