package main

import (
	"fmt"
	"time"
)

func power_h(n, p, res int) {
    if p <= 1 {
        fmt.Println("Result:", res*n)
        return
    }
    go power_h(n, p-1, res * n)
}

func power(n, p int) {
	power_h(n, p, 1)
}

func main() {
	n := 2
    p := 50
    go power(n, p)
    time.Sleep(100 * time.Millisecond) // waiting for goroutine to finish
}