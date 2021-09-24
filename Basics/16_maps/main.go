package main

import "fmt"

func main() {
	var m map[string]string = map[string]string{} //map[]
	var n map[string]string; //map[]

	fmt.Println(m==nil) //false
	
	n = make(map[string]string)
	
	fmt.Println(n, len(n)) //map[] 0
	m = map[string]string{"Zero": "1"} 
	m["One"] = "1"
	m["Two"] = "2"
	m["One"] = "2" //Override existing value
	
	// Print the map
	for k,v := range m {
		fmt.Println(k, v)
	}
	
	m = map[string]string{"Nil" : "nil"} //override map
	fmt.Println(m)
	
	// Check for key existance
	number, value := m["One"]
	
	if value {
		fmt.Println(number, value)
	} else {
		fmt.Println(number + " Does not exist")
	}
	
	
}