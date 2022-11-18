package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Env string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	log.Println(cfg.Env)
}
