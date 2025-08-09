package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	Row, Col int
}

type PointData struct {
	Cost int
	Next []Next
}

type Next struct {
	Point Point
	Cost  int
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

		field := readField(in, n, m)
		markSea(field)

		from := Point{}
		fmt.Fscanln(in, &from.Row, &from.Col)

		to := Point{}
		fmt.Fscanln(in, &to.Row, &to.Col)

		transitions := countTransitions(field, from, to)
		fmt.Fprintf(out, "%d\n", transitions)
	}
}

func countTransitions(field [][]string, from, to Point) int {
	if from.Row == to.Row && from.Col == to.Col {
		return 0
	}

	points := getPoints(field)

	fromData := points[from]
	fromData.Cost = 0
	deque := []Point{from}

	for len(deque) > 0 {
		current := deque[0]
		deque = deque[1:]

		currentData := points[current]

		if current.Row == to.Row && current.Col == to.Col {
			return currentData.Cost
		}

		for _, next := range currentData.Next {
			nextData := points[next.Point]
			newCost := currentData.Cost + next.Cost

			if newCost < nextData.Cost {
				nextData.Cost = newCost

				if next.Cost == 0 {
					deque = slices.Insert(deque, 0, next.Point)
				} else {
					deque = append(deque, next.Point)
				}
			}
		}
	}

	toData := points[to]
	return toData.Cost
}

func getPoints(field [][]string) map[Point]*PointData {
	res := map[Point]*PointData{}
	for rowID, row := range field {
		for colID := range row {
			from := Point{rowID, colID}
			data := &PointData{
				Cost: 999,
			}
			res[from] = data

			to := Point{from.Row, from.Col - 1}
			if isCellCorrect(field, to) {
				data.Next = append(data.Next, Next{
					Point: to,
					Cost:  calcCost(field, from, to),
				})
			}

			to = Point{from.Row, from.Col + 1}
			if isCellCorrect(field, to) {
				data.Next = append(data.Next, Next{
					Point: to,
					Cost:  calcCost(field, from, to),
				})
			}

			to = Point{from.Row - 1, from.Col}
			if isCellCorrect(field, to) {
				data.Next = append(data.Next, Next{
					Point: to,
					Cost:  calcCost(field, from, to),
				})
			}

			to = Point{from.Row + 1, from.Col}
			if isCellCorrect(field, to) {
				data.Next = append(data.Next, Next{
					Point: to,
					Cost:  calcCost(field, from, to),
				})
			}

			if field[from.Row][from.Col] == "~" {
				to = Point{from.Row - 1, from.Col - 1}
				if isCellCorrect(field, to) && field[to.Row][to.Col] == "~" {
					data.Next = append(data.Next, Next{
						Point: to,
						Cost:  calcCost(field, from, to),
					})
				}

				to = Point{from.Row + 1, from.Col - 1}
				if isCellCorrect(field, to) && field[to.Row][to.Col] == "~" {
					data.Next = append(data.Next, Next{
						Point: to,
						Cost:  calcCost(field, from, to),
					})
				}

				to = Point{from.Row - 1, from.Col + 1}
				if isCellCorrect(field, to) && field[to.Row][to.Col] == "~" {
					data.Next = append(data.Next, Next{
						Point: to,
						Cost:  calcCost(field, from, to),
					})
				}

				to = Point{from.Row + 1, from.Col + 1}
				if isCellCorrect(field, to) && field[to.Row][to.Col] == "~" {
					data.Next = append(data.Next, Next{
						Point: to,
						Cost:  calcCost(field, from, to),
					})
				}
			}
		}
	}

	return res
}

func calcCost(field [][]string, from, to Point) int {
	if field[from.Row][from.Col] != "~" && field[to.Row][to.Col] == "~" ||
		field[from.Row][from.Col] == "~" && field[to.Row][to.Col] != "~" {
		return 1
	}
	return 0
}

func isCellCorrect(field [][]string, p Point) bool {
	if p.Row < 0 || p.Row > len(field)-1 {
		return false
	}
	if p.Col < 0 || p.Col > len(field[0])-1 {
		return false
	}

	return true
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

func readField(in *bufio.Reader, rows, cols int) [][]string {
	field := make([][]string, rows+2)
	for i := 0; i < rows+2; i++ {
		if i == 0 || i == rows+1 {
			field[i] = make([]string, cols+2)
			for j := range field[i] {
				field[i][j] = "~"
			}
			continue
		}

		rowStr, _ := in.ReadString('\n')
		rowStr = strings.Trim(rowStr, "\n")
		rowStr = strings.ReplaceAll(rowStr, " ", "~")
		row := []string{"~"}
		row = append(row, strings.Split(rowStr, "")...)
		row = append(row, "~")
		field[i] = row
	}

	return field
}
