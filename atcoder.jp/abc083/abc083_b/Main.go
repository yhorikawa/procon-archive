package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var result int
	for i := 1; i <= n; i++ {
		sum := calc(i)
		if a <= sum && sum <= b {
			result += i
		}
	}
	fmt.Println(result)
}

func calc(n int) int {
	var result int
	for n > 0 {
		result += n % 10
		n /= 10
	}
	return result
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
