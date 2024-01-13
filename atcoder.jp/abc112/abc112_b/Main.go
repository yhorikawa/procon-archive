package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	var sc = bufio.NewScanner(r)
	var out = bufio.NewWriter(w)

	defer out.Flush()

	sc.Split(bufio.ScanWords) // スペース区切りの設定
	N := nextInt(sc)
	T := nextInt(sc)

	result := 10000
	for i := 0; i < N; i++ {
		c := nextInt(sc)
		t := nextInt(sc)
		if t <= T {
			result = min(result, c)
		}
	}

	if result == 10000 {
		fmt.Fprintln(out, "TLE")
		return
	}

	fmt.Fprintln(out, result)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString(sc *bufio.Scanner) string {
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
