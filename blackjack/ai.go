package blackjack

import (
	"BlackJackAI/deck"
	"fmt"
)

type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hands [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	// noop
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	// noop
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	fmt.Println("========NEW HAND=================")
	if shuffled {
		fmt.Println("The deck was just shuffled.")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand, "Score:", Score(hand...))
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand, (d)ouble, s(p)lit")
		fmt.Println("Jack White:", ai.Recomendaciones(hand, dealer))
		var input string
		var flag bool
		flag = true
		for flag {
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				return MoveHit
			case "s":
				return MoveStand
			case "d":
				if len(hand) == 2 {
					return MoveDouble
				} else {
					fmt.Println("You can't double this hand")
				}

			case "p":
				if len(hand) == 2 && hand[0] == hand[1] {
					return MoveSplit
				} else {
					fmt.Println("You can't split this hand")
				}
			default:
				fmt.Println("Invalid option:", input)
			}
		}
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("=========================FINAL HANDS=====================")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer, "\tScore:", Score(dealer...))
}

func (ai humanAI) Recomendaciones(hand []deck.Card, dealer deck.Card) string {
	score := Score(hand...)
	dScore := Score(dealer)
	var text string
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				if cardScore == 9 && (dScore == 7 || (dScore > 9)) {
					text = ("I  recommend stand")
					return text

				}
				text = ("I recommend split")
				return text
			}
		}
		if (score == 10 || score == 11) && !Soft(hand...) {
			text = ("I recommend double")
			return text
		}
	}
	if dScore >= 5 && dScore <= 6 {
		text = ("I recommend stand")
		return text
	}
	if score < 18 {
		if Soft(hand...) { //TENEMOS UN AS
			if score <= 17 {
				if (dScore == 5 || dScore == 6) && len(hand) == 2 {
					text = ("I recommend double")
					return text
				} else {
					text = ("I recommend hit")
					return text
				}
			}

		} else { //NO HAY UN AS
			if score <= 9 {

			} else {
				if (score == 10 || score == 11) && len(hand) == 2 {
					text = ("I recommend double")
					return text
				}
				if score >= 12 && score <= 16 && dScore <= 6 {
					text = ("I recommend stand")
					return text
				}
				if score >= 12 && score <= 16 && dScore >= 7 {
					text = ("I recommend hit")
					return text
				}
				if score >= 17 {
					text = ("I recommend stand")
					return text

				}
			}
		}
		// return blackjack.MoveHit
	}
	text = ("I recommend stand")
	return text
}
