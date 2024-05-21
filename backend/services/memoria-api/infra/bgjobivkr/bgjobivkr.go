package bgjobivkr

import "memoria-api/domain/interfaces"

type BGJobInvoker struct {
	invokeChan chan interfaces.BGJobInvokePayload
}

func NewBGJobInvoker(invokeChan chan interfaces.BGJobInvokePayload) interfaces.BGJobInvoker {
	return &BGJobInvoker{invokeChan: invokeChan}
}

func (i BGJobInvoker) CreateThumbnails(dto interfaces.BGJobInvokerCreateThumbnailsDTO) {
	payload := interfaces.BGJobInvokePayload{
		Type:  "create-thumbnails",
		Value: dto,
	}
	i.invokeChan <- payload
}
