package server

import (
	"context"
	"log"
	"net/http"

	"github.com/rlapz/rlapz_bot/config"
	"github.com/rlapz/rlapz_bot/handler"
)

type Server struct {
	config *config.Config
	server http.Server
}

func ServerInit(config *config.Config) Server {
	return Server{
		config,
		http.Server{
			Addr: config.ListenAddr,
		},
	}
}

func (self *Server) Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("header: ", r.Header)

		secret, ok := r.Header["X-Telegram-Bot-Api-Secret-Token"]

		if ok && r.Method == http.MethodPost {
			log.Println("secret: ", secret)
			handler.HandleFn(r, self.config.Url)
		}
	})

	log.Printf("running on: %s\n", self.config.ListenAddr)
	return self.server.ListenAndServe()
}

func (self *Server) Stop(ctx context.Context) error {
	log.Println("\nstopping...")
	return self.server.Shutdown(ctx)
}
