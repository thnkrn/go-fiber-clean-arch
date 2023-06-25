package log

type Logger interface {
	Debug(msg string)
	Error(msg string)
	Fatal(msg string)
	Info(msg string)
}
