package main

import (
	"bufio"
	"fmt"
	"os"
)

type season struct {
	from int
	to   int
	len  int
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {

		var l int
		fmt.Fscanln(in, &l)

		a := make([]int, 0, l)

		for j := 0; j < l; j++ {

			var v int
			if j == l-1 {
				fmt.Fscanln(in, &v)
			} else {
				fmt.Fscan(in, &v)
			}

			a = append(a, v)

		}

		seasons := getSeasons(a)

		seassonCounter := map[int]int{}
		for _, s := range seasons {

			series := getSeries(seasons, s.from, s.to) + 1

			counter := seassonCounter[s.len]
			if series > counter {
				seassonCounter[s.len] = series
			}
			// fmt.Printf("%+v, %d\n", s, series)

		}

		res := make([]int, l)
		for j := 0; j < l; j++ {
			res[j] = seassonCounter[j+1]
		}

		for j, r := range res {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			if j < l-1 {
				fmt.Fprint(out, r)
			} else {
				fmt.Fprintln(out, r)
			}
		}

	}

}

func getSeries(seasons []*season, from int, to int) int {

	res := 0

	for _, s := range seasons {
		if s.from == to {
			res++
			subres := getSeries(seasons, s.from, s.to)
			res += subres
		}
	}

	return res

}

func getSeasons(a []int) []*season {

	res := []*season{}

	if len(a) < 3 {
		return res
	}

	for i := 0; i < len(a)-2; i++ {

		from := i
		to := -1

		lenUp := 0
		for j := i; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				lenUp++
			} else {
				break
			}
		}

		if lenUp == 0 {
			continue
		}

		lenDown := 0
		for j := i + lenUp; j < i+lenUp*2; j++ {
			if j > len(a)-2 {
				break
			}
			if a[j] > a[j+1] {
				lenDown++
				to = j + 1
			} else {
				break
			}
		}

		if lenUp != lenDown {
			continue
		}

		res = append(res, &season{from: from, to: to, len: (to - from) / 2})

	}

	return res

}
