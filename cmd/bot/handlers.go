package bot

import (
	"errors"
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
	err := bot.MessageHandler(message)
	if err != nil {
		log.Printf("%v", err)
	}
}

func (bot *Bot) MessageHandler(msg string) error {
	if msg == "" {
		return errors.New("ошибка отправки пустого сообщения")
	}
	users, err := bot.Users.FindAll()
	if err != nil {
		return err
	}
	var usersWitError []int64
	for _, user := range users {
		_, err = bot.Bot.Send(user, msg)
		if err != nil {
			usersWitError = append(usersWitError, user.TelegramId)
		}
	}
	if len(usersWitError) != 0 {
		return fmt.Errorf("ошибка при отправке пользовотелям: %v", usersWitError)
	}
	return nil
}
