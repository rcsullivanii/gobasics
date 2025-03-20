package main

import (
	"fmt"
	"time"
)

var res int = 1   

func power(n, p int) {
	res = res * n
    if p == 1 {
		return
	}
	go power(n, p-1)
}

func main() {
    n := 2
    p := 50
    power(n, p)     
	time.Sleep(100 * time.Millisecond) 
	fmt.Println("Result:", res)
}