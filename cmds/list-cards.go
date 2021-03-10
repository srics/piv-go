package main

import (
	"fmt"
	"github.com/go-piv/piv-go/piv"
)

func main() {
	cards, err := piv.Cards()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	for i, card := range cards {
		fmt.Printf("Card %d: %s\n", i, card)
	}
}