package main

import "fmt"

func main() {
	var i, j int
	for i = 2; i < 10; i++ {
		if i%2 == 1 {
			for j = 1; j <= i; j++ {
				fmt.Printf("%d x %d = %d\n", i, j, i*j)
			}
			fmt.Println()
		}
	}
}
