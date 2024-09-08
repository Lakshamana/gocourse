package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// write in file
	// size, err := f.WriteString("Hello World!")
	size, err := f.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saving file successfully with %dB\n", size)
	f.Close()

	// read file
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// read as chunks
	f2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	// init reader cursor
	reader := bufio.NewReader(f2)
	buf := make([]byte, 4) // create buffer

	for {
		n, err := reader.Read(buf)  // read size of buffer
		if err != nil {  // error or EOF
			break
		}

		fmt.Println(string(buf[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
