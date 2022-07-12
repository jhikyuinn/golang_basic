package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}

	a[1] = 0

	fmt.Println(a)

	var b []int

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), " Nil Slice입니다.")
	}
}
