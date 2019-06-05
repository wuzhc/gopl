package display

import (
	"testing"
)

func BenchmarkDisplay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Display("v", []int{1, 2, 3})
	}
}

func Example_array() {
	v := []int{1, 2, 3, 4}
	Display("v", v)

	// Output:
	// v[0]=1
	// v[1]=2
	// v[2]=3
	// v[3]=4
}

func Example_map() {
	user := make(map[string]string)
	user["name"] = "wuzhc"
	user["address"] = "guangzhou"
	Display("user", user)

	// Output:
	// user[name]=wuzhc
	// user[address]=guangzhou
}
