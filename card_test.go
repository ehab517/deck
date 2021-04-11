package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	//Output:
	//Ace of Hearts

}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("worng number of cards in the new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	testCard := Card{Diamond, Ace}
	if cards[13] != testCard {
		t.Error("expected Ace of spades put found something else")
	}
}
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	testCard := Card{Diamond, Ace}
	if cards[13] != testCard {
		t.Error("expected Ace of spades put found something else")
	}
}

func TestJockers(t *testing.T) {
	cards := New(Jokers(3))
	counter := 0
	for _, c := range cards {
		if c.Suit == Joker {
			counter++
		}
	}

	if counter != 3 {
		t.Error("expected 3 jokers, Received: ", counter)
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}

	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("expected to filter all Twos and Threes, but some were found")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(4))

	if len(cards) != 52*4 {
		t.Error("expected a deck of ", 52*4, " cards, got a deck of ", len(cards), " cards")
	}
}
