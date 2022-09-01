package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println(len("Test Message"))
	fmt.Println("Test Message"[3])
	fmt.Println("Test Message" + " Another message")

	fmt.Println(add(34, 45))

}
