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
	S := nextString(sc)
	var t int
	for i := 2; i < N; i++ {
		f := uniqueStrings(S[:i])
		s := uniqueStrings(S[i:])
		var intersection []string
		for _, h := range f {
			for _, f := range s {
				if h == f {
					intersection = append(intersection, f)
				}
			}
		}
		t = max(t, len(intersection))
	}
	fmt.Fprintln(out, t)
}

func uniqueStrings(strs string) []string {
	m := map[string]struct{}{}
	for _, str := range strs {
		m[string(str)] = struct{}{}
	}

	ret := []string{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
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
