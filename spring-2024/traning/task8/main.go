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

	for ti := 0; ti < t; ti++ {
		var n, b, r, f int
		fmt.Fscanln(in, &n, &b, &r, &f)

		words := make([]string, n)
		for ni := 0; ni < n; ni++ {
			var word string
			fmt.Fscanln(in, &word)
			words[ni] = word
		}

		fmt.Fprintln(out, CodeNames(words, b, r, f))
	}
}

func CodeNames(words []string, b, r, f int) string {
	blueSubStrs := GetWordsSubStrings(words[:b])
	redSubStrs := GetWordsSubStrings(words[b : b+r])
	blackSubStrs := GetWordSubStrings(words[f-1])

	for _, word := range words {
		delete(blueSubStrs, word)
	}

	var maxPoints int
	var maxSubStr string
	for blueSubStr, blueSubStrPoints := range blueSubStrs {
		if _, exists := blackSubStrs[blueSubStr]; exists {
			continue
		}

		points := blueSubStrPoints

		redSubStrPoints, exists := redSubStrs[blueSubStr]
		if exists {
			points -= redSubStrPoints
		}

		if points > maxPoints {
			maxPoints = points
			maxSubStr = blueSubStr
		}

	}

	if maxSubStr == "" {
		return "tkhapjiabb 0"
	}

	return fmt.Sprintf("%s %d", maxSubStr, maxPoints)
}

func GetWordsSubStrings(words []string) map[string]int {
	res := make(map[string]int, 1000)
	for _, word := range words {
		substrs := GetWordSubStrings(word)
		for substr := range substrs {
			res[substr]++
		}
	}
	return res
}

func GetWordSubStrings(word string) map[string]int {
	res := make(map[string]int, len(word)*len(word))
	for i := 0; i < len(word); i++ {
		for j := i; j < len(word); j++ {
			res[word[i:j+1]] = 1
		}
	}
	return res
}
