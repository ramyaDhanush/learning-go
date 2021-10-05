package main

import (
	"bytes"
	"encoding/json"
	"os"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err) // terminate execution on error
	}
	defer f.Close() // close file at end
	
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	
	u := user{Name :"Bob", Age:10}
	if err := enc.Encode(u); err!=nil {
		panic(err)
	}
	
	fmt.Println(buf)
	// fmt.Println(string(buf[:]))
	fmt.Println(u)
	
	f.Write(buf.Bytes())
}