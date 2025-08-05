package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"sort"
)

type Row []int8
type Field []Row
type Rectangle struct {
	row1 int
	row2 int
	col1 int
	col2 int
}

const (
	point   int8 = 0
	star    int8 = 1
	visited int8 = 2
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	debug.SetMemoryLimit(200 * 1 << 20) // 2750 MB

	var fieldCounter int
	fmt.Fscanln(in, &fieldCounter)

	for i := 0; i < fieldCounter; i++ {
		runSet(in, out)
	}

}

func runSet(in *bufio.Reader, out *bufio.Writer) {

	var height int
	var length int
	fmt.Fscanln(in, &height, &length)
	field := makeField(height, length)

	for r := 0; r < height; r++ {
		var line string
		fmt.Fscanln(in, &line)
		for c := 0; c < length; c++ {
			if line[c:c+1] == "*" {
				field[r][c] = star
			} else {
				field[r][c] = point
			}
		}
	}

	res := analizeField(field)
	for _, r := range res {
		fmt.Fprintf(out, "%d ", r)
	}
	fmt.Fprintln(out)

}

func makeField(height int, lenght int) Field {
	field := make(Field, height)
	for i := range field {
		field[i] = make(Row, lenght)
	}
	return field
}

func analizeField(field Field) []int {

	res := findRectangles(field, 0, 0, len(field)-3, 0, len(field[0])-3)

	sort.Ints(res)

	return res

}

func findRectangles(field Field, level int, rowFrom int, rowTo, colFrom int, colTo int) []int {

	res := []int{}

	var rectangle Rectangle = Rectangle{}

	for row1 := rowFrom; row1 <= rowTo; row1++ {
		for col1 := colFrom; col1 <= colTo; col1++ {

			if field[row1][col1] == visited {
				continue
			}

			if field[row1][col1] == star &&
				field[row1+1][col1] == star &&
				field[row1][col1+1] == star {

				field[row1][col1] = visited

				rectangle.row1 = row1
				rectangle.col1 = col1

				for col2 := rectangle.col1 + 2; ; col2++ {
					if field[rectangle.row1][col2] == star &&
						field[rectangle.row1][col2-1] == star &&
						field[rectangle.row1+1][col2] == star {
						rectangle.col2 = col2
						break
					}
				}

				for row2 := rectangle.row1 + 2; ; row2++ {
					if field[row2][rectangle.col2] == star &&
						field[row2-1][rectangle.col2] == star &&
						field[row2][rectangle.col2-1] == star {
						rectangle.row2 = row2
						break
					}
				}

				res = append(res, level)

				if rectangle.row2-rectangle.row1 >= 6 &&
					rectangle.col2-rectangle.col1 >= 6 {
					subRes := findRectangles(field, level+1, rectangle.row1, rectangle.row2-3, rectangle.col1, rectangle.col2-3)
					res = append(res, subRes...)
				}

				setVisited(field, rectangle.row1, rectangle.row2, rectangle.col1, rectangle.col2)

			}

		}

	}

	return res

}

func setVisited(field Field, rowFrom, rowTo, colFrom, colTo int) {
	for row := rowFrom; row <= rowTo; row++ {
		for col := colFrom; col <= colTo; col++ {
			field[row][col] = visited
		}
	}

}
