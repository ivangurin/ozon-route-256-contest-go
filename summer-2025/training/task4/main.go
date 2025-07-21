package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
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
	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var r, c int
		fmt.Fscanln(in, &r, &c)

		field := make([][]string, 0, r)
		for j := 0; j < r; j++ {
			rowStr, _ := in.ReadString('\n')
			rowStr = strings.Trim(rowStr, "\n")
			row := strings.Split(rowStr, "")
			field = append(field, row)
		}

		isValid := checkField(field)
		if isValid {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}

}

func checkField(field [][]string) bool {
	hexes := map[string]map[Pos]bool{}
	for r, row := range field {
		for c, v := range row {
			if v != "." {
				p := Pos{r, c}
				_, exists := hexes[v]
				if !exists {
					hexes[v] = map[Pos]bool{}
				}
				hexes[v][p] = false
			}
		}
	}

	for _, poses := range hexes {
		for p := range poses {
			visit(poses, p)
			break
		}
	}

	for _, poses := range hexes {
		for _, visited := range poses {
			if !visited {
				return false
			}
		}
	}

	return true
}

func getNextMove(p Pos) []Pos {
	return []Pos{
		{p.row, p.col - 2},     // Left
		{p.row, p.col + 2},     // Right
		{p.row - 1, p.col - 1}, // Up-Left
		{p.row - 1, p.col + 1}, // Up-Right
		{p.row + 1, p.col - 1}, // Down-Left
		{p.row + 1, p.col + 1}, // Down-Right
	}
}

func visit(poses map[Pos]bool, p Pos) {
	visited, exists := poses[p]
	if !exists {
		return
	}
	if visited {
		return
	}

	poses[p] = true

	nextPoses := getNextMove(p)
	for _, nextPos := range nextPoses {
		visit(poses, nextPos)
	}
}
