package main

import (
"fmt"
"os"
)

func main() {
	/*
		openFile(filename, flag - int, permission - FileMode)
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC - 577
		os.O_WRONLY - 1
		os.O_CREATE - 64
		os.O_TRUNC - 512
	*/
	
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600) // Write(WR) only, if not exists create
	fmt.Printf("%T - %v\n", f, f)
	fmt.Println(os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.O_WRONLY, os.O_CREATE, os.O_TRUNC)
	if err != nil {
		panic(err)
	}
	
	defer f.Close() //close file at end
	
	f.Write([]byte("Writing data to the file using write method\n"))
	f.Write([]byte("Writing data to the file using write method\n"))
	f.Write([]byte("Writing data to the file using write method\n"))
	f.Write([]byte("Writing data to the file using write method\n"))
	f.Write([]byte("Writing data to the file using write method\n"))
	
	fmt.Println("Done")
}