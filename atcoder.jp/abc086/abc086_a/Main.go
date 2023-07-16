package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if (a*b)%2 == 0 {
		fmt.Println("Even")
		return
	}

	fmt.Println("Odd")
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func max(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}
