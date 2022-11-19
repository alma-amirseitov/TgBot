package bot

import (
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
