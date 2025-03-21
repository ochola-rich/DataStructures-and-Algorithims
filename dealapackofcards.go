package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for player := 1; player <= 4; player++ {
		fmt.Printf("Player %v:", player)
		switch player {
		case 1:
			fmt.Printf(" %v, %v, %v\n", deck[0], deck[1], deck[2])
		case 2:
			fmt.Printf(" %v, %v, %v\n", deck[3], deck[4], deck[5])
		case 3:
			fmt.Printf(" %v, %v, %v\n", deck[6], deck[7], deck[8])
		default:
			fmt.Printf(" %v, %v, %v\n", deck[9], deck[10], deck[11])
		}
	}
}
