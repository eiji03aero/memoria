package main

import (
	"memoria-api/route"
)

func main() {
	r := route.InitializeRouter()

	r.Run("0.0.0.0:4200")
}
