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
	return &elementImpl{
		elementType: t,
		key:         t.Name(),
		data:        ElementData(t.Name(), reflect.ValueOf(ele)),
		value:       nil,
	}
}

func ElementData(name string, value reflect.Value) interface{} {
	var elements []Element
	fmt.Println("name:", name)
	fmt.Println("print value:", value.String())
	switch value.Kind() {
	case reflect.Slice:
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			elements = append(elements, &elementImpl{
				elementType: nil,
				key:         value.Field(i).String(),
				data:        ElementData(value.Field(i).Type().Name(), value.Field(i)),
				value:       nil,
			})
		}
	}
	return elements

}
