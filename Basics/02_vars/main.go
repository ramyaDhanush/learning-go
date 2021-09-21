package main

import "fmt"
var name = "Brad"
var age int32= 37
var isCool = true

func main() {
	name, email := "Ramya"	, "ramya@ramya.com"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isCool)
	
	fmt.Printf("%T - %T\n%T\n", age, name, isCool)
	fmt.Println(name," ", email)

	var a = imp
	a()
}

func imp() {
	fmt.Println("Im in imp")
}

