package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscanln(in, &n, &m)

	c := make([]int, m)
	for i := 0; i < m; i++ {
		c[i] = i + 1
	}

	a := make([]int, n)
	var ai int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ai)
		a[i] = ai
	}
	fmt.Fscanln(in)

	res, err := distribute(a, c)
	if err != nil {
		fmt.Fprintln(out, "-1")
		return
	}

	fmt.Fprintf(out, "%s \n", strings.Trim(fmt.Sprint(res), "[]"))
}

func distribute(friendCards []int, totalCards []int) ([]int, error) {
	res := make([]int, 0, len(friendCards))
	for _, friendCard := range friendCards {

		found := false
		for i := friendCard; i < len(totalCards); i++ {
			if totalCards[i] > friendCard {
				res = append(res, totalCards[i])
				totalCards[i] = 0
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no suitable card")
		}

	}

	return res, nil
}
