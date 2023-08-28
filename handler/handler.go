package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rlapz/rlapz_bot/request"
	"github.com/rlapz/rlapz_bot/types"
)

func Handle(r *http.Request, url string) {
	var res types.Update
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		log.Println("HandleFn: json.NewDecoder: ", err.Error())
		return
	}

	if res.UpdateId == 0 {
		log.Println("HandleFn: invalid update id")
		return
	}

	log.Println("Chat type: ", res.Message.Chat.Type)

	switch res.Message.Chat.Type {
	case types.ChatTypePrivate:
		go handleChatPrivate(url, &res)
	case types.ChatTypeGroup:
		go handleChatGroup(url, &res)
	case types.ChatTypeSuperGroup:
		go handleChatSuperGroup(url, &res)
	}
}

func handleChatPrivate(url string, update *types.Update) {
	var reply = "Hello, " + update.Message.From.Username

	var req = fmt.Sprintf("{\"chat_id\":\"%v\", \"text\":\"%s\"}",
		update.Message.Chat.Id, reply,
	)

	request.Send(url, "/sendMessage", req)
}

func handleChatGroup(url string, update *types.Update) {
	var reply = "Hello, " + update.Message.From.Username + " ;)"

	var req = fmt.Sprintf("{\"chat_id\":\"%v\", \"text\":\"%s\"}",
		update.Message.Chat.Id, reply,
	)

	request.Send(url, "/sendMessage", req)
}

func handleChatSuperGroup(url string, update *types.Update) {

}
