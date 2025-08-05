package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	Total   int
	Printed string
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var counter int
	fmt.Fscanln(in, &counter)

	sets := make([]*Set, 0, counter)

	for i := 0; i < counter; i++ {
		var total int
		fmt.Fscanln(in, &total)
		var printed string
		fmt.Fscanln(in, &printed)
		sets = append(sets, &Set{total, printed})
	}

	for i, set := range sets {
		fmt.Fprint(out, getPagesToPrint(set.Total, getPages(set.Printed)))
		if i < len(sets)-1 {
			fmt.Fprint(out, "\n")
		}
	}

}

func getPages(s string) map[int]struct{} {

	res := map[int]struct{}{}

	list := strings.Split(s, ",")

	for _, row := range list {

		if strings.Contains(row, "-") {
			diap := convDiap(row)
			for k := range diap {
				res[k] = struct{}{}
			}
		} else {
			res[str2int(row)] = struct{}{}
		}

	}

	return res

}

func getPagesToPrint(total int, printed map[int]struct{}) string {

	res := ""

	pages := []int{}
	for i := 1; i <= total; i++ {
		if _, exists := printed[i]; exists {
			continue
		}
		pages = append(pages, i)
	}

	if len(pages) == 0 {
		return res
	}

	from := 0
	to := 0

	for i := 0; i < len(pages); i++ {

		if i == 0 {
			from = pages[i]
			to = pages[i]
			continue
		}

		if pages[i]-pages[i-1] == 1 {
			to = pages[i]
		} else {

			if res != "" {
				res = res + ","
			}

			page := ""
			if from == to {
				page = fmt.Sprintf("%d", from)
			} else {
				page = fmt.Sprintf("%d-%d", from, to)
			}

			res = res + page

			from = pages[i]
			to = pages[i]

		}

	}

	if res != "" {
		res = res + ","
	}

	page := ""
	if from == to {
		page = fmt.Sprintf("%d", from)
	} else {
		page = fmt.Sprintf("%d-%d", from, to)
	}

	res = res + page

	return res

}

func convDiap(s string) map[int]struct{} {

	fromTo := strings.Split(s, "-")

	res := map[int]struct{}{}

	for i := str2int(fromTo[0]); i <= str2int(fromTo[1]); i++ {
		res[i] = struct{}{}
	}

	return res

}

func str2int(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}
