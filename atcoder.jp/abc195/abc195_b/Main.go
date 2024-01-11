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
	A, B, W := nextInt(sc), nextInt(sc), nextInt(sc)

	W *= 1000
	min := 1000000000
	max := 0
	for i := 1; i <= 1000000; i++ {
		if A*i <= W && W <= B*i {
			min = minInt(min, i)
			max = maxInt(max, i)
		}
	}

	if min == 1000000000 {
		fmt.Fprintln(out, "UNSATISFIABLE")
		return
	}

	fmt.Fprintf(out, "%d %d\n", min, max)
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

func minInt(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func maxInt(values ...int) int {
	ret := values[0]
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}
