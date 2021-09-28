/*
 An array has a fixed size. 
 A slice, on the other hand, is a dynamically-sized, 
 flexible view into the elements of an array. 
 In practice, slices are much more common than arrays.
 
  The zero value of a slice is nil.
	A nil slice has a length and capacity of 0 and has no underlying array. 
	
	var s[int]
	s == nil is true
*/

package main

import "fmt"

func main() {
	// Arrays - fixed length with type defined 
	// Slices - array with dynamic length with type defined 
	
	var fruitArr[4] string
	var intArr[4] int32
	
	fruitArr[0] = "Apple"
	
	var a = 12.12
	fmt.Println(a)
	intArr[2] = 12
	
	floatArr := [2]float32{1.23, 7.67}
	floatArr64 := [2]float64{1.23, 7.67}

	fmt.Println(fruitArr)
	fmt.Printf("%f",floatArr)
	fmt.Println(floatArr)
	fmt.Printf("%f",floatArr64)
	fmt.Println(floatArr64)
	fmt.Println(intArr)
	
	sliceContainer := make([]int32, 3, 7)
	sliceContainerFloat := make([]float32, 3, 7)
	
	fmt.Println(sliceContainer)
	fmt.Println(sliceContainerFloat)
	
	var floatArr12 []float32 = []float32{1,2 , 4, 5}
	
	fmt.Println(floatArr12)
	floatArr12 = floatArr12[0:2]
	fmt.Println(floatArr12)
	
	emptySlice := []int{}
	fmt.Println(emptySlice, emptySlice == nil)
	
}