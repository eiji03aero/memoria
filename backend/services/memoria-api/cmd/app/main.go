package main

import (
	"log"

	"memoria-api/infra/bgjob"
	"memoria-api/infra/registry"
	"memoria-api/infra/route"

	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	regb := registry.NewBuilder()

	// -------------------- bgjob --------------------
	bgj := bgjob.New(regb)
	go bgj.Start()

	// -------------------- router --------------------
	r := route.InitializeRouter(regb)

	// -------------------- server --------------------
	log.Println("Starting memoria-api server")
	err := r.Run("0.0.0.0:4200")
	if err != nil {
		panic(err.Error())
	}
}
