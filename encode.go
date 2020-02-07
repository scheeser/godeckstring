package godeckstring

import (
	"encoding/base64"
	"encoding/binary"
	"sort"
)

// Encode takes the provided deck struct and encodes it as a valid Hearthstone deckstring.
func Encode(deck Deck) (string, error) {
	// The header is a known length.
	var headerWriteLength = 3
	var header = make([]uint64, headerWriteLength)
	header[0] = 0
	header[1] = Version
	header[2] = deck.FormatType

	// Sort the heroes
	sort.SliceStable(deck.Heroes, func(x int, y int) bool {
		return deck.Heroes[x].DbfID < deck.Cards[y].DbfID
	})

	// The hero will be the count of heroes + 1 for the length integer...most likely 1.
	numHeroes := len(deck.Heroes)
	heroesWriteLength := numHeroes + 1
	heroes := make([]uint64, heroesWriteLength)
	heroes[0] = uint64(numHeroes)

	for i, hero := range deck.Heroes {
		heroes[i+1] = hero.DbfID
	}

	// Sort cards by ascending DBF ID first. This should ensure future slices are built in sorted order.
	sort.SliceStable(deck.Cards, func(x int, y int) bool {
		return deck.Cards[x].DbfID < deck.Cards[y].DbfID
	})

	// Group cards by the number of copies
	var singleCopyCards, twoCopyCards []uint64
	var nCopyCards [][2]uint64
	for _, card := range deck.Cards {
		if card.Copies == 1 {
			singleCopyCards = append(singleCopyCards, card.DbfID)
			continue
		}
		if card.Copies == 2 {
			twoCopyCards = append(twoCopyCards, card.DbfID)
			continue
		}
		if card.Copies > 2 {
			pair := [2]uint64{card.DbfID, uint64(card.Copies)}
			nCopyCards = append(nCopyCards, pair)
		}
	}

	var cards []uint64
	// Add the single and two copy cards.
	cards = append(cards, uint64(len(singleCopyCards)))
	cards = append(cards, singleCopyCards...)
	cards = append(cards, uint64(len(twoCopyCards)))
	cards = append(cards, twoCopyCards...)

	// Add the n-copy cards.
	cards = append(cards, uint64(len(nCopyCards)))
	for _, pair := range nCopyCards {
		// The pair should always be of the form (dbf id, copies)
		cards = append(cards, pair[0], pair[1])
	}

	// Join the header, hero and cards slices
	totalToWrite := make([]uint64, headerWriteLength+heroesWriteLength+len(cards))
	copy(totalToWrite, header)
	copy(totalToWrite[headerWriteLength:], heroes)
	copy(totalToWrite[(headerWriteLength+heroesWriteLength):], cards)

	// Ensure byte buffer has enough capapcity to write the entire deck.
	byteBuffer := make([]byte, len(totalToWrite)*binary.MaxVarintLen64)
	var totalBytesWritten int
	for _, item := range totalToWrite {
		// The new byte to write needs appended to the end of the buffer. Offset from the last bytes that were written.
		bytesWritten := binary.PutUvarint(byteBuffer[totalBytesWritten:], item)
		totalBytesWritten = totalBytesWritten + bytesWritten
	}

	deckstring := base64.StdEncoding.EncodeToString(byteBuffer[:totalBytesWritten])

	return deckstring, nil
}
