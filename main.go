package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Test message")
	fmt.Println(time.Now())
	fmt.Println("random num", rand.Intn(12))
}
