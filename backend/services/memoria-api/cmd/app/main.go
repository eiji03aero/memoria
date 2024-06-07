package main

import (
	"memoria-api/infra/bgjob"
	"memoria-api/infra/registry"
	"memoria-api/infra/route"

	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	regb := registry.NewBuilder()
	reg, err := regb.Build(registry.BuilderBuildDTO{})
	if err != nil {
		panic(err)
	}

	// -------------------- bgjob --------------------
	bgj := bgjob.New(regb)
	go bgj.Start()

	// -------------------- router --------------------
	r := route.InitializeRouter(regb)

	// -------------------- server --------------------
	reg.NewLogger().Info("Starting memoria-api server")
	err = r.Run("0.0.0.0:4200")
	if err != nil {
		panic(err)
	}
}
