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
	var s string
	fmt.Fscanln(in, &s)

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		var f, t int
		var ss string
		fmt.Fscanln(in, &f, &t, &ss)
		s = s[:f-1] + ss + s[t:]

	}

	fmt.Fprintln(out, s)
}
