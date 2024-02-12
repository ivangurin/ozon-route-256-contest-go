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

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {

		var m int
		fmt.Fscanln(in, &m)

		var q string
		fmt.Fscanln(in, &q)

		fmt.Fprintln(out, checkQ(q))

	}

}

func checkQ(q string) string {

	x, y, xy, xz, yz, zu, zuy := 0, 0, 0, 0, 0, 0, 0

	for i := 0; i < len(q); i++ {

		s := string(q[i])

		switch s {
		case "X":
			x++
		case "Y":
			if x > 0 {
				x--
				xy++
			} else {
				y++
			}
		case "Z":
			if y > 0 {
				y--
				yz++
			} else if x > 0 {
				x--
				xz++
				zu = xy
			} else if xy > 0 {
				xy--
				yz++
				x++
				zuy++
			} else {
				return "No"
			}
		}

	}

	// fmt.Printf("No, x: %d, y: %d, xy: %d, xz: %d, yz: %d, zu: %d\n", x, y, xy, xz, yz, zu)

	// if x > 0 && x <= zuy {
	// 	y = y + x
	// 	x = 0
	// }

	if x == 0 && y <= 2*zu && y <= 2*xz {
		return "Yes"
	}

	return "No"

}
