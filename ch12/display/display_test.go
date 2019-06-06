package display

import (
	"testing"
)

func BenchmarkDisplay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Display("v", []int{1, 2, 3})
	}
}

func Example_struct() {
	user := struct {
		Id      int
		Name    string
		Address string
		Course  struct {
			Id    int
			Title string
		}
	}{
		Id:      1,
		Name:    "wuzhc",
		Address: "guangzhou",
		Course: struct {
			Id    int
			Title string
		}{
			Id:    1,
			Title: "xxx",
		},
	}
	Display("user", user)

	// Output:
	// user.Id=1
	// user.Name=wuzhc
	// user.Address=guangzhou
	// user.Course.Id=1
	// user.Course.Title=xxx
}

func Example_ptr() {
	type User struct {
		Id   int
		Name string
	}
	var user = &User{Id: 1, Name: "wuzhc"}
	Display("user", user)

	// Output:
	// *user.Id=1
	// *user.Name=wuzhc
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)

	// Output:
	// i=3
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
