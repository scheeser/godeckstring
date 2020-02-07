# godeckstring
Go library for encoding/decoding Hearthstone [deckstrings](https://playhearthstone.com/en-us/blog/20720853).

A great resource describing the format can be found at [https://hearthsim.info/docs/deckstrings/](https://hearthsim.info/docs/deckstrings/).

## Install

```
go get github.com/scheeser/godeckstring
```

```go
import github.com/scheeser/godeckstring
```

## Usage

### Decode

```go
package main

import (
	"fmt"

	"github.com/scheeser/godeckstring"
)

func main() {
	deckstring := "AAECAaoIBO/3AuGoA+O0A9PAAw2cArSRA7SXA8aZA7ulA8+lA9SlA9WlA/mlA7etA7mtA/6uA6qvAwA="

	decodedDeck, _ := godeckstring.Decode(deckstring)
	fmt.Printf("%+v\n", decodedDeck)
}
```

Outputs:
```go
{Name: FormatType:2 Heroes:[{DbfID:1066}] Cards:[{DbfID:48111 Copies:1} {DbfID:54369 Copies:1} {DbfID:55907 Copies:1} {DbfID:57427 Copies:1} {DbfID:284 Copies:2} {DbfID:51380 Copies:2} {DbfID:52148 Copies:2} {DbfID:52422 Copies:2} {DbfID:53947 Copies:2} {DbfID:53967 Copies:2} {DbfID:53972 Copies:2} {DbfID:53973 Copies:2} {DbfID:54009 Copies:2} {DbfID:54967 Copies:2} {DbfID:54969 Copies:2} {DbfID:55166 Copies:2} {DbfID:55210 Copies:2}]}
```

Translating from DBF ID to Hero or Card can be done using tools like [HearthstoneJSON](https://hearthstonejson.com/).

### Encode

```go
package main

import (
	"fmt"

	"github.com/scheeser/godeckstring"
)

func main() {
	deck := godeckstring.Deck{
		FormatType: godeckstring.FormatTypeStandard,
		Heroes: []godeckstring.Hero{
			godeckstring.Hero{
				DbfID: 1066, // Thrall
			},
		},
		Cards: []godeckstring.Card{
			godeckstring.Card{
				DbfID:  52148, // Mutate
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  54369, // Corrupt the Waters
				Copies: 1,
			},
			godeckstring.Card{
				DbfID:  52422, // Sludge Slurper
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  51380, // EVIL Cable Rat
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  54009, // EVIL Totem
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  284, // Novice Engineer
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  53947, // Questing Explorer
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  53967, // Sandstorm Elemental
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  53973, // Weaponized Wasp
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  55210, // Devoted Maniac
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  54969, // Corrupt Elementalist
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  54967, // Dragon's Pack
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  55166, // Shield of Galakrond
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  55907, // Kronx Dragonhoof
				Copies: 1,
			},
			godeckstring.Card{
				DbfID:  57427, // Galakrond, The Tempest
				Copies: 1,
			},
			godeckstring.Card{
				DbfID:  53972, // Mogu Fleshshaper
				Copies: 2,
			},
			godeckstring.Card{
				DbfID:  48111, // Shudderwock
				Copies: 1,
			},
		},
	}

	deckstring, _ := godeckstring.Encode(deck)

	fmt.Println(deckstring)
}
```

Outputs:
```
AAECAaoIBO/3AuGoA+O0A9PAAw2cArSRA7SXA8aZA7ulA8+lA9SlA9WlA/mlA7etA7mtA/6uA6qvAwA=
```
