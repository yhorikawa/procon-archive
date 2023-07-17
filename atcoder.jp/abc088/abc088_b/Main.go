package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	_ = nextInt()
	str := nextString()
	inputs := strings.Split(str, " ")
	var a []int
	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		a = append(a, p)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	var alice, bob int
	for i, v := range a {
		if i%2 == 0 {
			alice += v
			continue
		}

		bob += v
	}
	fmt.Println(alice - bob)
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
