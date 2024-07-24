package deck

import (
	"sort"
	"time"
	"math/rand"
)


func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	// for each suits
	// for each rank
	// add card{suits, rank} to cards
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i,j int) bool ) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i,j int)bool {
	return func(i,j int)bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}


func absRank(c Card) int {
	return int(c.Suit)*int(c.Rank) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i,j := range perm {
		ret[i] = cards[j]
	}
	return ret
}