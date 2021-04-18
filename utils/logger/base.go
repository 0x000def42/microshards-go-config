package logger

type Logger interface {
	Servise() string

	Debug(message string, err error)
	Info(message string, err error)
	Warn(message string, err error)
	Error(message string, err error)
	Fatal(message string, err error)
}
