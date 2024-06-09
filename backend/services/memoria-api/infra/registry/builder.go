package registry

import (
	"context"

	"memoria-api/domain/interfaces"
	"memoria-api/infra/caws"
	"memoria-api/infra/db"

	"gorm.io/gorm"
)

type Builder struct {
	bgjobInvokeChan chan interfaces.BGJobInvokePayload
}

func NewBuilder() *Builder {
	return &Builder{
		bgjobInvokeChan: make(chan interfaces.BGJobInvokePayload),
	}
}

type BuilderBuildDTO struct {
	InitDB *bool
}

func (b *Builder) Build(dto BuilderBuildDTO) (reg interfaces.Registry, err error) {
	database, err := func() (database *gorm.DB, err error) {
		if dto.InitDB != nil && *dto.InitDB == false {
			return
		}

		database, err = db.New()
		return
	}()

	awsCfg, err := caws.LoadConfig(context.TODO())
	if err != nil {
		return
	}

	reg = &Registry{
		DB:              database,
		awsCfg:          awsCfg,
		bgjobInvokeChan: b.bgjobInvokeChan,
	}
	return
}

func (b *Builder) GetBGJobInvokeChan() chan interfaces.BGJobInvokePayload {
	return b.bgjobInvokeChan
}
