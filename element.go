package jsonfunc

import (
	"fmt"
	"reflect"
)

type Element interface {
	Key() interface{}
	Value() Element
}

type elementImpl struct {
	elementType reflect.Type
	key         interface{}
	data        interface{}
	value       *elementImpl
}

func (e elementImpl) Key() interface{} {
	return e.key
}

func (e elementImpl) Value() Element {
	return e.value
}

func NewElement(ele interface{}) Element {
	t := reflect.TypeOf(ele)
	//v := reflect.ValueOf(ele)
	fmt.Println(t.Field(0).Tag.Get("name"))

	return &elementImpl{
		elementType: t,
		key:         nil,
		data:        nil,
		value:       nil,
	}
}
