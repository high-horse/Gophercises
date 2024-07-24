package deck

// import (
// 	"fmt"
// 	"testing"
// )

// func ExampleCard() {
// 	fmt.Println(Card{Rank: Ace, Suit: Heart})
// 	fmt.Println(Card{Rank: Two, Suit: Spade})
// 	fmt.Println(Card{Rank: Nine, Suit: Diamond})
// 	fmt.Println(Card{Rank: Jack, Suit: Club})
// 	fmt.Println(Card{Suit: Joker})



// 	// Output: 
// 	// Ace of Hearts
// 	// Two of Spades
// 	// Nine of Diamonds
// 	// Jack of Clubs
// 	// Joker
// }

// func TestNew(t *testing.T) {
// 	cards := New()
// 	// 13 ranks * 4 suits
// 	if len(cards) != 13*4 {
// 		t.Error("Expected 13*4 cards, but got", len(cards))
// 	}
// }

// func TestDefaultSort(t *testing.T) {
// 	cards := New(DefaultSort)
// 	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
// 		t.Error("Expected Ace of Spades as first card, but got", cards[0])
// 	}
// }

// func TestSort(t *testing.T) {
// 	cards := New(Sort(Less))
// 	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
// 		t.Error("Expected Ace of Spades as first card, but got", cards[0])
// 	}
// }