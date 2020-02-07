package godeckstring

import (
	"fmt"
)

// ReserveByte starts the string and is always 0.
const ReserveByte = 0

// Version currently is always 1.
const Version = 1

// The set of existing Formats: Standard, Wild and Unknown.
const (
	FormatTypeUnknown  = 0
	FormatTypeWild     = 1
	FormatTypeStandard = 2
)

var formatTypes = map[int]string{
	FormatTypeUnknown:  "Unknown",
	FormatTypeWild:     "Wild",
	FormatTypeStandard: "Standard",
}

// GetFormatType Convert numeric format type to a more friendly string. If the format type is know one of the known values an error will be returned.
func GetFormatType(ft uint64) (string, error) {
	format, found := formatTypes[int(ft)]

	if found {
		return format, nil
	}

	return "", fmt.Errorf("%d is an invalid format type", ft)
}
