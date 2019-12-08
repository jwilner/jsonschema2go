// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"encoding/json"
	"fmt"
	"github.com/jwilner/jsonschema2go/boxed"
)

// Bar gives you some dumb info
type Bar struct {
	Foo Foo `json:"foo,omitempty"`
}

func (m *Bar) Validate() error {
	if err := m.Foo.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Foo"}, err.Path()...),
				jsonPath: append([]interface{}{"foo"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

type Baz struct {
	Name boxed.String `json:"name"`
}

func (m *Baz) Validate() error {
	return nil
}

func (m *Baz) MarshalJSON() ([]byte, error) {
	inner := struct {
		Name *string `json:"name,omitempty"`
	}{}
	if m.Name.Set {
		inner.Name = &m.Name.String
	}
	return json.Marshal(inner)
}

type Foo struct {
	Baz Baz `json:"baz,omitempty"`
}

func (m *Foo) Validate() error {
	if err := m.Baz.Validate(); err != nil {
		if err, ok := err.(valErr); ok {
			return &validationError{
				errType:  err.ErrType(),
				message:  err.Message(),
				path:     append([]interface{}{"Baz"}, err.Path()...),
				jsonPath: append([]interface{}{"baz"}, err.JSONPath()...),
			}
		}
		return err
	}
	return nil
}

type valErr interface {
	ErrType() string
	JSONPath() []interface{}
	Path() []interface{}
	Message() string
}

type validationError struct {
	errType, message string
	jsonPath, path   []interface{}
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONPath() []interface{} {
	return e.jsonPath
}

func (e *validationError) Path() []interface{} {
	return e.path
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.path, e.message)
}

var _ valErr = new(validationError)