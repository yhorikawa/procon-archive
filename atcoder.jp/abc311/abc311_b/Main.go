package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	sc.Split(bufio.ScanWords) // スペース区切りの設定
	N := nextInt()
	D := nextInt()
	var s [][]string
	for i := 0; i < N; i++ {
		input := strings.Split(nextString(), "")
		s = append(s, input)
	}

	S := array_transpose(s)

	ans := 0
	count := 0
	var flag bool
	for d := 0; d < D; d++ {
		flag = true
		for n := 0; n < N; n++ {
			if S[d][n] == "x" {
				flag = false
				break
			}
		}
		if flag {
			count++
			if count > ans {
				ans = count
			}
		}
		if !flag {
			if count > ans {
				ans = count
			}
			count = 0
		}
	}
	fmt.Fprintln(out, ans)
}

func array_transpose(array [][]string) [][]string {
	rowLen := len(array)
	colLen := len(array[0])
	result := make([][]string, colLen)
	for i := range result {
		result[i] = make([]string, rowLen)
	}

	for i, row := range array {
		for j, val := range row {
			result[j][i] = val
		}
	}

	return result
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
