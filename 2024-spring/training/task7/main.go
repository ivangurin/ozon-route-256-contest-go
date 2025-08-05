package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type person struct {
	window int
	index  int
	action string
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

	for ti := 0; ti < t; ti++ {
		var n, m int
		fmt.Fscanln(in, &n, &m)

		persons := make([]person, m)
		for mi := 0; mi < m; mi++ {
			var w int
			fmt.Fscan(in, &w)
			persons[mi] = person{
				window: w,
				index:  mi,
			}
		}
		fmt.Fscanln(in)

		res, err := reorderPersons(persons, n)
		if err != nil {
			fmt.Fprintln(out, "x")
			continue
		}

		fmt.Fprintln(out, res)
	}

}

func reorderPersons(persons []person, n int) (string, error) {

	if len(persons) > n {
		return "", fmt.Errorf("too many persons")
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].window < persons[j].window
	})

	if persons[0].window > 1 {
		persons[0].window--
		persons[0].action = "-"
	} else {
		persons[0].action = "0"
	}

	for i := 1; i < len(persons); i++ {

		if (persons[i].window - persons[i-1].window) == 1 {
			persons[i].action = "0"
		}

		if (persons[i].window - persons[i-1].window) > 1 {
			persons[i].window--
			persons[i].action = "-"
		}

		if (persons[i].window - persons[i-1].window) == 0 {
			persons[i].window++
			persons[i].action = "+"
		}

		if (persons[i].window - persons[i-1].window) < 0 {
			return "", fmt.Errorf("no solution")
		}

		if persons[i].window > n {
			return "", fmt.Errorf("no solution")
		}

	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].index < persons[j].index
	})

	sb := strings.Builder{}
	for _, p := range persons {
		sb.WriteString(p.action)
	}

	return sb.String(), nil
}
