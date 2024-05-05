package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

		var rows, cols int

		fmt.Fscanln(in, &rows, &cols)

		matrix := make([][]int, rows)
		for i := range matrix {
			matrix[i] = make([]int, cols)
		}

		var line string
		for row := 0; row < rows; row++ {
			fmt.Fscanln(in, &line)
			for col := 0; col < cols; col++ {
				matrix[row][col] = str2int(string(line[col]))
			}
		}

		row, col := checkMatrix(matrix)
		fmt.Fprintln(out, row+1, col+1)

	}

}

func str2int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func checkMatrix(m [][]int) (int, int) {

	for s := 5; s > 1; s-- {

		// fmt.Printf("score: %d\n", s)
		ok, r, c := checkMatrixScore(m, s)
		if ok {

			scores := map[int]int{}
			for ri := range m {
				scores[m[ri][0]]++
			}
			if len(scores) == 1 {
				r = 1
			}

			return r, c
		}

	}

	return 0, 0

}

func checkMatrixScore(m [][]int, s int) (bool, int, int) {

	var rr, rc int = -1, -1 // точный отвт
	var ar, ac int = -1, -1 // гипотеза 1 или или
	var br, bc int = -1, -1 // гипотеза 2 или или

	for ri, row := range m {

		count := 0
		c := -1

		for ci, score := range row {
			if ci == rc {
				continue
			}
			if score < s {
				count++
				c = ci
			}
		}

		if count > 1 {
			if rr == -1 {
				if bc >= 0 {
					// fmt.Printf("FAIL0: %d %d\n", ri, c)
					return false, -1, -1
				}
				if ac >= 0 {
					rc = ac
					ac = -1
					ar = -1
					// fmt.Printf("CC0: %d\n", rc)
				}
				rr = ri
				// fmt.Printf("RR1: %d\n", rr)
			} else {
				return false, -1, -1
			}
		} else if count == 1 {
			if ri == rr || c == rc {

			} else if rr >= 0 && rc >= 0 {
				// fmt.Printf("FAIL1: %d %d\n", ri, c)
				return false, -1, -1
			} else if rr >= 0 {
				rc = c
				// fmt.Printf("CC1: %d\n", rc)
			} else if rc >= 0 {
				rr = ri
				// fmt.Printf("RR2: %d\n", rr)
			} else if ar == -1 {
				ar = ri
				ac = c
				// fmt.Printf("A: %d %d\n", ar, ac)
			} else if ac == c {
				rc = c
				ar = br
				ac = bc
				br = -1
				bc = -1
				// fmt.Printf("CC3: %d\n", rc)
			} else if bc == c {
				rc = c
				br = -1
				bc = -1
				// fmt.Printf("CC4: %d\n", rc)
			} else if br == -1 {
				br = ri
				bc = c
				// fmt.Printf("B: %d %d\n", br, bc)
			} else {
				// fmt.Printf("FAIL: %d %d\n", ri, c)
				return false, -1, -1
			}

		}

	}

	if rr == -1 && rc == -1 && bc == -1 {
		rc = ac
		rr = 0
	}

	if rr == -1 {
		rr = ar
	}
	if rr == -1 {
		rr = br
	}
	if rr == -1 {
		rr = 0
	}
	if rc == -1 {
		rc = bc
	}
	if rc == -1 {
		rc = ac
	}
	if rc == -1 {
		rc = 0
	}

	return true, rr, rc

}
