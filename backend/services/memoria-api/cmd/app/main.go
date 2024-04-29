package main

import (
	"memoria-api/route"
)

func main() {
	r := route.InitializeRouter()

	err := r.Run("0.0.0.0:4200")
	if err != nil {
		panic(err.Error())
	}
}
