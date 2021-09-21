package main

import "fmt"

func greeting(name string) string {
	return "Hello!! " + name
}

func add(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("User"))
	fmt.Println(add(1, 3))

}