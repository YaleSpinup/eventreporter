package eventreporter

import (
	"fmt"
	"net/http"
)

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
	Report(Event) error
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(name string, config map[string]string) (Reporter, error) {
	switch name {
	case "slack":
		reporter, err := NewSlackReporter(config)
		if err != nil {
			return nil, err
		}

		return reporter, nil
	case "email":
		reporter, err := NewSlackReporter(config)
		if err != nil {
			return nil, err
		}

		return reporter, nil
	default:
		return nil, fmt.Errorf("Unknown event reporter name, %s", name)
	}
}
