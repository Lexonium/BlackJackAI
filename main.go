package main

import (
	"fmt"

	"BlackJackAI/blackjack"
	"BlackJackAI/deck"
)

type basicAI struct {
	score int
	seen  int
	decks int
}

func (ai *basicAI) Bet(shuffled bool) int {
	if shuffled {
		ai.score = 0
		ai.seen = 0
	}
	trueScore := ai.score / ((ai.decks*52 - ai.seen) / 52)
	switch {
	case trueScore >= 14:
		fmt.Println("I bet 100")
		return 100
	case trueScore >= 8:
		fmt.Println("I bet 50")
		return 50
	default:
		fmt.Println("I bet 10")
		return 10
	}
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	score := blackjack.Score(hand...)
	dScore := blackjack.Score(dealer)
	fmt.Println("Player:", hand, "Score:", score)
	fmt.Println("Dealer:", dealer, "Score:", dScore)
	fmt.Println("What will you do? (h)it, (s)tand, (d)ouble, s(p)lit")
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := blackjack.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				if cardScore == 9 && (dScore == 7 || (dScore > 9)) {
					fmt.Println("I stand")
					return blackjack.MoveStand
				}
				fmt.Println("I split")
				return blackjack.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !blackjack.Soft(hand...) {
			fmt.Println("I double")
			return blackjack.MoveDouble
		}
	}
	if dScore >= 5 && dScore <= 6 {
		fmt.Println("I stand")
		return blackjack.MoveStand
	}
	if score < 18 {
		if blackjack.Soft(hand...) { //TENEMOS UN AS
			if score <= 17 {
				if (dScore == 5 || dScore == 6) && len(hand) == 2 {
					fmt.Println("I double")
					return blackjack.MoveDouble
				} else {
					fmt.Println("I hit")
					return blackjack.MoveHit
				}
			}

		} else { //NO HAY UN AS
			if score <= 9 {
				return blackjack.MoveHit
			} else {
				if (score == 10 || score == 11) && len(hand) == 2 {
					fmt.Println("I double")
					return blackjack.MoveDouble
				}
				if score >= 12 && score <= 16 && dScore <= 6 {
					fmt.Println("I stand")
					return blackjack.MoveStand
				}
				if score >= 12 && score <= 16 && dScore >= 7 {
					fmt.Println("I hit")
					return blackjack.MoveHit
				}
				if score >= 17 {
					fmt.Println("I stand")
					return blackjack.MoveStand
				}
			}
		}
		// return blackjack.MoveHit
	}
	fmt.Println("I stand")
	return blackjack.MoveStand
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	for _, card := range dealer {
		ai.count(card)
	}
	for _, hand := range hands {
		for _, card := range hand {
			ai.count(card)
		}
	}
}

func (ai *basicAI) count(card deck.Card) {
	score := blackjack.Score(card)
	switch {
	case score >= 10:
		ai.score--
	case score <= 6:
		ai.score++
	}
	ai.seen++
}

func main() {
	opts := blackjack.Options{
		Decks:           4,
		Hands:           3,
		BlackjackPayout: 1.5,
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!====================NEW GAME===================!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	game := blackjack.New(opts)

	//winnings := game.Play(blackjack.HumanAI()) //SI QUIERES QUE JUEGUE UNA PERSONA
	winnings := game.Play(&basicAI{decks: 4}) //SI QUIERES QUE JUEGUE EL AI
	fmt.Println("Final Winnings:", winnings)
}
