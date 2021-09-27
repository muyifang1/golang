package main

import (
	"reflect"
)

type MyType struct {
	Name string `json:"name,this is the name desc"`
	Age string `json:"age,this is the age desc"`
}

func main() {
	mt := MyType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	age := myType.Field(1)
	tag := name.Tag.Get("json")
	println(tag)
	tag = age.Tag.Get("json")
	println(tag)

	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}
