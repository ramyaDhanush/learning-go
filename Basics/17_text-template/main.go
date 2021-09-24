package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name, Secret string
}

func main() {
	secret := secret{"Professor", "I'm from Money Heist"}

	templatePath := "C:/Users/ramya.dhanushkodi/go/src/github.com/ramyaDhanush/learning-go/Basics/17_text-template/template.txt"
	
	t, err := template.New("template.txt").ParseFiles(templatePath)
	
	if err != nil {
		fmt.Println(err)
	} 
	
	err = t.Execute(os.Stdout, secret)
	
	if err != nil {
		fmt.Println(err)
	} 
	
	fmt.Println()
}