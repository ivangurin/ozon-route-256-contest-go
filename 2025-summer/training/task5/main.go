package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Operation struct {
	p1 int
	p2 int
}

type Point struct {
	x int
	y int
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

	for i := 0; i < t; i++ {
		var n, m, k int
		fmt.Fscanln(in, &n, &m, &k)

		paper := make([][]string, 0, n)
		for j := 0; j < n; j++ {
			rowStr, _ := in.ReadString('\n')
			rowStr = strings.Trim(rowStr, "\n")
			row := strings.Split(rowStr, "")
			paper = append(paper, row)
		}

		operations := make([]Operation, 0, k)
		for j := 0; j < k; j++ {
			operation := Operation{}
			fmt.Fscanln(in, &operation.p1, &operation.p2)
			operations = append(operations, operation)
		}

		for _, operation := range operations {
			paper = FoldIt(paper, operation)

			for _, row := range paper {
				fmt.Fprintln(out, strings.Join(row, ""))
				// fmt.Println(strings.Join(row, ""))
			}
			fmt.Fprintln(out)
			// fmt.Println()
		}
	}
}

// Решение на 11 баллов без сгибания по диагонали
func FoldIt(paper [][]string, operation Operation) [][]string {
	p1 := GetPoint(paper, operation.p1)
	p2 := GetPoint(paper, operation.p2)

	if p1.x == p2.x {
		if p1.y > p2.y {
			return FoldLeft(paper, p1.x)
		} else {
			return FoldRight(paper, p1.x)
		}
	} else if p1.y == p2.y {
		if p1.x > p2.x {
			return FoldDown(paper, p1.y)
		} else {
			return FoldUp(paper, p1.y)
		}
	}

	return paper
}

func GetPoint(paper [][]string, point int) Point {
	p := Point{}
	if point <= len(paper[0])+1 {
		p.x = point - 1
		p.y = 0
	} else if point <= len(paper[0])+len(paper)+1 {
		p.x = len(paper[0])
		p.y = point - len(paper[0]) - 1
	} else if point <= len(paper[0])*2+len(paper)+1 {
		p.x = len(paper[0])*2 + len(paper) + 1 - point
		p.y = len(paper)
	} else {
		p.x = 0
		p.y = len(paper[0])*2 + len(paper)*2 + 1 - point
	}

	return p
}

func FoldLeft(paper [][]string, x int) [][]string {
	field := CreateField(paper)

	line := 10 + x
	for row := 0; row < 30; row++ {
		for col := 0; col < 10; col++ {
			if field[row][line+col] == "-" {
				continue
			}
			if field[row][line-col-1] == "#" {
				field[row][line+col] = "-"
				continue
			}

			field[row][line-col-1] = field[row][line+col]
			field[row][line+col] = "-"

		}
	}

	field = CutEmptyBorders(field)
	field = ExtractObject(field)

	return field
}

func FoldRight(paper [][]string, x int) [][]string {
	field := CreateField(paper)

	line := 10 + x
	for row := 0; row < 30; row++ {
		for col := 0; col < 10; col++ {
			if field[row][line-col-1] == "-" {
				continue
			}
			if field[row][line+col] == "#" {
				field[row][line-col-1] = "-"
				continue
			}

			field[row][line+col] = field[row][line-col-1]
			field[row][line-col-1] = "-"

		}
	}

	field = CutEmptyBorders(field)
	field = ExtractObject(field)

	return field
}

func FoldUp(paper [][]string, y int) [][]string {
	field := CreateField(paper)

	line := 10 + y
	for col := 0; col < 30; col++ {
		for row := 0; row < 10; row++ {
			if field[line+row][col] == "-" {
				continue
			}
			if field[line-row-1][col] == "#" {
				field[line+row][col] = "-"
				continue
			}

			field[line-row-1][col] = field[line+row][col]
			field[line+row][col] = "-"

		}
	}

	field = CutEmptyBorders(field)
	field = ExtractObject(field)

	return field
}

func FoldDown(paper [][]string, y int) [][]string {
	field := CreateField(paper)

	line := 10 + y
	for col := 0; col < 30; col++ {
		for row := 0; row < 10; row++ {
			if field[line-row-1][col] == "-" {
				continue
			}
			if field[line+row][col] == "#" {
				field[line-row-1][col] = "-"
				continue
			}

			field[line+row][col] = field[line-row-1][col]
			field[line-row-1][col] = "-"

		}
	}

	field = CutEmptyBorders(field)
	field = ExtractObject(field)

	return field
}

func CreateField(paper [][]string) [][]string {
	field := make([][]string, 30)
	for i := range field {
		field[i] = make([]string, 30)
	}
	for row := 0; row < 30; row++ {
		for col := 0; col < 30; col++ {
			field[row][col] = "-"
		}
	}
	for row := 10; row < 10+len(paper); row++ {
		for col := 10; col < 10+len(paper[0]); col++ {
			field[row][col] = paper[row-10][col-10]
		}
	}
	return field
}

func CutEmptyBorders(field [][]string) [][]string {
	for rowID := 0; rowID < len(field); rowID++ {
		dotsCounter := 0
		reshCounter := 0
		for colID := 0; colID < len(field[rowID]); colID++ {
			if field[rowID][colID] == "." {
				dotsCounter++
			}
			if field[rowID][colID] == "#" {
				reshCounter++
			}
		}
		if dotsCounter > 0 && reshCounter == 0 {
			for colID := 0; colID < len(field[rowID]); colID++ {
				field[rowID][colID] = "-"
			}
		}
		if reshCounter > 0 {
			break
		}
	}

	for rowID := len(field) - 1; rowID >= 0; rowID-- {
		dotsCounter := 0
		reshCounter := 0
		for colID := 0; colID < len(field[rowID]); colID++ {
			if field[rowID][colID] == "." {
				dotsCounter++
			}
			if field[rowID][colID] == "#" {
				reshCounter++
			}
		}
		if dotsCounter > 0 && reshCounter == 0 {
			for colID := 0; colID < len(field[rowID]); colID++ {
				field[rowID][colID] = "-"
			}
		}
		if reshCounter > 0 {
			break
		}
	}

	for colID := 0; colID < len(field[0]); colID++ {
		dotsCounter := 0
		reshCounter := 0
		for rowID := 0; rowID < len(field); rowID++ {
			if field[rowID][colID] == "." {
				dotsCounter++
			}
			if field[rowID][colID] == "#" {
				reshCounter++
			}
		}
		if dotsCounter > 0 && reshCounter == 0 {
			for rowID := 0; rowID < len(field); rowID++ {
				field[rowID][colID] = "-"
			}
		}
		if reshCounter > 0 {
			break
		}
	}

	for colID := len(field[0]) - 1; colID >= 0; colID-- {
		dotsCounter := 0
		reshCounter := 0
		for rowID := 0; rowID < len(field); rowID++ {
			if field[rowID][colID] == "." {
				dotsCounter++
			}
			if field[rowID][colID] == "#" {
				reshCounter++
			}
		}
		if dotsCounter > 0 && reshCounter == 0 {
			for rowID := 0; rowID < len(field); rowID++ {
				field[rowID][colID] = "-"
			}
		}
		if reshCounter > 0 {
			break
		}
	}

	return field
}

func ExtractObject(field [][]string) [][]string {
	var minX, maxX, minY, maxY *int
	for rowID, row := range field {
		for colID, val := range row {
			if val != "-" {
				if minX == nil || colID < *minX {
					minX = &colID
				}
				if maxX == nil || colID > *maxX {
					maxX = &colID
				}
				if minY == nil || rowID < *minY {
					minY = &rowID
				}
				if maxY == nil || rowID > *maxY {
					maxY = &rowID
				}
			}
		}
	}

	res := make([][]string, *maxY-*minY+1)
	for rowID := range res {
		res[rowID] = make([]string, *maxX-*minX+1)
		for colID := range res[rowID] {
			res[rowID][colID] = " "
		}
	}

	for rowID := *minY; rowID <= *maxY; rowID++ {
		for colID := *minX; colID <= *maxX; colID++ {
			if field[rowID][colID] != "-" {
				res[rowID-*minY][colID-*minX] = field[rowID][colID]
			}
		}
	}

	return res
}
