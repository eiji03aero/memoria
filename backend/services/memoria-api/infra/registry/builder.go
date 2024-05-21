package registry

import (
	"context"

	"memoria-api/domain/interfaces"
	"memoria-api/infra/caws"
	"memoria-api/infra/db"
)

type Builder struct {
	bgjobInvokeChan chan interfaces.BGJobInvokePayload
}

func NewBuilder() *Builder {
	return &Builder{
		bgjobInvokeChan: make(chan interfaces.BGJobInvokePayload),
	}
}

func (b *Builder) Build() (reg interfaces.Registry, err error) {
	db, err := db.New()
	if err != nil {
		return
	}

	awsCfg, err := caws.LoadConfig(context.TODO())
	if err != nil {
		return
	}

	reg = &Registry{
		DB:              db,
		awsCfg:          awsCfg,
		bgjobInvokeChan: b.bgjobInvokeChan,
	}
	return
}

func (b *Builder) GetBGJobInvokeChan() chan interfaces.BGJobInvokePayload {
	return b.bgjobInvokeChan
}
