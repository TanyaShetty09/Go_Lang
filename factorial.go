package main

import "fmt"

func main() {

	fact := 1
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	fmt.Println(fact)
}
