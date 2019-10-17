package jsonfunc_test

import (
	"github.com/goextension/jsonfunc"
	"testing"
)

type FieldStruct struct {
	Field1 string `json:"field_1"`
}

type Fields struct {
	Fields     []FieldStruct `json:"fields"`
	OtherField FieldStruct   `json:"other_field"`
}

func TestNewElement(t *testing.T) {
	jsonfunc.NewElement(Fields{})
}
