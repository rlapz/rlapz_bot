package config

import (
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

const base_url = "https://api.telegram.org/bot"

type Config struct {
	Token      string
	Url        string
	Secret     string
	ListenAddr string
}

func ConfigInit(env string) Config {
	if err := godotenv.Load(env); err != nil {
		log.Fatalln("config.ConfigInit: ", err.Error())
	}

	var token = os.Getenv("TOKEN")
	if len(token) == 0 {
		log.Fatalln("config.ConfigInit: .env.TOKEN: empty")
	}

	var secret = os.Getenv("SECRET")
	if len(secret) == 0 {
		log.Fatalln("config.ConfigInit: .env.SECRET: empty")
	}

	var listenAddr = os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		log.Fatalln("config.ConfigInit: .env.LISTEN_ADDR: empty")
	}

	return Config{
		Token:      token,
		Url:        base_url + token,
		Secret:     secret,
		ListenAddr: listenAddr,
	}
}
