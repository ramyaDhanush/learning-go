package main

import "fmt"

func greeting(name string) string {
	return "Hello!! " + name
}

func add(num1, num2 int) int {
	return num1 + num2
}

func variadic(integers ...int) {
	for _, integer := range integers {
		fmt.Println(integer)
	}
}

func main() {
	fmt.Println(greeting("User"))
	fmt.Println(add(1, 3))
	variadic()
	variadic(1, 23,4)
	// variadic({1, 23,4}) - error

}