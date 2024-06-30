package handler

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/character"
)

func Trace(ctx telebot.Context) error {
	err := ctx.Send(fmt.Sprintf("*Bot*: `%d`\n*Chat*: `%d` <%s>\n*From*: `%d`\n*Message*: `%s`",
		ctx.Bot().Me.ID, ctx.Chat().ID, ctx.Chat().Type, ctx.Sender().ID, ctx.Text()),
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	if err != nil {
		return fmt.Errorf("handler.Trace: %w", err)
	}

	return nil
}

func CacheAssetsClear(ctx telebot.Context) error {
	if len(ctx.Args()) < 2 {
		return Respond(ctx)
	}

	dir := ctx.Args()[0]
	name := strings.Join(ctx.Args()[1:], " ")

	if strings.HasPrefix(dir, "character/") {
		chars := character.Search(name)
		if len(chars) != 1 {
			return Send(ctx, "No such character")
		}

		cacheFilePath := getCacheFilePath(fmt.Sprintf("assets/%s/%s/%s.png", dir, chars[0].Elemental.ID, chars[0].ID))
		err := os.Remove(cacheFilePath)

		var e *os.PathError

		ok := errors.As(err, &e)
		if ok && !errors.Is(e.Err, syscall.ENOENT) {
			return Reply(ctx, fmt.Sprintf("clear %s cache failed: %s", cacheFilePath, err))
		}

		cacheFileIDPath := cacheFilePath + ".id"
		err = os.Remove(cacheFileIDPath)
		ok = errors.As(err, &e)

		if ok && !errors.Is(e.Err, syscall.ENOENT) {
			return Reply(ctx, fmt.Sprintf("clear %s cache failed: %s", cacheFileIDPath, err))
		}

		return Send(ctx, fmt.Sprintf("cache clear successfully: %s(.id)", cacheFilePath))
	}

	return Respond(ctx)
}
