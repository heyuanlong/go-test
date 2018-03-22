package main

import (
	"fmt"
	"reflect"
)
type IBean interface {
	
}
type testS struct {
}

func test1(bean IBean)  {
	typeOf := reflect.TypeOf(bean)
	typeName := typeOf.String()
	fmt.Println(typeName)
}

func main() {
	var t  *testS
	var t2 int
	var t3 *int
	test1(t)
	test1(t2)
	test1(t3)
}