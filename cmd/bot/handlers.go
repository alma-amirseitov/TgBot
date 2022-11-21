package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alma-amirseitov/TgBot/internal/models"
	"gopkg.in/telebot.v3"
)

func (bot *Bot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)

		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет" + ctx.Sender().FirstName)
}

func (bot *Bot) MessagePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("ParseForm() err: %v", err)
		return
	}
	message := r.FormValue("message")
	fmt.Printf("message = %s\n", message)
}

func (bot *Bot) MessageHandler(ctx telebot.Context) error {

	existUsers, err := bot.Users.FindAll()
	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}
	for _, existUser := range existUsers {
		log.Printf("пользователя %v", existUser.TelegramId)
		return ctx.Send("sdfdf")
	}
	return ctx.Send("Привет")
}
