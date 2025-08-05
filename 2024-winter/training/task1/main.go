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

	var setsCounter int
	fmt.Fscanln(in, &setsCounter)

	sets := make([]map[int]int, setsCounter)
	for i := 0; i < setsCounter; i++ {
		sets[i] = make(map[int]int, 10)
		for j := 0; j < 10; j++ {
			var size int
			fmt.Fscan(in, &size)
			sets[i][size]++
		}
	}

	for i, set := range sets {
		correct := true
		for key, ship := range set {
			if key+ship != 5 {
				correct = false
				break
			}
		}
		if correct {
			fmt.Fprint(out, "YES")
		} else {
			fmt.Fprint(out, "NO")
		}

		if i < len(sets)-1 {
			fmt.Fprint(out, "\n")
		}
	}

}
