package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rlapz/rlapz_bot/config"
	"github.com/rlapz/rlapz_bot/handler"
)

func main() {
	var config = config.ConfigInit(".env")
	var srv = http.Server{
		Addr: config.ListenAddr,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tok = r.Header.Get("X-Telegram-Bot-Api-Secret-Token")
		if tok == config.Secret && r.Method == http.MethodPost {
			handler.Handle(r, config.Url)
		}
	})

	go func(ctx *http.Server) {
		log.Printf("http.Server: running on: %s\n", config.ListenAddr)
		if err := ctx.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalln("http.ListenAndServe: ", err.Error())
			}

			log.Println("http.ListenAndServe: http server closed")
		}
	}(&srv)

	var qChan = make(chan os.Signal, 1)
	signal.Notify(qChan, os.Interrupt)
	<-qChan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("http.ListenAndServe: ", err.Error())
	}
}
