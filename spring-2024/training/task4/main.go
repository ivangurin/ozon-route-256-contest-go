package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)

	for ni := 0; ni < n; ni++ {
		var s int
		fmt.Fscanln(in, &s)

		t := make([]int, 0, n)
		for si := 0; si < s; si++ {
			var ti int
			fmt.Fscan(in, &ti)
			t = append(t, ti)
		}
		fmt.Fscanln(in)

		res := calcPlaces(t)

		fmt.Fprintf(out, "%s \n", strings.Trim(fmt.Sprint(res), "[]"))
	}
}

func calcPlaces(t []int) []int {

	tSorted := make([]int, len(t))
	copy(tSorted, t)
	sort.Ints(tSorted)

	places := make(map[int]int)
	for i := 0; i < len(tSorted); i++ {
		if i == 0 {
			places[tSorted[i]] = 1
			continue
		}

		if (tSorted[i] - tSorted[i-1]) <= 1 {
			places[tSorted[i]] = places[tSorted[i-1]]
		} else {
			places[tSorted[i]] = i + 1
		}

	}

	res := make([]int, len(t))
	for i := 0; i < len(t); i++ {
		res[i] = places[t[i]]
	}

	return res
}
