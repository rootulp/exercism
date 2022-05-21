package poker

import "fmt"

type Hand struct {
	rawHand string
	cards   []Card
}

func NewHand(rawHand string) (hand Hand, err error) {
	cards, err := parseCards(rawHand)
	if err != nil {
		return Hand{}, err
	}
	return Hand{
		rawHand: rawHand,
		cards:   cards,
	}, nil

}

func parseHands(rawHands []string) (hands []Hand, err error) {
	for _, rawHand := range rawHands {
		hand, err := NewHand(rawHand)
		if err != nil {
			return []Hand{}, err
		}
		hands = append(hands, hand)
	}
	return hands, nil
}

func (h Hand) value() int {
	// TODO
	return 0
}

func (h Hand) String() string {
	return fmt.Sprintf("rawHand %v cards %v value %v\n", h.rawHand, h.cards, h.value())
}

type ByScore []Hand

func (s ByScore) Len() int          { return len(s) }
func (s ByScore) Swap(a int, b int) { s[a], s[b] = s[b], s[a] }
func (s ByScore) Less(a int, b int) bool {
	handA := s[a]
	handB := s[b]

	return handA.value() < handB.value()
}
