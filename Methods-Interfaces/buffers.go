package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	buffer := bytes.NewBufferString("Hello ")
	fmt.Println(buffer.String())
	buffer.WriteString("World\n")
	fmt.Println(buffer.String())
	buffer.WriteString("Hello Gophers\n")
	fmt.Println(buffer.String())

	buffer.Reset()

	buffer.WriteString("New text added to buffer") // without \n
	fmt.Println(buffer.String())

	buffer.Write([]byte("Nice Test"))
	fmt.Println(buffer.String())

	// implementing writer interface with buffer to write toa  file
	// creating a value of the type person
	p := person{
		first: "Tomzeras",
	}

	// creating the txt file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// creating the buffer
	var b bytes.Buffer

	p.writeOut(f) // writing to the file
	fmt.Println("Write to file")
	p.writeOut(&b) // writing to the buffer
	fmt.Println("Write to buffer")
	fmt.Println(b.String()) // printing from the buffer

	b.WriteString(p.first) // writing to buffer again
	fmt.Println("Write to buffer")
	fmt.Println(b.String())
	f.Write([]byte(b.String())) // writing to file again
}
