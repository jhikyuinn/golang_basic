package main

import "fmt"

func main() {
	var num1, num2, q, n int

	fmt.Scanln(&num1, &num2)

	q = num1 / num2
	n = num1 % num2

	fmt.Printf("몫 : %d, 나머지 : %d", q, n)
}
