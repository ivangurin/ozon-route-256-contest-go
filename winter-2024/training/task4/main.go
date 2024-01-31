package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var template *regexp.Regexp

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	template, _ = regexp.Compile(`([A-Z]{1}[0-9]{1,2}[A-Z]{2})`)

	var counter int
	fmt.Fscanln(in, &counter)

	lines := make([]string, counter)
	for i := 0; i < counter; i++ {
		var s string
		fmt.Fscan(in, &s)
		lines[i] = s
	}

	for i, line := range lines {
		fmt.Fprint(out, convert(line))
		if i < len(lines)-1 {
			fmt.Fprint(out, "\n")
		}
	}

}

func convert(line string) string {

	numbers := template.FindAllStringSubmatch(line, -1)

	lenNums := 0
	res := ""

	for _, number := range numbers {
		lenNums += len(number[0])
		if res != "" {
			res = res + " "
		}
		res = res + number[0]
	}

	if len(line) != lenNums {
		return "-"
	} else {
		return res
	}

}
