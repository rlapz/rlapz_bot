package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rlapz/rlapz_bot/request"
	"github.com/rlapz/rlapz_bot/types"
)

func HandleFn(r *http.Request, url string) {
	var res types.Update
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		log.Println("HandleFn: json.NewDecoder: ", err.Error())
		return
	}

	if res.UpdateId == 0 {
		log.Println("HandleFn: invalid update id")
		return
	}

	if res.Message.Chat.Type == types.ChatTypePrivate {
		var reply = "Hello, " + res.Message.From.Username

		var req = fmt.Sprintf("{\"chat_id\":\"%v\", \"text\":\"%s\"}",
			res.Message.Chat.Id, reply,
		)

		go request.Send(url, "/sendMessage", req)
	}
}
