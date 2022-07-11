package main

import "fmt"

func main() {
	var i, j int
	var result int

	for i = 0; i < 10; i++ {
		for j = 9; j >= 0; j-- {
			if i != j {
				result = i*10 + j + j*10 + i
				if result == 99 {
					fmt.Printf("%d%d + %d%d = %d\n", i, j, j, i, result)
				}
			}
		}
	}
}
