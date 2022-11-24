package main

import (
	"flag"
	"github.com/alma-amirseitov/TgBot/cmd/app"
	"github.com/alma-amirseitov/TgBot/cmd/bot"
	"github.com/alma-amirseitov/TgBot/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}
	if !db.Migrator().HasTable(&models.User{}) {
		db.Migrator().CreateTable(&models.User{})
	}

	bot := bot.Bot{
		Bot:   bot.InitBot(cfg.BotToken),
		Users: &models.UserModel{Db: db},
	}
	app := app.Application{
		Bot: &bot,
	}
	go func() {
		err = app.Serve(":8081")
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	bot.Bot.Handle("/start", bot.StartHandler)
	bot.Bot.Start()

}
