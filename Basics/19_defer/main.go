package main

import "fmt"

func retur() string{
	defer fmt.Println("Deferred")
	return "Hello World"
}


func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	
	fmt.Println(retur())
}

/* Output

hello
Deferred
Hello World
world

*/