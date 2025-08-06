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
	var fw, fh, hw, hh, t int
	fmt.Fscanln(in, &fw, &fh, &hw, &hh, &t)

	field := MakeField(fw, fh)
	DrawHexagons(field, hw, hh, t)
	PrintField(field, out)
}

func MakeField(w, h int) [][]string {
	field := make([][]string, h+2)
	for r := range field {
		field[r] = make([]string, w+2)
		if r == 0 || r == h+1 {
			field[r][0] = "+"
			field[r][len(field[r])-1] = "+"
			for c := 1; c < w+1; c++ {
				field[r][c] = "-"
			}
			continue
		}
		for c := 1; c < w+1; c++ {
			field[r][c] = " "
		}
		field[r][0] = "|"
		field[r][len(field[r])-1] = "|"
	}
	return field
}

func DrawHexagons(field [][]string, hw, hh, t int) {
	var oddLine int
	var evenLine int
	var rowStart int
	var colStart int
	var hexagons int
	for {
		fieldLen := len(field[0]) - 2

		hexagons = 0
		if (oddLine+evenLine)%2 == 0 {
			oddLine++

			if fieldLen >= hh+hw+hh {
				hexagons++
				fieldLen -= hh + hw + hh
			}
			hexagons += fieldLen / (hh + hw + hh + hw)

			rowStart = 1 + (oddLine-1)*(hh+hh)
			colStart = 1
		} else {
			evenLine++

			fieldLen -= hh
			hexagons += fieldLen / (hw + hh + hw + hh)

			rowStart = 1 + hh + (evenLine-1)*(hh+hh)
			colStart = hh + hw + 1
		}

		for i := 0; i < hexagons; i++ {
			DrawHexagon(field, rowStart, colStart, hw, hh)

			t--
			if t == 0 {
				break
			}

			colStart += hh + hw + hh + hw
		}
		if t == 0 {
			break
		}
	}
}

func DrawHexagon(field [][]string, r, c, w, h int) {
	for col := c + h; col < c+h+w; col++ {
		field[r][col] = "_"
	}

	for row := r + 1; row <= r+h; row++ {
		for col := c; col < c+h; col++ {
			if (h - (col - c)) == (row - r) {
				field[row][col] = "/"
				field[r+1+h+(r+1+h-row)-1][col] = "\\"
			}
		}
		for col := c + h + w; col < c+h+w+h; col++ {
			if col-c == row-r+h+w-1 {
				field[row][col] = "\\"
				field[r+1+h+(r+1+h-row)-1][col] = "/"
				continue
			}
		}
	}

	for col := c + h; col < c+h+w; col++ {
		field[r+h+h][col] = "_"
	}
}

func PrintField(field [][]string, out *bufio.Writer) {
	for _, row := range field {
		str := ""
		for _, val := range row {
			str += val
		}
		fmt.Fprintln(out, str)
	}
}
