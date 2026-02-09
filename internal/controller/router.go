package controller

import "net/http"

func InitRouter() {
	http.HandleFunc("/login", LoginHandler)
}
