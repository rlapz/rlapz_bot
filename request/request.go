package request

import (
	"log"
	"net/http"
	"strings"
)

func Send(url, method, data string) {
	var _url = url + method
	var reader = strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, _url, reader)
	if err != nil {
		log.Println("request.Send: http.NewRequest: ", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Println("request.Send: http.Client.Do: ", err.Error())
		return
	}

	defer resp.Body.Close()

	log.Println("request.Send: status: ", resp.Status)
}
