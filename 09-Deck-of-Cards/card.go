// go:generate stringer -type=Suit,Rank

package deck

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

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

type Card struct {
	Suit 
	Rank
}

func (c Card) String() string {
	return "Ace of Heart"
}