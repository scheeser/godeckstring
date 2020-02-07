package godeckstring

import (
	"testing"
)

// Note: It appears the PlayHearthstone.com deckcode follows different sorting of cards. Without testing extensively it appears that primary sort is mana cost and then dbf id. This is based on the deck below being encoded as: 'AAECAaoIBOGoA+O0A9PAA+/3Ag20lwPGmQO0kQP5pQOcArulA8+lA9WlA6qvA7mtA7etA/6uA9SlAwA='. Building the list in the mobile client yields the expected code.
func TestMixedEncode(t *testing.T) {

	encodedDeck, _ := Encode(mixedDeck)

	if encodedDeck != mixedDeckcode {
		t.Errorf("expected deckcode '%s' got '%s'.", mixedDeckcode, encodedDeck)
	}
}

func TestTwoCopyEncode(t *testing.T) {
	encodedDeck, _ := Encode(twoCopyDeck)

	if encodedDeck != twoCopyDeckcode {
		t.Errorf("expected deckcode '%s' got '%s'.", twoCopyDeckcode, encodedDeck)
	}
}

func TestHighlanderEncode(t *testing.T) {
	encodedDeck, _ := Encode(highlanderDeck)

	if encodedDeck != highlanderDeckcode {
		t.Errorf("expected deckcode '%s' got '%s'.", highlanderDeckcode, encodedDeck)
	}
}

func TestNCopyEncode(t *testing.T) {
	encodedDeck, _ := Encode(nCopyDeck)

	if encodedDeck != nCopyDeckcode {
		t.Errorf("expected deckcode '%s' got '%s'.", nCopyDeckcode, encodedDeck)
	}
}

func TestSingleCardEncode(t *testing.T) {
	encodedDeck, _ := Encode(singleCardDeck)

	if encodedDeck != singleCardDeckcode {
		t.Errorf("expected deckcode '%s' got '%s'.", singleCardDeckcode, encodedDeck)
	}
}
