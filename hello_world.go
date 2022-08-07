package main

import (
	"bytes"
	"fmt"
	"rsc.io/quote"

	"github.com/minghinmatthewlam/hello_world/utils"
)

type CustomSlice []int

func main() {
	fmt.Printf("Hello world...My name is %v\n", quote.Go())
	testBuffer()
	utils.RandomFunc()
	fmt.Println("Finished test")
}

func testBuffer() {
	b := bytes.NewBuffer([]byte("hi"))
	buffer := make([]byte, 3)
	if n, err := b.Read(buffer); err != nil {
		fmt.Printf("Read %d bytes with err: %e", n, err)
	} else {
		fmt.Println("read", n, "bytes")
		fmt.Println(buffer)
	}

	if n, err := b.Read(buffer); err != nil {
		fmt.Printf("Read %d bytes with err: %e", n, err)
	}
}

func PrintNumbers(txt string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(txt, i)
	}
}

func removeElement(s CustomSlice, d *map[string]int) CustomSlice {
	(*d)["hello"] = 1
	return s[:len(s)-1]
}
