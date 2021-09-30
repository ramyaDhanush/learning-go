// package part1
// package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Package flag implements command-line flag parsing. Read mode https://pkg.go.dev/flag

/* Usage
Define flags using flag.String(), Bool(), Int(), etc.
*/

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question, answer'")
	fmt.Println(csvFilename, *csvFilename)
	// flag.Parse()
	
	file, err := os.Open(*csvFilename)
	
	if err!=nil { // wrond file path
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	
	if err!=nil { 
		exit("Failed to parse the provided CSV file.")
	}
	
	problems := parseLines(lines)
	
	correct := 0
	
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)
		
		var answer string
		fmt.Scanf("%s\n", &answer)
		
		if answer == p.a {
			correct++
			fmt.Println("Correct!!!")
		}
		
	}	
	
	fmt.Printf("You scored %d out of %d problems", correct, len(problems))
}


// define struct to hold quiz data
type problem struct { //question - q, answer - a
	q, a string
}

// parse data from 2D array of strings to struct of problem
func parseLines(lines [][]string) []problem {
	ret :=make([]problem, len(lines))
	
	for i, line := range lines {
		ret[i] = problem{
			q:line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	
	return ret
}

// to terminate execution in terms of error
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

/* 
--------------------
		  Approach
--------------------
1. Get csv file input
2. Read csv file & validate the content
3. Iterate through quiz questions 
4. Get answer from user input
5. Compare answers with csv file answer

*/