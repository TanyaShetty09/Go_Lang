package main

import "fmt"

func main() {

	sum := 0
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		sum = sum + i*i

	}
	fmt.Println(sum)
}
