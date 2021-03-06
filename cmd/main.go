package main

import "fmt"

func main() {
	// buf := bytes.NewBuffer([]byte(`{
	// 	"data": [
	// 		{"a": 1, "b": "hello"},
	// 		{"a": 2, "b": "world"}
	// 	]
	// 	}`))
	// j, _ := json.NewFromReader(buf)
	// arr := j.Get("data").GetIndex(0).MustMap()
	// brr := j.Get("data").GetIndex(1).MustMap()
	// fmt.Println(arr["a"], arr["b"])
	// fmt.Println(brr["a"], brr["b"])
	var a = "b"
	m := make(map[string]int)
	m[a] = 1
	fmt.Println(m["b"])
}
