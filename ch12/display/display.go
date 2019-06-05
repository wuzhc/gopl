// reflect反射
package display

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	v := []int{1, 2, 3, 4}
	display("v", reflect.ValueOf(v))

	v1 := make(map[string]string)
	v1["name"] = "wuzhc"
	v1["address"] = "guangzhou"
	display("v1", reflect.ValueOf(v1))
}

func Display(name string, v interface{}) {
	display(name, reflect.ValueOf(v))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.String:
		return v.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Chan, reflect.Slice, reflect.Map, reflect.Func, reflect.Ptr:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Srtuct, reflect.Interface
		return v.Type().String() + " v"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s=invalid\n", path)
	case reflect.Slice, reflect.Array:
		// 对数组和切片的处理
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%v]", path, i), v.Index(i))
		}
	case reflect.Map:
		//对map的处理
		for _, k := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, k.String()), v.MapIndex(k))
		}
	default:
		fmt.Printf("%s=%s\n", path, formatAtom(v))
	}
}
