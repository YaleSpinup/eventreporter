package eventreporter

type Level uint32

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type Event struct {
	Message string
	Level   Level
}

type Reporter interface {
	Report(*Event) error
}
