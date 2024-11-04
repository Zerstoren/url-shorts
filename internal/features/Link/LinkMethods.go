package featureLink

import (
	"fmt"
	"math"
	"slices"
)

var allowedChars = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func CreateNewLink(url string) (Link, error) {
	request := newLinkRequest()

	if err := request.createLink(url); err != nil {
		return nil, fmt.Errorf("error creating new link: %v", err)
	}

	link, ok := request.GetFirst()
	if !ok {
		return nil, fmt.Errorf("error getting first link")
	}

	return link, nil
}

func GetLinkByCode(code string) (string, error) {
	request := newLinkRequest()

	if err := request.getByCode(code); err != nil {
		return "", fmt.Errorf("error getting link by code: %v", err)
	}

	link, ok := request.GetFirst()
	if !ok {
		return "", fmt.Errorf("error getting first link")
	}

	return link.GetTarget(), nil
}

func codeFromId(id uint) string {
	if id == 0 {
		return allowedChars[0]
	}

	code := ""

	for id > 0 {
		remainder := id % 62

		code = allowedChars[remainder] + code
		id = id / 62
	}

	return code
}

func idFromCode(code string) uint {
	id := uint(0)

	for i, char := range code {
		index := slices.Index(allowedChars, string(char))
		if index == -1 {
			return 0
		}

		remainder := uint(index)

		id += remainder * uint(math.Pow(62, float64(len(code)-i-1)))
	}

	return id
}
