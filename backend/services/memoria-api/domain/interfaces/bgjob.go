package interfaces

type BGJob interface {
	Start() (err error)
}

type BGJobInvokePayload struct {
	Type  string
	Value any
}

type BGJobInvoker interface {
	CreateThumbnails(dto BGJobInvokerCreateThumbnailsDTO)
}

type BGJobInvokerCreateThumbnailsDTO struct {
	MediumID string
}
