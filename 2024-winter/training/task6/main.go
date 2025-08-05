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

	var counter int
	fmt.Fscanln(in, &counter)

	sets := make([]string, counter)

	for i := 0; i < counter; i++ {
		var s string
		fmt.Fscanln(in, &s)
		sets[i] = s
	}

	for ind, str := range sets {
		if ind > 0 {
			fmt.Fprintln(out)
		}
		fmt.Fprint(out, eval(str)+"\n-")
	}

}

func eval(s string) string {

	row, col := 0, 0
	field := [][]string{}
	field = append(field, []string{})

	for _, l := range s {

		symb := string(l)

		// fmt.Println(symb, ":", row, col, field)

		switch symb {
		case "L":
			if col > 0 {
				col--
			}
		case "R":
			if col < len(field[row]) {
				col++
			}
		case "U":
			if row > 0 {
				row--
				if col > len(field[row]) {
					col = len(field[row])
				}
			}
		case "D":
			if row < len(field)-1 {
				row++
				if col > len(field[row]) {
					col = len(field[row])
				}
			}
		case "B":
			col = 0
		case "E":
			col = len(field[row])
		case "N":
			newField := [][]string{}
			newField = append(newField, field[:row]...)
			newRow := []string{}
			newRow = append(newRow, field[row][:col]...)
			newField = append(newField, newRow)
			newRow = []string{}
			newRow = append(newRow, field[row][col:]...)
			newField = append(newField, newRow)
			if row < len(field)-1 {
				newField = append(newField, field[row+1:]...)
			}
			field = newField
			row++
			col = 0

		default:
			newRow := []string{}
			if col > 0 {
				newRow = append(newRow, field[row][:col]...)
			}
			newRow = append(newRow, symb)
			newRow = append(newRow, field[row][col:]...)
			field[row] = newRow
			col++
		}

	}

	res := ""

	for ind, row := range field {
		if ind > 0 {
			res = res + "\n"
		}
		for _, s := range row {
			res = res + s
		}
	}

	return res

}
