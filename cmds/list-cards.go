package main

import (
	"fmt"
	"github.com/go-piv/piv-go/piv"
	"strings"
)

func main() {
	cards, err := piv.Cards()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	for i, card := range cards {
		fmt.Printf("Card %d: %s\n", i, card)

		if !strings.Contains(strings.ToLower(card), "yubikey") {
			continue
		}
		yk, err := piv.Open(card)
		if err != nil {
			fmt.Printf("Error opening card: %s\n", err.Error())
			continue
		}
		fmt.Printf("  Version: %d.%d.%d\n", yk.Version().Major, yk.Version().Minor, yk.Version().Patch)
		serial, err := yk.Serial()
		if err == nil {
			fmt.Printf("  Serial: %d\n", serial)
		}
	}
}
