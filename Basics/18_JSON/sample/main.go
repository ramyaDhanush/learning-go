package main

import (
	"fmt"
	"encoding/json"
)

type student struct {
	Name string `json:"full_name"`
	State string
	Id int
}

func main() {
	data := `{"full_name": "Genny","state": "TN","id": 12}`
	
	fmt.Println(data)
	
	var john student 
	
	err := json.Unmarshal([]byte(data), &john)
	
	if err!=nil {
		fmt.Println("Sad -- ", err)
	}
	
	fmt.Println(john)
	
	sTo, err := json.Marshal(john)
	
	if err!=nil {
		fmt.Println("Sad -- ", err)
	}
	
	fmt.Println(sTo)
	fmt.Println("Before ", john)
	
	json.Unmarshal(sTo, &john)
	fmt.Println("After ", john)
}