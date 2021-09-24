// Add(), Remove(), Get()
package main

import (
	"fmt"
	"strings"
)

type node struct { //node wrapper
	value int
	next *node
}

func (n node) String() string {
	return fmt.Sprintf("%d ", n.value)
}

type linkedlist struct { //linkedlist wrapper
	head *node
	len int
}

func (l *linkedlist) add(value int) { // modifies the linkedlist by adding node
	newNode := new(node)
	newNode.value = value
	
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		
		for ; iterator.next != nil ; iterator = iterator.next {}
		
		iterator.next = newNode
	}
	
	l.len++
}

func (l *linkedlist) remove(value int) { // modifies the linkedlist by removing node
	var prev *node
	
	for cur := l.head; cur != nil; cur = cur.next { 
		if cur.value == value {
			l.len--
			if l.head == cur {
				l.head = cur.next
			}else {
				prev.next = cur.next
				return
			}
		}
		prev = cur
	}
	
}

func (l linkedlist) get(value int) *node{ // doesn't modify LL, return the node value
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedlist) String() string { // doesn't modify LL, return the string 
sb := strings.Builder{}

	for iterator:=l.head; iterator != nil; iterator = iterator.next {
		
		sb.WriteString(fmt.Sprintf("%s", iterator))
		
	}
	return sb.String()
}
	


func main() {
	fmt.Println()
	
	l := linkedlist{}
	l.add(12)
	l.add(1234)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	fmt.Println(l)
	fmt.Println(l.get(12))
	l.remove(3)
	l.remove(5)
	l.remove(12)
	
	fmt.Println(l)
	fmt.Println(l.len)
}