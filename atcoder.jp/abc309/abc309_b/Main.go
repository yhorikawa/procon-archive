package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var t [][]string

	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		row := strings.Split(input, "")
		t = append(t, row)
	}

	var result [][]string
	for i := 0; i < n; i++ {
		if i == 0 {
			row := append([]string{t[i+1][0]}, t[i]...)
			row = row[:len(row)-1]
			result = append(result, row)
			continue
		}

		if i == n-1 {
			row := append(t[i], t[i-1][n-1])
			row = row[1:]
			result = append(result, row)
			continue
		}

		row := append([]string{}, t[i]...)
		row[0] = t[i+1][0]
		row[n-1] = t[i-1][n-1]
		result = append(result, row)
	}

	for _, v := range result {
		fmt.Println(strings.Join(v, ""))
	}
}
