package main

import (
	"fmt"
)

func main() {
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Printf("mySlice1 %+v\n", mySlice1)
	fmt.Printf("mySlice2 %+v\n", mySlice2)
	fmt.Printf("mySlice3 %+v\n", mySlice3)
	fmt.Printf("mySlice4 %+v\n", mySlice4)

	//myStringSlice1 := make([]string, 0, 2)
	myStringSlice1 := "test for string slice"
	fmt.Println(myStringSlice1[1])
	fmt.Printf("myStringSlice1 %s",myStringSlice1)
}
