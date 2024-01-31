package main

import (
	"bufio"
	"fmt"
	"os"
)

type value struct {
	value  string
	points string
}

type Card struct {
	Value string
	Suit  string
}

type Cards []Card

type Deck struct {
	cards   Cards
	removed map[Card]bool
}

type Players []Cards

var values []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

var points map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var suits []string = []string{"S", "C", "D", "H"}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {

	var setCounter int
	fmt.Fscanln(in, &setCounter)

	var card1 string
	var card2 string

	for i := 0; i < setCounter; i++ {

		var playersCounter int
		fmt.Fscanln(in, &playersCounter)

		players := make(Players, playersCounter)

		for j := 0; j < playersCounter; j++ {
			fmt.Fscanln(in, &card1, &card2)
			cards := Cards{{card1[:1], card1[1:]}, {card2[:1], card2[1:]}}
			players[j] = cards
		}

		winCards := GetWinCards(players)

		fmt.Fprintln(out, len(winCards))
		for _, c := range winCards {
			fmt.Fprintf(out, "%s%s\n", c.Value, c.Suit)
		}

	}

}

func NewDeck() *Deck {

	cards := make(Cards, 0, len(values)*len(suits))

	for _, v := range values {
		for _, s := range suits {
			cards = append(cards, Card{v, s})
		}
	}

	return &Deck{
		cards:   cards,
		removed: map[Card]bool{},
	}

}

func (d *Deck) RemoveCard(card Card) {
	d.removed[card] = true
}

func (d *Deck) GetCards() Cards {
	res := make(Cards, 0, len(d.cards)-len(d.removed))
	for _, c := range d.cards {
		if _, removed := d.removed[c]; !removed {
			res = append(res, c)
		}
	}
	return res
}

func GetWinCards(players Players) Cards {

	res := Cards{}

	deck := NewDeck()

	for _, playerCards := range players {
		for _, card := range playerCards {
			deck.RemoveCard(card)
		}
	}

	deckCards := deck.GetCards()

	for _, deckCard := range deckCards {

		player1points := 0
		playerNpoints := 0

		for i, playerCards := range players {

			points := EvalCombination(&playerCards[0], &playerCards[1], &deckCard)

			if i == 0 {
				player1points = points
			} else {
				playerNpoints = points
			}

			if player1points < playerNpoints {
				break
			}

		}

		if player1points >= playerNpoints {
			res = append(res, deckCard)
		}

	}

	return res

}

func EvalCombination(card1 *Card, card2 *Card, card3 *Card) int {

	res := 0

	if card1.Value == card2.Value && card1.Value == card3.Value {
		res = 3000 * points[card1.Value]
	} else if card1.Value == card2.Value {
		res = 200 * points[card1.Value]
	} else if card1.Value == card3.Value {
		res = 200 * points[card1.Value]
	} else if card2.Value == card3.Value {
		res = 200 * points[card2.Value]
	} else {
		res = Max(points[card1.Value], points[card2.Value], points[card3.Value])
	}

	return res

}

func Max(values ...int) int {
	res := 0
	for i, v := range values {
		if i == 0 {
			res = v
		} else {
			if res < v {
				res = v
			}
		}
	}
	return res
}
