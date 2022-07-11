package main

import "fmt"

func main() {

	var n int

	for n = 1; n < 100; n++ {
		if n%7 == 0 {
			fmt.Printf("%d ", n)
		} else if n%9 == 0 {
			fmt.Printf("%d ", n)
		}
	}
}
