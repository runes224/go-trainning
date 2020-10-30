package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

var num1, num2, num3 = 0, 0, 1

func fibonacci() func() int {
	return func() int {
		num1, num2, num3 = num2, num3, num2+num3
		return num1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
