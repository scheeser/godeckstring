package godeckstring

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
)

// Decode reference from: https://hearthsim.info/docs/deckstrings/

// Decode Takes the provided codified deck and translates the contents to a struct.
func Decode(deckcode string) (deck Deck, funcErr error) {
	// TODO: Should probably deal with lines in the deckcode that start wtih #
	// TODO: "The first line beginning with ### preceding the deck string, which will be used as the deck name if available."

	decodeBytes, err := base64.StdEncoding.DecodeString(deckcode)
	if err != nil {
		funcErr = fmt.Errorf("decode error: %s", err)
		return
	}

	buf := bytes.NewReader(decodeBytes)

	// Parse reserve byte
	reserve, err := binary.ReadUvarint(buf)

	if err != nil {
		funcErr = fmt.Errorf("Problem reading reserved byte: %s", err)
		return
	}

	if reserve != ReserveByte {
		funcErr = fmt.Errorf("The reserve byte %d != %d", reserve, ReserveByte)
		return
	}

	// Parse version
	ver, err := binary.ReadUvarint(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading version: %s", err)
		return
	}

	if ver != Version {
		funcErr = fmt.Errorf("The version %d != %d", ver, Version)
		return
	}

	// Parse format
	format, err := binary.ReadUvarint(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading format: %s", err)
		return
	}

	_, err = GetFormatType(format)
	if err != nil {
		return deck, err
	}

	deck.FormatType = format

	// Parse the hero
	heroIDs, err := readLengthArrayPair(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading hero: %s", err)
		return
	}

	for _, dbfID := range heroIDs {
		hero := Hero{
			DbfID: dbfID,
		}

		deck.Heroes = append(deck.Heroes, hero)
	}

	// Parse single copy cards
	singleCopyIDs, err := readLengthArrayPair(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading single copy cards: %s", err)
		return
	}

	deck.Cards = append(deck.Cards, buildCards(singleCopyIDs, 1)...)

	// Parse double copy cards
	doubleCopyIDs, err := readLengthArrayPair(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading double copy cards: %s", err)
		return
	}

	deck.Cards = append(deck.Cards, buildCards(doubleCopyIDs, 2)...)

	// Deal with n-copy cards. Without a valid deckstring from the game it's hard to ensure if this implementation is correct.
	nCopyCount, err := binary.ReadUvarint(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading n-copy pair count: %s", err)
		return
	}

	if nCopyCount == 0 {
		return deck, nil
	}

	// For each n-copy cards read the pair of dbf id and number of copies
	for i := 0; i < int(nCopyCount); i++ {
		dbfID, err := binary.ReadUvarint(buf)
		if err != nil {
			funcErr = fmt.Errorf("Problem reading DBF ID with index %d: %s", i, err)
			return
		}

		copies, err := binary.ReadUvarint(buf)
		if err != nil {
			funcErr = fmt.Errorf("Problem reading number of copies for DBF ID %d at index %d: %s", dbfID, i, err)
			return
		}

		deck.Cards = append(deck.Cards, buildCard(dbfID, int(copies)))
	}

	return deck, nil
}

func readLengthArrayPair(buf io.ByteReader) (dbfIDs []uint64, funcErr error) {
	count, err := binary.ReadUvarint(buf)
	if err != nil {
		funcErr = fmt.Errorf("Problem reading count: %s", err)
		return
	}

	if count == 0 {
		return
	}

	dbfIDs = make([]uint64, count, count)
	for i := 0; i < int(count); i++ {
		dbfID, err := binary.ReadUvarint(buf)
		if err != nil {
			funcErr = fmt.Errorf("Problem reading DBF ID with index %d: %s", i, err)
			return
		}

		dbfIDs[i] = dbfID
	}

	return
}

func buildCards(dbfIDs []uint64, copies int) []Card {
	numCards := len(dbfIDs)
	cards := make([]Card, numCards, numCards)

	for i, dbfID := range dbfIDs {
		cards[i] = buildCard(dbfID, copies)
	}

	return cards
}

func buildCard(dbfID uint64, copies int) Card {
	return Card{
		DbfID:  dbfID,
		Copies: copies,
	}
}
