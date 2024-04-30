package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {

		var n int
		fmt.Fscanln(in, &n)

		var events string
		fmt.Fscanln(in, &events)

		if !isEventsCorrect(events) {
			fmt.Fprintln(out, "No")
			continue
		}

		fmt.Fprintln(out, "Yes")
	}
}

func isEventsCorrect(events string) bool {

	var x, y, xy, xyxz int
	for i := 0; i < len(events); i++ {

		switch string(events[i]) {
		case "X":
			x++
		case "Y":

			if x > 0 {
				x--
				xy++
				continue
			}

			if xyxz > 0 {
				xyxz--
				xy++
				x++
				continue
			}

			y++

		case "Z":

			if y > 0 {
				y--
				continue
			}

			if x > 0 {
				x--
				if xy > 0 {
					xyxz++
					xy--
				}
				continue
			}

			if xy > 0 {
				xy--
				x++
				continue
			}

			if xyxz > 0 {
				xyxz--
				x++
				continue
			}

			return false

		}

	}

	if x == 0 && y == 0 {
		return true
	}

	return false
}
