package main

import (
	"bufio"
	"fmt"
	"os"
)

var possibleCommands map[string]int8

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	possibleCommands = map[string]int8{
		"MR": 0,
		"MC": 0,
		"MD": 0,
		"RC": 0,
		"RD": 0,
		"CM": 0,
		"DM": 0,
	}

	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {

		var commands string
		fmt.Fscanln(in, &commands)

		fmt.Fprintln(out, checkCommands(commands))

	}

}

func checkCommands(commands string) string {

	if string(commands[0]) != "M" {
		return "NO"
	}

	if string(commands[len(commands)-1]) != "D" {
		return "NO"
	}

	if len(commands) == 1 {
		return "NO"
	}

	for j := 0; j < len(commands)-1; j++ {

		command := string(commands[j]) + string(commands[j+1])

		if _, ok := possibleCommands[command]; !ok {
			return "NO"
		}

	}

	return "YES"

}
