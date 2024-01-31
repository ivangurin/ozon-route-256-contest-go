package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var counter int
	fmt.Fscanln(in, &counter)

	sets := make([][]int, counter)

	for i := 0; i < counter; i++ {

		var l int
		fmt.Fscanln(in, &l)

		sets[i] = make([]int, l)

		for j := 0; j < l; j++ {

			var v int
			if j == l-1 {
				fmt.Fscanln(in, &v)
			} else {
				fmt.Fscan(in, &v)
			}

			sets[i][j] = v

		}

	}

	for _, set := range sets {

		res := compress(set)
		fmt.Fprint(out, len(res))
		fmt.Fprint(out, "\r\n")
		for i, v := range res {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, v)
		}

		fmt.Fprint(out, "\r\n")

	}

}

func compress(set []int) []int {

	res := []int{}

	if len(set) == 1 {
		res = append(res, set[0], 0)
		return res
	}

	cursor := 0
	nums := 0
	currSign := ""
	prevSign := ""

	var ended bool
	var delta int

	for i := range set {

		if i == 0 {
			continue
		}

		ended = false

		delta = set[i] - set[i-1]

		if delta == 1 {
			currSign = "+"
			nums++
		}

		if delta == -1 {
			currSign = "-"
			nums++
		}

		if prevSign == "" {
			prevSign = currSign
		}

		if math.Abs(float64(delta)) != 1 || prevSign != currSign {

			if nums == 0 {
				res = append(res, set[cursor], 0)
			} else {
				res = append(res, set[cursor], set[i-1]-set[cursor])
			}

			cursor = i
			nums = 0
			currSign = ""

			if float64(delta) == 1 && prevSign == currSign {
				ended = true
			}

		}

		prevSign = currSign

	}

	if !ended {

		if nums == 0 {
			res = append(res, set[cursor], 0)
		} else {
			res = append(res, set[cursor], set[len(set)-1]-set[cursor])
		}

	}

	return res

}
