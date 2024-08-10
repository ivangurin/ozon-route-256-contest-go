package main

import (
	"bufio"
	"fmt"
	"os"
)

type rate struct {
	from        string
	to          string
	numerator   int
	denominator int
}

type bank struct {
	rates []rate
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	banks := make([]bank, 0, 3)
	for i := 0; i < 3; i++ {

		bank := bank{
			rates: []rate{
				{from: "RUB", to: "USD"},
				{from: "RUB", to: "EUR"},
				{from: "USD", to: "RUB"},
				{from: "USD", to: "EUR"},
				{from: "EUR", to: "RUB"},
				{from: "EUR", to: "USD"},
			},
		}

		for i, rate := range bank.rates {
			fmt.Fscanln(in, &rate.numerator, &rate.denominator)
			bank.rates[i] = rate
		}

		banks = append(banks, bank)

	}

	fmt.Println(banks)

}

// func exchange(banks []bank) float64 {

// }
