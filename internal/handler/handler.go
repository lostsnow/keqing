package handler

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func Help(c telebot.Context) error {
	return c.Send("斩尽牛杂")
}

func Trace(c telebot.Context) error {
	return c.Send(fmt.Sprintf("*Bot*: `%d`\n*Chat*: `%d` <%s>\n*From*: `%d`\n*Message*: `%s`",
		c.Bot().Me.ID, c.Chat().ID, c.Chat().Type, c.Sender().ID, c.Text()),
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
}
