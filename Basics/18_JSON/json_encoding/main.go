package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type City struct {
	Name string `json:"city_name"`
	GDP string `json:"-"`
	Population int `json:"city_pop"`
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	City City `json:"city"`
	CreatedAt customTime `json:"created_at"`
}

type customTime struct {
	time.Time 
}

const layout = "1990-01-01"

func (c customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Format(layout))), nil
}
func (c *customTime) UnMarshalJSON(v []byte) error {
var err error
	c.Time, err = time.Parse(layout, strings.Trim("\"", string(v)))
	
	// strings.
	if err != nil {
		return err
	}
	return nil
}



func main() {
	// t := time.Now()
	u := User{
		Name : "bob",
		Age : 37,
		City : City{
			Name : "London",
			GDP : "1234",
		Population : 2137289,
		},
		CreatedAt: customTime{time.Now()},
	}
	
	out, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(out))
}