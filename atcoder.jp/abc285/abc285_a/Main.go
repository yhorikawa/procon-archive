package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if 2*a == b || 2*a+1 == b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
