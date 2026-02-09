package main

import (
	"log"

	router "github.com/saf1o/go-test/internal/controller"
	db "github.com/saf1o/go-test/internal/model"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	r := router.Setup()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
