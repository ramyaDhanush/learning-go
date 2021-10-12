package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Name    string `json:"user_name"`
	Age     int
	Subject string
}

func main() {
	f, err := os.Open("input.json")
	if err != nil {
		panic(err)
	}
	
	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	u := Student{}
	
	if err := json.Unmarshal(jsonBytes, &u); err != nil {
		panic(err)
	}
	
	fmt.Println(u)
	
}