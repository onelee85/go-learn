package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("D:/learn/data.txt", d1, 0644)
	check(err)

	f, err := os.Create("D:/learn/data2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes \n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n", n4)
	w.Flush()
}

func readFile() {
	dat, err := ioutil.ReadFile("D:/learn/data.txt")
	check(err)
	p(string(dat))

	f, err := os.Open("D:/learn/data.txt")
	check(err)

	b1 := make([]byte, 3)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)

	fmt.Printf("5 bytes: %s\n", string(b4))
}

func main() {
	writeFile()
	readFile()
}
