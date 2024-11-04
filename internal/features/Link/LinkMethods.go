package featureLink

import "fmt"

var allowedChars = [62]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func CreateNewLink(url string) (Link, error) {
	request := newLinkRequest()

	// TODO
	// For fine concurrency work it's must go to some queue and queue create codes
	// what based on their position and gave correct incremental name
	// But for local live and dev isn't important
	if err := request.findLastItem(); err != nil {
		return nil, fmt.Errorf("error search for last link: %v", err)
	}

	link, ok := request.getFirst()
	if !ok {
		return nil, fmt.Errorf("error getting first link")
	}

	code := getNewCode(link.GetId())

	if err := request.createLink(url, code); err != nil {
		return nil, fmt.Errorf("error creating new link: %v", err)
	}

	link, ok = request.getFirst()
	if !ok {
		return nil, fmt.Errorf("error getting first link")
	}

	return link, nil
}

func getNewCode(id uint) string {
	//groups := id / 62
	last := id % 62

	fmt.Println(last)

	return allowedChars[last]
}
