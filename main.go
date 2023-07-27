package main

import (
	"fmt"
	"gomap/lib"
)

func main() {
	mapTest := new(lib.Set[string]).New()
	mapTest.Add("foo bar")
	mapTest.Add("bar")
	fmt.Println(mapTest.In("foo bar"))

	fmt.Println(mapTest.In("bar"))
	fmt.Println(mapTest.Len())

}
