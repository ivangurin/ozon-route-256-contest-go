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
	var n, q int
	fmt.Fscanln(in, &n, &q)

	var messageID int
	var lastMassMessageID int
	var users = make(map[int]int)
	for i := 0; i < q; i++ {
		var t, id int
		fmt.Fscanln(in, &t, &id)

		if t == 1 {
			messageID++
			if id == 0 {
				lastMassMessageID = messageID
				users = make(map[int]int)
			} else {
				users[id] = messageID
			}
			continue
		}

		if _, exists := users[id]; exists {
			fmt.Fprintln(out, users[id])
		} else {
			fmt.Fprintln(out, lastMassMessageID)
		}
	}
}
