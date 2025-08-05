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
		var w, h int
		fmt.Fscanln(in, &w, &h)

		field := MakeField(w, h)
		DrawHexagon(field, w, h)
		OutputField(field, out)
	}
}

func MakeField(w, h int) [][]string {
	field := make([][]string, 2*h+1)
	for r := range field {
		field[r] = make([]string, w+2*h)
	}
	return field
}

func DrawHexagon(field [][]string, w, h int) {
	for c := 0; c < h; c++ {
		field[0][c] = " "
	}
	for c := h; c < h+w; c++ {
		field[0][c] = "_"
	}

	for r := 1; r <= h; r++ {
		for c := 0; c < h; c++ {
			if h-c == r {
				field[r][c] = "/"
				field[h+h-r+1][c] = "\\"
				continue
			}
			field[r][c] = " "
			field[h+h-r+1][c] = " "
		}
		for c := h; c < h+w; c++ {
			field[r][c] = " "
			field[h+h-r+1][c] = " "
		}
		for c := h + w; c < h+w+h; c++ {
			if c == h+w+r-1 {
				field[r][c] = "\\"
				field[h+h-r+1][c] = "/"
				continue
			}
			if c < h+w+r-1 {
				field[r][c] = " "
				field[h+h-r+1][c] = " "
			}
		}
	}

	for c := h; c < h+w; c++ {
		field[len(field)-1][c] = "_"
	}
}

func OutputField(field [][]string, out *bufio.Writer) {
	for _, row := range field {
		str := ""
		for _, val := range row {
			str += val
		}
		fmt.Fprintln(out, str)
	}
}
