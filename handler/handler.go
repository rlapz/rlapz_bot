package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rlapz/rlapz_bot/types"
)

func HandleFn(r *http.Request, url string) {
	var res types.Update
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		log.Println("HandleFn: ", err.Error())
		return
	}

	if res.UpdateId == 0 {
		log.Println("HandleFn: invalid update id")
		return
	}

	if res.Message.Chat.Type == types.ChatTypePrivate {
		var reply = "Hello, " + res.Message.From.Username

		log.Printf("id: %d, msg: %s\n", res.Message.From.Id, res.Message.Text)

		var req = fmt.Sprintf("{\"chat_id\":\"%v\", \"text\":\"%s\"}",
			res.Message.Chat.Id, reply,
		)

		var _url = url + "/sendMessage"
		hreq, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(req))
		if err != nil {
			log.Println("NewRequest: ", err.Error())
			return
		}

		log.Println("url: ", _url)

		hreq.Header.Set("Content-Type", "application/json; charset=UTF-8")
		var cl http.Client
		resp, err := cl.Do(hreq)
		if err != nil {
			log.Println("Do: ", err.Error())
			return
		}

		defer resp.Body.Close()

		log.Println("status: ", resp.Status)
	}
}
