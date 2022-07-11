package main

import "fmt"

func main() {
	var n, i, j int

	fmt.Scanf("%d", &n)

	for i = n; i > 0; i-- {
		for j = 0; j < n-i; j++ {
			fmt.Print("o ")
		}
		fmt.Println("* ")
	}
}
