package error_factory

import (
	"fmt"

	"github.com/0x000def42/microshards-go-config/utils/access"
)

// Define AccessError
type AccessError struct {
	Service string
	Method  string
	Actor   access.Actor
}

func (err AccessError) Error() string {
	return fmt.Sprintf("Access denied while calling \"%s.%s\" with access \"%s\"", err.Service, err.Method, err.Actor)
}

// Define ErrorFactory
type ErrorFactory struct {
	Servise string
}

func (factory ErrorFactory) Access(method string, actor access.Actor) error {
	return &AccessError{
		Service: factory.Servise,
		Method:  method,
		Actor:   actor,
	}
}
