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

	sc.Split(bufio.ScanWords) // スペース区切りの設定
	N, Y := nextInt(), nextInt()

	for a := 0; a <= N; a++ {
		for b := 0; b <= N-a; b++ {
			c := N - a - b
			if 10000*a+5000*b+1000*c == Y {
				fmt.Fprintln(out, a, b, c)
				return
			}
		}
	}

	fmt.Fprintln(out, -1, -1, -1)
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
