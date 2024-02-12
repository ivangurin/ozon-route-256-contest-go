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

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {

		var qty int
		var percom int
		fmt.Fscanln(in, &qty, &percom)

		res := float64(0)
		for j := 0; j < qty; j++ {
			var value float64
			fmt.Fscanln(in, &value)
			comis := value / 100 * float64(percom)
			comis = math.Round(comis*100) / 100
			res += (comis - float64(int(comis)))
		}

		fmt.Fprintf(out, "%.2f\n", res)

	}

}
