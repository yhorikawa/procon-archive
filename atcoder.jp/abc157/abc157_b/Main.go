package main

import (
	"bufio"
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
	bingoCard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		bingoCard[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			bingoCard[i][j] = nextInt(sc)
		}
	}
	N := nextInt(sc)

	for i := 0; i < N; i++ {
		b := nextInt(sc)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if bingoCard[j][k] == b {
					bingoCard[j][k] = 0
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		if bingoCard[i][0] == 0 && bingoCard[i][1] == 0 && bingoCard[i][2] == 0 {
			out.WriteString("Yes\n")
			return
		}
		if bingoCard[0][i] == 0 && bingoCard[1][i] == 0 && bingoCard[2][i] == 0 {
			out.WriteString("Yes\n")
			return
		}
	}

	if bingoCard[0][0] == 0 && bingoCard[1][1] == 0 && bingoCard[2][2] == 0 {
		out.WriteString("Yes\n")
		return
	}

	if bingoCard[0][2] == 0 && bingoCard[1][1] == 0 && bingoCard[2][0] == 0 {
		out.WriteString("Yes\n")
		return
	}

	out.WriteString("No\n")
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
