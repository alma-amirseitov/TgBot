package bot

import "gopkg.in/telebot.v3"

func (bot *Bot) StartHandler(ctx telebot.Context) error {
	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}
