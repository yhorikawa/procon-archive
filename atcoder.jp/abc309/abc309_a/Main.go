package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b-a == 1 && a%3 != 0 {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
