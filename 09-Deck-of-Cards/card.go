//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // This is special case.
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota // Skip 0
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}


func New() []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	// for each suits
	// for each rank
	// add card{suits, rank} to cards

	return cards
}