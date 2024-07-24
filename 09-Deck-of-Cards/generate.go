//go:generate stringer -type=Suit,Rank

package deck

type Suit uint8

const (
    Spade Suit = iota
    Diamond
    Club
    Heart
    Joker // This is a special case.
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
