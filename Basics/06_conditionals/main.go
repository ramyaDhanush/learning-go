package main

import (
	"fmt"
	"runtime"
)

func compare(num1, num2 int) {
	if num1>num2 {
		printResult("greater", num1, num2)
	}	else {
		printResult("smaller", num1, num2)
	}
}

func printResult (condition string, num1, num2 int) {
		fmt.Printf("%d is %s than %d\n", num1, condition,  num2)
}

func main() {
	compare(1, 3)
	compare(6, 9)
	compare(30, 12)
	
	venues := []string{"Home", "Work", "School", "Bar", "Gym"}
	
	for _, venue := range venues {
		switch(venue) {
			case "Home":
				fmt.Println("Home")
				// break
			case "Work":
				fmt.Println("Work")
				// break
			case "School":
				fmt.Println("School")
				// break
			case "Bar":
				fmt.Println("Bar")
				// break
			default:
				fmt.Println("Gym")
		}
	}
	
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}