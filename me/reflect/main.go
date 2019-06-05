// reflect反射
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	v := struct {
		Id   int
		Name string
	}{
		Id:   1,
		Name: "xxx",
	}
	fmt.Println(formatAtom(v))
}

func formatAtom(v interface{}) string {
	value := reflect.ValueOf(v) // type : reflect.Value
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.String:
		return value.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Bool:
		if value.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Chan, reflect.Slice, reflect.Map, reflect.Func, reflect.Ptr:
		return value.Type().String() + "0x" + strconv.FormatUint(uint64(value.Pointer()), 16)
	default: // reflect.Array, reflect.Srtuct, reflect.Interface
		return value.Type().String() + " value"
	}
}

func display(path string, v interface{}) {

}
