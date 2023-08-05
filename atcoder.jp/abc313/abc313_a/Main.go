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

	//sc.Split(bufio.ScanWords) // スペース区切りの設定
	_ = nextInt()
	p := nextString()
	P := strings.Split(p, " ")
	var pnums []int
	for _, str := range P {
		num, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		pnums = append(pnums, num)
	}
	if len(pnums) == 1 {
		fmt.Fprintln(out, 0)
		return
	}
	arr := pnums[1:]
	maxp := max(arr...)
	x := maxp - pnums[0]
	if x < 0 {
		fmt.Fprintln(out, 0)
		return
	}
	fmt.Fprintln(out, x+1)
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
