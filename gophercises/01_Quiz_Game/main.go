package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Package flag implements command-line flag parsing. Read mode https://pkg.go.dev/flag

/* Usage
Define flags using flag.String(), Bool(), Int(), etc.
*/

func main() {

	// flag input : csv
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question, answer'")
	
	fmt.Println(csvFilename, *csvFilename)
	
	// flag input : timeLimit
	timeLimit := flag.Int("limit", 1, "Time limit for quiz in seconds")
	
	// to parse the command line into the defined flags. 
	flag.Parse()
	
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
	
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	problemLoop:
		for i, p := range problems {
			fmt.Printf("Problem %d: %s = \n", i+1, p.q)
			
			answerChannel := make(chan string)
			
			// goroutine to get user input for a question
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerChannel <- answer 
			}()
			
			// Check for timer overflow 
			select {
			case <-timer.C:
				fmt.Printf("You scored %d out of %d problems", correct, len(problems))
				break problemLoop 
			case answer := <-answerChannel:
				if answer == p.a {
					correct++
				}
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
5. Compare answers with csv file answer based on timer 

*/