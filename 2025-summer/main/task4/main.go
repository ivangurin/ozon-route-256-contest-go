package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)

	for tt := 0; tt < t; tt++ {
		var n, m int
		fmt.Fscanln(in, &n, &m)

		field := ReadField(in, n)
		markSea(field)

		from := Point{}
		fmt.Fscanln(in, &from.row, &from.col)
		from.row--
		from.col--

		to := Point{}
		fmt.Fscanln(in, &to.row, &to.col)
		to.row--
		to.col--

		if IsRouteExists(field, from, to) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func IsRouteExists(field [][]string, from, to Point) bool {
	if from.row == to.row && from.col == to.col {
		return true
	}
	if field[from.row][from.col] == "~" || field[to.row][to.col] == "~" {
		return false
	}

	allRoads := findRoads(field)

	roads := map[Point]struct{}{}
	getRoadsFrom(allRoads, from, &roads)

	_, exist := roads[to]
	if exist {
		return true
	}

	return false
}

func findRoads(field [][]string) map[Point][]Point {
	res := map[Point][]Point{}
	for rowID, row := range field {
		for colID := range row {
			if field[rowID][colID] != "~" {
				from := Point{rowID, colID}

				to := Point{from.row, from.col - 1}
				if isEarth(field, to) {
					res[from] = append(res[from], to)
				}

				to = Point{from.row, from.col + 1}
				if isEarth(field, to) {
					res[from] = append(res[from], to)
				}

				to = Point{from.row - 1, from.col}
				if isEarth(field, to) {
					res[from] = append(res[from], to)
				}

				to = Point{from.row + 1, from.col}
				if isEarth(field, to) {
					res[from] = append(res[from], to)
				}
			}
		}
	}

	return res
}

func getRoadsFrom(roads map[Point][]Point, p Point, res *map[Point]struct{}) {
	(*res)[p] = struct{}{}
	for _, to := range roads[p] {
		_, ok := (*res)[to]
		if ok {
			continue
		}
		getRoadsFrom(roads, to, res)
	}
}

func isEarth(field [][]string, p Point) bool {
	if p.row < 1 || p.row > len(field)-1 {
		return false
	}
	if p.col < 1 || p.col > len(field[0])-1 {
		return false
	}

	if field[p.row][p.col] != "~" {
		return true
	}

	return false
}

func markSea(field [][]string) {
	for rowID, row := range field {
		for colID := range row {
			if !isSea(field, rowID, colID) {
				field[rowID][colID] = " "
			}
		}
	}
}

func isSea(field [][]string, rowID, colID int) bool {
	if rowID == 0 || rowID == len(field)-1 {
		return true
	}
	if colID == 0 || colID == len(field[0])-1 {
		return true
	}
	if field[rowID][colID] != "~" {
		return true
	}
	if field[rowID-1][colID] == "_" &&
		field[rowID][colID-1] == "/" && field[rowID][colID+1] == "\\" &&
		field[rowID+1][colID-1] == "\\" && field[rowID+1][colID+1] == "/" &&
		field[rowID+1][colID] == "_" {
		return false
	}

	if field[rowID][colID-1] == "/" && field[rowID-1][colID+1] == " " {
		return false
	}

	if field[rowID][colID-1] == "\\" && field[rowID-1][colID] == " " {
		return false
	}

	if field[rowID][colID-1] == " " {
		return false
	}

	if field[rowID][colID-1] == "/" && field[rowID-1][colID] == "_" {
		var check1 bool
		counter := 0
		for r := rowID; r < len(field); r++ {
			if field[r][colID-counter] == "\\" {
				counter--
			} else if field[r][colID-counter-1] == "/" {
				counter++
			} else {
				break
			}

			if field[r][colID-1] == "\\" && field[r][colID] != "_" {
				break
			}
			if field[r][colID-1] != "\\" && field[r][colID] == "_" {
				break
			}
			if field[r][colID-1] == "\\" && field[r][colID] == "_" {
				if counter == 0 {
					check1 = true
					break
				}
			}
		}
		if !check1 {
			return true
		}

		var check2 bool
		var rightColID int
		for c := colID; c < len(field[0]); c++ {
			if field[rowID-1][c] == "_" && field[rowID-1][c+1] != "_" && field[rowID][c+1] != "\\" {
				break
			}
			if field[rowID-1][c] == "_" && field[rowID-1][c+1] != "_" && field[rowID][c+1] == "\\" {
				check2 = true
				rightColID = c
				break
			}
		}
		if !check2 {
			return true
		}

		var check3 bool
		for r := rowID + 1; r < len(field); r++ {
			if field[r][rightColID+1] == "/" && field[r][rightColID] != "_" {
				break
			}
			if field[r][rightColID+1] != "/" && field[r][rightColID] == "_" {
				break
			}
			if field[r][rightColID+1] == "/" && field[r][rightColID] == "_" {
				check3 = true
				break
			}
		}
		if !check3 {
			return true
		}

		return false
	}

	return true
}

func ReadField(in *bufio.Reader, rows int) [][]string {
	field := make([][]string, rows)
	for i := range field {
		rowStr, _ := in.ReadString('\n')
		rowStr = strings.Trim(rowStr, "\n")
		rowStr = strings.ReplaceAll(rowStr, " ", "~")
		field[i] = strings.Split(rowStr, "")
	}
	return field
}
