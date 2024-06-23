package main

import (
	"fmt"
	"time"
)

var b int

func main() {
	go func(){
	time.Sleep(time.Millisecond)
	fmt.Println("competitive with delay")
}()
	

	go fmt.Println("competitive")
	go fmt.Println("competitive")
	time.Sleep(2 * time.Second)
}