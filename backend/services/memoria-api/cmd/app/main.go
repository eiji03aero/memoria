package main

import (
	"log"

	"memoria-api/infra/route"
)

func main() {
	r := route.InitializeRouter()

	log.Println("Starting memoria-api server")
	err := r.Run("0.0.0.0:4200")
	if err != nil {
		panic(err.Error())
	}
}
