package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	var sc = bufio.NewScanner(r)
	var out = bufio.NewWriter(w)

	defer out.Flush()

	sc.Split(bufio.ScanWords) // スペース区切りの設定
	N, M := nextInt(sc), nextInt(sc)
	s := make([]int, N)

	for i := 0; i < M; i++ {
		si, ci := nextInt(sc), nextInt(sc)
		// １桁目が0は許容しない
		if N != 1 && si == 1 && ci == 0 {
			fmt.Fprintln(out, -1)
			return
		}

		// 同じ桁に異なる数字が入ることは許容しない
		if s[si-1] != 0 && s[si-1] != ci {
			fmt.Fprintln(out, -1)
			return
		}

		s[si-1] = ci
	}

	if N == 1 {
		if s[0] == 0 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, s[0])
		}
		return
	}

	if s[0] == 0 {
		s[0] = 1
	}

	str := []string{}
	for _, v := range s {
		str = append(str, strconv.Itoa(v))
	}
	fmt.Fprintln(out, strings.Join(str, ""))
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
