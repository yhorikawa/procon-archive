package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	// sc.Split(bufio.ScanWords) // スペース区切りの設定
	N := nextInt()
	d := make([]int, N)
	count := 0
	m := make(map[int]bool)

	for i := 0; i < N; i++ {
		d[i] = nextInt()
		if !m[d[i]] {
			m[d[i]] = true
			count++
		}
	}

	fmt.Fprintln(out, count)
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
