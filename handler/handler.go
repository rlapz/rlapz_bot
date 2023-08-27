package handler

import (
	"encoding/json"
	"log"
	"net/http"

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

	log.Println("text: ", res.Message.Text)
}
