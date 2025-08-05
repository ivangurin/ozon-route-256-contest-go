package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	sign  string
	tempr int
}

type Persons []Person

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var counter int
	fmt.Fscanln(in, &counter)

	sets := make([]Persons, counter)

	for i := 0; i < counter; i++ {

		var personCounter int
		fmt.Fscanln(in, &personCounter)

		sets[i] = make(Persons, personCounter)

		for j := 0; j < personCounter; j++ {

			var s string
			var t int
			fmt.Fscanln(in, &s, &t)

			sets[i][j] = Person{s, t}

		}

	}

	for _, persons := range sets {

		from := 15
		to := 30

		for _, person := range persons {

			if person.sign == ">=" {
				if from < person.tempr {
					from = person.tempr
				}
			} else {
				if to > person.tempr {
					to = person.tempr
				}
			}

			if from <= to {
				fmt.Fprint(out, from, "\r\n")
			} else {
				fmt.Fprint(out, -1, "\r\n")
			}

		}

		fmt.Fprint(out, "\r\n")

	}

}
