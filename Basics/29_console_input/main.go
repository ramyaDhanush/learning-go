package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Birth Year: ")
	scanner.Scan()
	input1, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Your Name: ")
	scanner.Scan()
	input2 := scanner.Text()
	
	fmt.Printf("Hey %s you are %d years old", input2, 2021-input1)
}