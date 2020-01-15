package main

import (
	"fmt"
	"reflect"
)

func main() {
	var person person
	//testType(person)
	testValue(person)
}

type person struct {
	Name string
	age  int64
}

func testType(i interface{}) {
	v := reflect.TypeOf(i) //类型的种类
	fmt.Println(v)
	fmt.Printf("name = %s, kind = %s", v.Name(), v.Kind())
	//main.person
	//name = person, kind = struct

}

func testValue(i interface{}){
	v := reflect.ValueOf(i)
	fmt.Printf("kind = %s", v.Kind())
}
