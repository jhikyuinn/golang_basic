package main

import "fmt"

func main() {

	var num1, num2, num3 int
	var nnum1 float32
	var nnum2 uint
	var nnum3 int64
	fmt.Scanln(&num1, &num2, &num3)

	nnum1 = float32(num1)
	nnum2 = uint(num2)
	nnum3 = int64(num3)

	fmt.Printf("%T, %f", nnum1, nnum1)
	fmt.Println()
	fmt.Printf("%T, %d", nnum2, nnum2)
	fmt.Println()
	fmt.Printf("%T, %d", nnum3, nnum3)
}
