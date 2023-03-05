package handler

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"gopkg.in/telebot.v3"

	"github.com/lostsnow/keqing/pkg/character"
)

func Trace(ctx telebot.Context) error {
	return ctx.Send(fmt.Sprintf("*Bot*: `%d`\n*Chat*: `%d` <%s>\n*From*: `%d`\n*Message*: `%s`",
		ctx.Bot().Me.ID, ctx.Chat().ID, ctx.Chat().Type, ctx.Sender().ID, ctx.Text()),
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
}

func CacheAssetsClear(ctx telebot.Context) error {
	if len(ctx.Args()) < 2 {
		return ctx.Respond()
	}

	dir := ctx.Args()[0]
	name := strings.Join(ctx.Args()[1:], " ")
	if strings.HasPrefix(dir, "character/") {
		chars := character.Search(name)
		if len(chars) != 1 {
			return Send(ctx, "No such character")
		}

		cacheFilePath := getCacheFilePath(fmt.Sprintf("assets/%s/%s/%s.png", dir, chars[0].Elemental.Id, chars[0].Id))
		err := os.Remove(cacheFilePath)
		e, ok := err.(*os.PathError)
		if ok && e.Err != syscall.ENOENT {
			return ctx.Reply(fmt.Sprintf("clear %s cache failed: %s", cacheFilePath, err))
		}
		cacheFileIdPath := fmt.Sprintf("%s.id", cacheFilePath)
		err = os.Remove(cacheFileIdPath)
		e, ok = err.(*os.PathError)
		if ok && e.Err != syscall.ENOENT {
			return ctx.Reply(fmt.Sprintf("clear %s cache failed: %s", cacheFileIdPath, err))
		}
		return ctx.Send(fmt.Sprintf("cache clear successfully: %s(.id)", cacheFilePath))
	}

	return ctx.Respond()
}
