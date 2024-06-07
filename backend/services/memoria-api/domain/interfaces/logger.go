package interfaces

type Logger interface {
	Debug(messages ...any)
	Info(messages ...any)
	Warn(messages ...any)
	Error(messages ...any)
}
