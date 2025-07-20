package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		var a, r, c int
		fmt.Fscanln(in, &a, &r, &c)

		arts := make([][]string, 0, a)
		for j := 0; j < a; j++ {
			if j > 0 {
				fmt.Fscanln(in)
			}

			art := make([]string, 0, r)
			for l := 0; l < r; l++ {
				row, _ := in.ReadString('\n')
				row = strings.Trim(row, "\n")
				art = append(art, row)
			}

			arts = append(arts, art)
		}

		res := drawMountains(arts, r, c)
		for _, row := range res {
			fmt.Fprintln(out, row)
		}
		fmt.Fprintln(out)
	}
}

func drawMountains(arts [][]string, rows, cols int) []string {
	if len(arts) == 1 {
		return arts[0]
	}

	firstArt := arts[0]
	for _, nextArt := range arts[1:] {
		for r := 0; r < rows; r++ {
			firstArtRow := strings.Split(firstArt[r], "")
			nextArtRow := strings.Split(nextArt[r], "")

			for c := 0; c < cols; c++ {
				if firstArtRow[c] == "." {
					firstArtRow[c] = nextArtRow[c]
				}
			}

			firstArt[r] = strings.Join(firstArtRow, "")
		}
	}

	return firstArt
}
