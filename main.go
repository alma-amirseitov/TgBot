package main

import (
	"flag"
	"github.com/alma-amirseitov/TgBot/cmd/bot"
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/telebot.v3"
)

type Config struct {
	Env      string
	BotToken string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	bot := bot.Bot{
		Bot: bot.InitBot(cfg.BotToken),
	}

	bot.Bot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("Привет, " + ctx.Sender().FirstName)
	})
	bot.Bot.Start()
}
