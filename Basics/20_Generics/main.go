/*
Implmentation of Generics
1. Type of parameter
2. Constraints on parameter

Generics Implementation
1. Minimize new concepts 
2. Complexity falls on writer of generic code
3. Writer & User can work independently
4. Short build times & fast execution time
5. Preserve clarity and simplicity of go
*/

package main

import "fmt"

// Generics
// func print[T any](output ...T) {
// 	fmt.Println(output)
// }

func main() {
	fmt.Println("")
	print(1, 2, 3)
	print(1.78, 2.0, 3.12)
	print("1", "2", "3")
}

/*

*/