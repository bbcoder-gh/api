package endpoint

import (
	"fmt"
)

type methodType int8

const (
	Get methodType = iota + 1
	Post
	Put
	Delete
)

func (m methodType) String() string {
	switch m {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Put:
		return "PUT"
	case Delete:
		return "DELETE"
	}
	return fmt.Sprintf("%s(%d)", "Unsupported", m)
}

type IEndpoint interface {
	Method() string
	Path() string
}
type Endpoint struct {
	MethodType methodType
	Name       string
}

func (e *Endpoint) String() string {
	return fmt.Sprintf("%s %s", e.MethodType, e.Name)
}

func (e *Endpoint) Method() string {
	return e.MethodType.String()
}

func (e *Endpoint) Path() string {
	return e.Name
}
