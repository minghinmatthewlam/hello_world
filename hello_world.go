package main

import (
	"bytes"
	"fmt"
	"net/http"
	"rsc.io/quote"
	"time"
)

type CustomSlice []int

type key int

const (
	cache key = iota
	global
	isEnded
)

func main() {
	fmt.Printf("Hello world...My name is %v\n", quote.Go())

	firstChan := make(chan string)
	secondChan := make(chan string)
	exit := make(chan string)

	go func() {
		t := time.NewTicker(1 * time.Second)
		for now := range t.C {
			firstChan <- fmt.Sprintf("one at time %v", now)
		}
	}()
	go func() {
		t := time.NewTicker(2 * time.Second)
		for now := range t.C {
			secondChan <- fmt.Sprintf("two at time %v", now)
		}
	}()

	go func() {
		t := time.NewTimer(10 * time.Second)
		exit <- time.Time.String(<-t.C)
	}()

	for {
		select {
		case msg1 := <-firstChan:
			fmt.Println(msg1)
		case msg2 := <-secondChan:
			fmt.Println(msg2)
		case <-exit:
			fmt.Println("exiting...")
			break
		}

	}
	fmt.Println("Finished test")
}

func helloResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request with context: ", r.Context())
	fmt.Fprintf(w, "hello world")
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
