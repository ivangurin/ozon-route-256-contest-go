package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row int
	col int
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

		a := coord{}
		b := coord{}

		var rows, cols int
		fmt.Fscanln(in, &rows, &cols)

		matrix := make([][]string, rows)
		for i := range matrix {
			matrix[i] = make([]string, cols)
		}

		var line string
		for row := 0; row < rows; row++ {
			fmt.Fscanln(in, &line)
			for col := 0; col < cols; col++ {
				matrix[row][col] = string(line[col])
				if matrix[row][col] == "A" {
					a.row = row
					a.col = col
				}
				if matrix[row][col] == "B" {
					b.row = row
					b.col = col
				}
			}
		}

		if a.row == 0 && a.col == 0 {
			goDownRight(matrix, b, "b")
		} else if a.row == len(matrix)-1 && a.col == len(matrix[0])-1 {
			goUpLeft(matrix, b, "b")
		} else if b.row == 0 && b.col == 0 {
			goDownRight(matrix, a, "a")
		} else if b.row == len(matrix)-1 && b.col == len(matrix[0])-1 {
			goUpLeft(matrix, a, "a")
		} else if a.row == b.row {
			if a.col < b.col {
				goUpLeft(matrix, a, "a")
				goDownRight(matrix, b, "b")
			} else {
				goUpLeft(matrix, b, "b")
				goDownRight(matrix, a, "a")
			}
		} else if a.row <= b.row {
			goUpLeft(matrix, a, "a")
			goDownRight(matrix, b, "b")
		} else {
			goUpLeft(matrix, b, "b")
			goDownRight(matrix, a, "a")
		}

		// for i := range matrix {
		// 	fmt.Println(matrix[i])
		// }
		// fmt.Println()

		fmt.Fprint(out, matrixToStr(matrix))

	}

}

func goUpLeft(matrix [][]string, pos coord, mark string) {
	if pos.row > 0 {
		if matrix[pos.row-1][pos.col] == "#" {
			pos.col--
			matrix[pos.row][pos.col] = mark
		}
	}
	pos = goUp(matrix, pos, mark)
	pos = goLeft(matrix, pos, mark)
}

func goUp(matrix [][]string, pos coord, mark string) coord {
	for i := pos.row - 1; i >= 0; i-- {
		pos.row = i
		matrix[pos.row][pos.col] = mark
	}
	return pos
}

func goLeft(matrix [][]string, pos coord, mark string) coord {
	for i := pos.col - 1; i >= 0; i-- {
		pos.col = i
		matrix[pos.row][pos.col] = mark
	}
	return pos
}

func goDownRight(matrix [][]string, pos coord, mark string) {
	if pos.row < len(matrix)-2 {
		if matrix[pos.row+1][pos.col] == "#" {
			pos.col++
			matrix[pos.row][pos.col] = mark
		}
	}
	pos = goDown(matrix, pos, mark)
	pos = goRight(matrix, pos, mark)
}

func goDown(matrix [][]string, pos coord, mark string) coord {
	for i := pos.row + 1; i < len(matrix); i++ {
		pos.row = i
		matrix[pos.row][pos.col] = mark
	}
	return pos
}

func goRight(matrix [][]string, pos coord, mark string) coord {
	for i := pos.col + 1; i < len(matrix[0]); i++ {
		pos.col = i
		matrix[pos.row][pos.col] = mark
	}
	return pos
}

func matrixToStr(m [][]string) string {
	sb := strings.Builder{}
	for _, row := range m {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}
