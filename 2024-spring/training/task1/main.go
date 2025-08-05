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
	var n int
	fmt.Fscanln(in, &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
