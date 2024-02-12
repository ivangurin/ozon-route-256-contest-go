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

	var el int
	fmt.Fscanln(in, &el)

	existLogins := make([]string, el)

	for i := 0; i < el; i++ {

		var login string
		fmt.Fscanln(in, &login)

		existLogins[i] = login

	}

	var nl int
	fmt.Fscanln(in, &nl)

	newLogins := make([]string, nl)

	for i := 0; i < nl; i++ {

		var login string
		fmt.Fscanln(in, &login)

		newLogins[i] = login

	}

	forbid := make(map[string]struct{}, 1000000)

	sb := strings.Builder{}

	for _, login := range existLogins {

		forbid[login] = struct{}{}

		for i := 0; i < len(login)-1; i++ {
			sb.Reset()
			sb.WriteString(string(login[0:i]))
			sb.WriteString(string(login[i+1]))
			sb.WriteString(string(login[i]))
			sb.WriteString(string(login[i+2:]))
			forbid[sb.String()] = struct{}{}
		}

	}

	for _, login := range newLogins {
		if _, ok := forbid[login]; ok {
			fmt.Fprintln(out, "1")
		} else {
			fmt.Fprintln(out, "0")
		}

	}

}
