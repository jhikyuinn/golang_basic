package main

import "fmt"

func main() {
	var dan int
	var i int
	fmt.Scanf("%d", &dan)

	for i = 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}
}
