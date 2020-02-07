package godeckstring

// Hero Represents the hero. Decks are associated to the hero and not the underlying class.
type Hero struct {
	DbfID uint64
}

// Card Identifies the card by DBF ID and the number of copies of the card in the deck.
type Card struct {
	DbfID  uint64
	Copies int
}

// Deck The contents of a deck. Format, heros, cards, etc.
type Deck struct {
	Name       string
	FormatType uint64
	Heroes     []Hero
	Cards      []Card
}
