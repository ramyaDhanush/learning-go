package main

import "fmt"

func main() {
	const zero = 0
	const (one = 12 
	two 
	three)
	
	fmt.Println(zero, one, two)
	
	const (number0 int = iota
		number1
		number2
	)
	
	const (_  = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	
	fmt.Println(kb, mb, gb, tb, pb, eb)

}