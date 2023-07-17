package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	sc.Split(bufio.ScanWords) // スペース区切りの設定
	S := nextString()
	if S[0:1] != "A" {
		fmt.Fprintln(out, "WA")
		return
	}

	s := S[2 : len(S)-1]
	count := strings.Count(s, "C")
	if count != 1 {
		fmt.Fprintln(out, "WA")
		return
	}

	c := strings.Index(S, "C")

	for i, v := range S {
		if i == 0 {
			continue
		}

		if i == c {
			continue
		}

		if unicode.IsUpper(v) {
			fmt.Fprintln(out, "WA")
			return
		}
	}
	fmt.Fprintln(out, "AC")
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
