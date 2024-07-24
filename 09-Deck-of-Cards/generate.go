
package deck

// import "fmt"
//go:generate stringer -type=Suit,Rank

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

// The String method should be uncommented and corrected if needed.
// func (c Card) String() string {
//     if c.Suit == Joker {
//         return c.Suit.String()
//     }
//     return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
// }
func (c Card) String() string {
    if c.Suit == Joker {
        return c.Suit.String()
    }
    return c.Rank.String() + " of " + c.Suit.String() + "s"
}
