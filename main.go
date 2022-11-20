package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/alma-amirseitov/TgBot/cmd/bot"
	"github.com/alma-amirseitov/TgBot/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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

	go serverForPOST()

	http.HandleFunc("/message", bot.MessagePostHandler)

	bot.Bot.Handle("/start", bot.StartHandler)
	bot.Bot.Handle("/message", bot.MessageHandler)
	bot.Bot.Start()

}

func serverForPOST() {
	fmt.Printf("Starting server for HTTP POST...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
