package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var counter int
	fmt.Fscanln(in, &counter)

	type date struct {
		day   string
		month string
		year  string
	}

	type datesT []date

	dates := make(datesT, counter)
	for i := 0; i < counter; i++ {
		var d, m, y string
		fmt.Fscan(in, &d, &m, &y)
		if len(d) == 1 {
			d = "0" + d
		}
		if len(m) == 1 {
			m = "0" + m
		}
		dates[i] = date{d, m, y}
	}

	for i, date := range dates {
		if i > 0 && i < len(dates) {
			fmt.Fprint(out, "\n")
		}
		_, err := time.Parse("02.01.2006", fmt.Sprintf("%s.%s.%s", date.day, date.month, date.year))
		if err != nil {
			fmt.Fprint(out, "NO")
			continue
		}
		fmt.Fprint(out, "YES")
	}
}
