package handler

import (
	"fmt"
	"net/http"
)

func HandleFn(r *http.Request, url string) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Body)
}
