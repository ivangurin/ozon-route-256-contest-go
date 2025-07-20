package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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
		var a int
		fmt.Fscanln(in, &a)

		answers := make([]string, 0, a)
		for j := 0; j < a; j++ {
			answer, _ := in.ReadString('\n')
			answer = strings.TrimRight(answer, "\n")
			answers = append(answers, answer)
		}

		res := WhoIs(answers)

		for _, resRow := range res {
			fmt.Fprintln(out, resRow)
		}
	}
}

func WhoIs(answers []string) []string {
	persons := map[string]int{}
	action := ""
	for _, answer := range answers {
		answerParts := strings.Split(answer, ":")
		name := answerParts[0]
		replay := strings.Trim(answerParts[1], " ")
		replayParts := strings.Split(replay, " ")

		_, exists := persons[name]
		if !exists {
			persons[name] = 0
		}

		if replayParts[1] == "am" {
			if replayParts[2] == "not" {
				persons[name]--
			} else {
				persons[name] += 2
			}
		}

		if replayParts[1] == "is" {
			if replayParts[2] == "not" {
				persons[replayParts[0]]--
			} else {
				persons[replayParts[0]]++
			}
		}

		action = replayParts[len(replayParts)-1]
	}

	var hasMaxScore bool
	var maxScore int
	for _, score := range persons {
		if !hasMaxScore || score > maxScore {
			hasMaxScore = true
			maxScore = score
		}
	}

	res := []string{}
	for name, score := range persons {
		if score == maxScore {
			res = append(res, name+" is "+string(action[:len(action)-1])+".")
		}
	}

	slices.Sort(res)

	return res
}
