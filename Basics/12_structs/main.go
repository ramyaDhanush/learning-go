/*

struct is a collection of fields.
Struct fields are accessed using a dot. 
Struct fields can be accessed through a struct pointer. 
 A struct literal denotes a newly allocated struct value by listing the values of its fields. 
*/

package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")

	fmt.Println(person2.greet())
	

	

	
	// john := student{name: "John"}
	// fmt.Println(john)
	// john = student{subjects: []string{"Maths", "Chemistry", "Physics"}, rollno: 12}
	// fmt.Println(john)
	// john.name = "John Smith"
	// fmt.Println(john)
	// john = student{"Jonny", []string{"History"}, 8}
	// fmt.Println(john)
	// john.rollno++
	// fmt.Println(john)
	
	jane := results{"Jane", school{"GO Univ", "Go Town"}, 14}
	fmt.Println(jane)
	jane.printResult()
}

type student struct {
	name string
	subjects []string
	rollno int
}

type school struct {
	name string
	place string
}

type results struct {
	name string
	education school
	rollno int
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)


func (r results) printResult () {
	fmt.Println(r)
}