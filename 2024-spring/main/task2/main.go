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
	var n, t int
	fmt.Fscanln(in, &n, &t)
	var letters = map[string]int{}
	for i := 0; i < n; i++ {
		var l string
		fmt.Fscan(in, &l)
		letters[l]++
	}
	fmt.Fscanln(in)
	for i := 0; i < t; i++ {
		var pin string
		fmt.Fscan(in, &pin)
		if isPinCorrect(letters, pin) {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

func isPinCorrect(letters map[string]int, pin string) bool {
	pinLetters := map[string]int{}
	for _, l := range pin {
		pinLetters[string(l)]++
	}

	if len(letters) != len(pinLetters) {
		return false
	}

	for letter, letterCounter := range letters {
		pinLetterCounter, exists := pinLetters[letter]
		if !exists {
			return false
		}
		if letterCounter != pinLetterCounter {
			return false
		}
	}

	for pinLetter := range pinLetters {
		_, exists := letters[pinLetter]
		if !exists {
			return false
		}
	}

	return true
}
