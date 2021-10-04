package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// fileReader()
	// stringReader()
	connReader()
}

func fileReader() {
	// open file to read
	f, err := os.Open("document.txt")
	
	if err != nil {
		panic(err) // terminate program incase of error
	}
	defer f.Close() // close file at end of this function
	
	readerToStdout(f, 10)
	// buf := make([]byte, 10) // buffer to hold data while reading as stream
	
	// // loop file till EOF
	// for {
	// 	n, err := f.Read(buf) // read file 'f' - read into 'buf'
		
	// 	// check for end of file (no data to read)
	// 	if err == io.EOF {
	// 		break
	// 	}
		
	// 	// misc error
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
		
	// 	// no error
	// 	if n > 0 {
	// 		fmt.Println(string(buf[:n])) // slice the buffer to avoid reading old values from buffer
	// 	}
		
	// }
}

func stringReader() {
	s := strings.NewReader("this is a string")
	// buf := make([]byte, 5)
	// for {
	// 	n, err := s.Read(buf)
		
	// 	if err == io.EOF {
	// 		break
	// 	}
		
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
		
	// 	if n>0 {
	// 		fmt.Println(string(buf[:n]))
	// 	}
	// }
	readerToStdout(s, 10)
}

func connReader() {
	conn, err := net.Dial("tcp", "example.com:80") // (network, address string)
	/* 
	If network is tcp or udp, then address string should be "host:port"
	*/
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // close connection
	
	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")
	// buf := make([]byte, 50)
	
	// for {
	// 	n, err := conn.Read(buf)
		
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
		
	// 	if n>0 {
	// 		fmt.Println(string(buf[:n]))
	// 	}
	// }
	
	readerToStdout(conn, 50)
}

func readerToStdout(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	
	for {
		n, err := r.Read(buf)
		
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		
		if n>0 {
			fmt.Println(string(buf[:n]))
		}
	}
}