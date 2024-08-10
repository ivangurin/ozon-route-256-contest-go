package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscanln(in, &n, &m)

	a := make([]int, n)
	var ai int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ai)
		a[i] = ai
	}
	fmt.Fscanln(in)

	res, err := distribute(a, m)
	if err != nil {
		fmt.Fprintln(out, "-1")
		return
	}

	fmt.Fprintf(out, "%s \n", strings.Trim(fmt.Sprint(res), "[]"))
}

func distribute(friendCards []int, totalCards int) ([]int, error) {
	if len(friendCards) > totalCards {
		return nil, fmt.Errorf("no suitable card")
	}

	type fiend struct {
		id   int
		card int
	}

	friends := make([]fiend, len(friendCards))
	for id, card := range friendCards {
		friends[id] = fiend{id, card}
	}

	sort.Slice(friends, func(i, j int) bool {
		if friends[i].card < friends[j].card {
			return true
		}
		if friends[i].card == friends[j].card {
			if friends[i].id < friends[j].id {
				return true
			}
		}
		return false
	})

	lastCard := 0
	res := make([]int, len(friends))
	for _, friend := range friends {

		if friend.card > lastCard {
			lastCard = friend.card
		}

		found := false
		for nextCard := lastCard + 1; nextCard <= totalCards; nextCard++ {
			res[friend.id] = nextCard
			lastCard = nextCard
			found = true
			break
		}
		if !found {
			return nil, fmt.Errorf("no suitable card")
		}

	}

	return res, nil
}
