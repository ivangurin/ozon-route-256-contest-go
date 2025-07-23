package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type action struct {
	p1 int
	p2 int
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

		field := make([][]string, 0, n)
		for j := 0; j < n; j++ {
			rowStr, _ := in.ReadString('\n')
			rowStr = strings.Trim(rowStr, "\n")
			row := strings.Split(rowStr, "")
			field = append(field, row)
		}

		actions := make([]action, 0, k)
		for j := 0; j < k; j++ {
			action := action{}
			fmt.Fscanln(in, &action.p1, &action.p2)
			actions = append(actions, action)
		}

		res := MakeIt(field, actions)

		_ = res
	}
}

func MakeIt(field [][]string, actions []action) [][]string {

	return nil
}
