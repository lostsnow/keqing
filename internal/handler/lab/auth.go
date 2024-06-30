package lab

import (
	"bytes"
	"context"
	"fmt"

	"github.com/litsea/logger"
	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/internal/db"
	"github.com/lostsnow/keqing/internal/handler"
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
		err := handler.Reply(ctx, i18n.T(ctx, "Please send me a private chat message"))
		if err != nil {
			return fmt.Errorf("auth.AuthQrcode: %w", err)
		}
	}

	userID := ctx.Sender().ID
	if qrcodeCheckPool.IsRunning(userID) {
		err := handler.Reply(ctx, i18n.T(ctx, "QR code login is being checked"))
		if err != nil {
			return fmt.Errorf("auth.AuthQrcode: %w", err)
		}
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
		_ = handler.Reply(ctx, i18n.T(ctx, "Upsert user failed"))

		return fmt.Errorf("upsert user %d failed: %w", userID, err)
	}

	req := api.NewQrcodeFetch()

	resp, err := req.Do()
	if err != nil {
		_ = handler.Reply(ctx, i18n.T(ctx, "Fetch QR code failed"))

		return fmt.Errorf("lab.AuthQrcode: %w", err)
	}

	qr, err := req.ToImage(resp.URL)
	if err != nil {
		_ = handler.Reply(ctx, i18n.T(ctx, "Generate QR code failed"))

		return fmt.Errorf("lab.AuthQrcode: %w", err)
	}

	qrcodeCheckPool.Add(ctx, userID, api.NewQrcodeQuery(resp.AppID, resp.Device, resp.Ticket))

	err = handler.ReplyPhoto(ctx, &telebot.Photo{
		File:    telebot.FromReader(bytes.NewReader(qr)),
		Caption: i18n.T(ctx, "Please scan the QR code with the miyoushe App"),
	})
	if err != nil {
		return fmt.Errorf("auth.AuthQrcode: %w", err)
	}

	return nil
}
