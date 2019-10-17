package jsonfunc

import "reflect"

type Element struct {
	Type  reflect.Type //element type
	Key   interface{}  //element key
	Value *Element     //sub element
	Data  interface{}  //element data
}
