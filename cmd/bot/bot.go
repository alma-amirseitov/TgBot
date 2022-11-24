package bot

import (
	"errors"
	"fmt"
	"github.com/alma-amirseitov/TgBot/internal/models"
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

type Bot struct {
	Bot   *telebot.Bot
	Users *models.UserModel
}

func InitBot(token string) *telebot.Bot {

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}

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

	return ctx.Send("Привет " + ctx.Sender().FirstName)
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
