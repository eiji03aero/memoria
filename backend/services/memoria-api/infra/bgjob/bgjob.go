package bgjob

import (
	"strconv"
	"sync"

	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/infra/registry"
)

type BGJob struct {
	regb       *registry.Builder
	resultChan chan *WorkResult
	workerQty  int
	wg         *sync.WaitGroup
}

type WorkResult struct {
	Type string
	Err  error
}

func New(regb *registry.Builder) interfaces.BGJob {
	return &BGJob{
		regb:       regb,
		resultChan: make(chan *WorkResult),
		workerQty:  3,
		wg:         &sync.WaitGroup{},
	}
}

func (b *BGJob) Start() (err error) {
	reg, err := b.regb.Build(registry.BuilderBuildDTO{})
	if err != nil {
		return
	}

	for i := 0; i < b.workerQty; i++ {
		reg.NewLogger().Info("BGJob starting worker " + strconv.Itoa(i))
		go b.startWorker()
	}

	b.wg.Wait()
	return
}

func (b *BGJob) startWorker() {
	reg, err := b.regb.Build(registry.BuilderBuildDTO{})
	if err != nil {
		return
	}

	b.wg.Add(1)
	defer b.wg.Done()

	for payload := range b.regb.GetBGJobInvokeChan() {
		reg.NewLogger().Info("BGJob received payload", payload.Type, payload.Value)

		var e error
		reg, e := b.regb.Build(registry.BuilderBuildDTO{})
		if e != nil {
			b.sendResult("registry-builder-build", e)
			continue
		}

		switch payload.Type {
		case "create-thumbnails":
			e = b.createThumbnails(reg, payload.Value.(interfaces.BGJobInvokerCreateThumbnailsDTO))
		}
		b.sendResult(payload.Type, e)
	}

	return
}

func (b *BGJob) sendResult(t string, e error) {
	b.resultChan <- &WorkResult{Type: t, Err: e}
}

// -------------------- Works --------------------
func (b *BGJob) createThumbnails(reg interfaces.Registry, dto interfaces.BGJobInvokerCreateThumbnailsDTO) (err error) {
	err = usecase.NewMedium(reg).CreateThumbnails(usecase.MediumCreateThumbnailsDTO{
		MediumID: dto.MediumID,
	})
	return
}
