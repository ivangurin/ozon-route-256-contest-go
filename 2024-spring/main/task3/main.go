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
	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Fscanln(in, &s)
		if isCorrect(s) {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

func isCorrect(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	if s[0] != s[len(s)-1] {
		return false
	}

	for i := 1; i < len(s)-1; i++ {
		if s[i] == s[0] {
			continue
		}
		if s[i-1] != s[i+1] {
			return false
		}
	}

	return true
}
