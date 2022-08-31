package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Test message")
	fmt.Println(time.Now())
	fmt.Println("random num", rand.Intn(12))
	fmt.Println(math.Sqrt(36))
	fmt.Println("Pi", math.Pi)
}
