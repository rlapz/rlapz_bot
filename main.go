package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rlapz/rlapz_bot/config"
	"github.com/rlapz/rlapz_bot/server"
)

func main() {
	var config = config.ConfigInit(".env")
	var srv = server.ServerInit(&config)

	go func() {
		if err := srv.Run(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalln("server.Run: ", err.Error())
			}

			log.Println("server.Run: http server closed")
		}
	}()

	var qChan = make(chan os.Signal, 1)
	signal.Notify(qChan, os.Interrupt)
	<-qChan

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Println("server.Stop: ", err.Error())
	}
}
