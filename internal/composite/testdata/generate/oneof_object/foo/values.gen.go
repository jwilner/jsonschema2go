// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"encoding/json"
	"fmt"
)

// Bar is generated from https://example.com/testdata/generate/oneof_object/foo/example.json
// Bar gives you some dumb info
type Bar struct {
	_         [0]byte
	Direction interface{}
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/oneof_object/foo/example.json
func (m *Bar) Validate() error {
	return nil
}

func (m *Bar) UnmarshalJSON(data []byte) error {
	var discrim struct {
		Direction string `json:"direction"`
	}
	if err := json.Unmarshal(data, &discrim); err != nil {
		return err
	}
	switch discrim.Direction {
	case "l":
		m.Direction = new(Left)
	case "r":
		m.Direction = new(Right)
	default:
		return fmt.Errorf("unknown discriminator: %v", discrim.Direction)
	}
	return json.Unmarshal(data, m.Direction)
}

func (m *Bar) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Direction)
}

// Left is generated from https://example.com/testdata/generate/oneof_object/foo/example.json#/oneOf/0
type Left struct {
	_         [0]byte
	Direction *string `json:"direction,omitempty"`
	Value     *int64  `json:"value,omitempty"`
}

var (
	leftDirectionEnum = map[string]bool{"l": true}
)

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/oneof_object/foo/example.json#/oneOf/0
func (m *Left) Validate() error {
	if m.Direction != nil && !leftDirectionEnum[*m.Direction] {
		return &validationError{
			errType:  "enum",
			path:     []interface{}{"Direction"},
			jsonPath: []interface{}{"direction"},
			message:  fmt.Sprintf(`must be "l" but got %v`, *m.Direction),
		}
	}
	return nil
}

// Right is generated from https://example.com/testdata/generate/oneof_object/foo/example.json#/oneOf/1
type Right struct {
	_         [0]byte
	Direction *string  `json:"direction,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}

var (
	rightDirectionEnum = map[string]bool{"r": true}
)

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/oneof_object/foo/example.json#/oneOf/1
func (m *Right) Validate() error {
	if m.Direction != nil && !rightDirectionEnum[*m.Direction] {
		return &validationError{
			errType:  "enum",
			path:     []interface{}{"Direction"},
			jsonPath: []interface{}{"direction"},
			message:  fmt.Sprintf(`must be "r" but got %v`, *m.Direction),
		}
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
