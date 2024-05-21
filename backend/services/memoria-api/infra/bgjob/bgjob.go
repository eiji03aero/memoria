package bgjob

import (
	"log"

	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/infra/registry"
)

type BGJob struct {
	regb *registry.Builder
}

func New(regb *registry.Builder) interfaces.BGJob {
	return &BGJob{regb: regb}
}

func (b *BGJob) Start() (err error) {
	for payload := range b.regb.GetBGJobInvokeChan() {
		log.Println("BGJob received payload", payload.Type, payload.Value)

		reg, err := b.regb.Build()
		if err != nil {
			panic(err)
		}

		switch payload.Type {
		case "create-thumbnails":
			b.createThumbnails(reg, payload.Value.(interfaces.BGJobInvokerCreateThumbnailsDTO))
		}
	}

	return
}

func (b *BGJob) createThumbnails(reg interfaces.Registry, dto interfaces.BGJobInvokerCreateThumbnailsDTO) (err error) {
	err = usecase.NewMedium(reg).CreateThumbnails(usecase.MediumCreateThumbnailsDTO{
		MediumID: dto.MediumID,
	})
	return
}
