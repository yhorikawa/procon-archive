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

	// sc.Split(bufio.ScanWords) // スペース区切りの設定
	_ = nextInt()
	S := nextString()
	array := strings.Split(S, "")
	var a, b, c bool
	for i, v := range array {
		if v == "A" {
			a = true
		}

		if v == "B" {
			b = true
		}

		if v == "C" {
			c = true
		}

		if a && b && c {
			fmt.Fprintln(out, i+1)
			break
		}
	}
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
