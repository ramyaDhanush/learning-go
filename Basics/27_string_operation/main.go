package main

import (
"fmt"
"strings"
"strconv"
)

func main() {
	dis := "4.7|5.1|5.6|6.2|6.8|7.5|8.2|9.1|10|11|12|13|15|16|18|20"
	// dis := "256|512|768|1024"
	// unit := "ms"
	// unit := "µs"
	unit := "kΩ"
	// unit:="V"
	val := "0|1|2|3|4|5|6|7|8|9|10|11|12|13|14|15"
	// val := "0|1|2|3|4|5|6|7"
	


	res1 := strings.Split(dis, "|")
	res2 := strings.Split(val, "|")
	
	ints := make([]int, len(res2))

	for i, s := range res2 {
	    ints[i], _ = strconv.Atoi(s)
	}
	
	type objs struct {
		display string
		value int
	}
	
	final :=make([]objs, len(ints))
	
	for i := 0; i < len(res1); i++ {
		final[i] = objs{ fmt.Sprintf("%s %s", res1[i], unit), ints[i]}
	}
	
	for _, i := range final {
		fmt.Printf("{ \"display\" : \"%s\", \"value\" : %d },\n", i.display, i.value)
	}
	
	var f [100]int
	u := "dec"
	for idx, _ := range f {
		fmt.Printf("{ \"display\" : \"%d %s\", \"value\" : %d },\n", idx, u, idx)
	}
}