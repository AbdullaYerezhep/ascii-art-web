package main

import (
	"log"
	"project/internal/app"
)

func main() {
	err := app.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
