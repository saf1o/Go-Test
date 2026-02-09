package main

import (
	"log"
	"net/http"

	"github.com/saf1o/go-test/internal/controller"

	"github.com/saf1o/go-test/internal/model"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/game?parseTime=true"

	if err := model.InitDB(dsn); err != nil {
		log.Fatal(err)
	}

	controller.InitRouter()

	log.Panicln("server start :8080")
	http.ListenAndServe(":8080", nil)
}
