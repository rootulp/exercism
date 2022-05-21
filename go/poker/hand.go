package poker

import (
	"fmt"
	"reflect"
	"sort"
)

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

func (h Hand) highestCard() int {
	ranks := []Rank{}
	for _, card := range h.cards {
		ranks = append(ranks, card.rank)
	}

	sort.Sort(ByRank(ranks))
	fmt.Printf("sortedRanks %v\n", ranks)
	highestRank := ranks[len(ranks)-1]

	return int(highestRank)
}

func (h Hand) handType() HandType {
	switch {
	case h.isStraightFlush():
		return StraightFlush
	case h.isFourOfKind():
		return FourOfKind
	case h.isFullHouse():
		return FullHouse
	case h.isFlush():
		return Flush
	case h.isStraight():
		return Straight
	case h.isThreeOfKind():
		return ThreeOfKind
	case h.isTwoPairs():
		return TwoPairs
	case h.isTwoOfKind():
		return Pair
	default:
		return Nothing
	}
}

func (h Hand) String() string {
	return fmt.Sprintf("rawHand %v cards %v handType %v highestCard %v\n", h.rawHand, h.cards, h.handType(), h.highestCard())
}

func (h Hand) isStraightFlush() bool {
	return h.isStraight() && h.isFlush()
}

func (h Hand) isFourOfKind() bool {
	rankToOccurences := h.getRankToOccurences()
	for _, occurences := range rankToOccurences {
		if occurences == 4 {
			return true
		}
	}
	return false
}

func (h Hand) isFullHouse() bool {
	rankToOccurences := h.getRankToOccurences()
	occurences := getValues(rankToOccurences)
	fullHouseA := []int{2, 3}
	fullHouseB := []int{3, 2}
	return reflect.DeepEqual(occurences, fullHouseA) || reflect.DeepEqual(occurences, fullHouseB)
}

func (h Hand) isFlush() bool {
	suitToOccurences := h.getSuitToOccurences()
	for _, occurences := range suitToOccurences {
		if occurences != 5 {
			return false
		}
	}
	return true
}

func (h Hand) isStraight() bool {
	// TODO
	return false
}

func (h Hand) isThreeOfKind() bool {
	rankToOccurences := h.getRankToOccurences()
	for _, occurences := range rankToOccurences {
		if occurences == 3 {
			return true
		}
	}
	return false
}

func (h Hand) isTwoPairs() bool {
	rankToOccurences := h.getRankToOccurences()
	occurences := getValues(rankToOccurences)
	twoPairs := []int{2, 2, 1}
	return reflect.DeepEqual(occurences, twoPairs)
}

func (h Hand) isTwoOfKind() bool {
	rankToOccurences := h.getRankToOccurences()
	for _, occurences := range rankToOccurences {
		if occurences == 2 {
			return true
		}
	}
	return false
}

func (h Hand) getRankToOccurences() (occurences map[Rank]int) {
	occurences = map[Rank]int{}
	for _, card := range h.cards {
		if _, ok := occurences[card.rank]; !ok {
			occurences[card.rank] = 0
		}
		occurences[card.rank] += 1
	}
	return occurences
}

func (h Hand) getSuitToOccurences() (occurences map[Suit]int) {
	occurences = map[Suit]int{}
	for _, card := range h.cards {
		if _, ok := occurences[card.suit]; !ok {
			occurences[card.suit] = 0
		}
		occurences[card.suit] += 1
	}
	return occurences
}

type ByScore []Hand

func (s ByScore) Len() int          { return len(s) }
func (s ByScore) Swap(a int, b int) { s[a], s[b] = s[b], s[a] }
func (s ByScore) Less(a int, b int) bool {
	handA := s[a]
	handB := s[b]

	if handA.handType() != handB.handType() {
		return handA.handType() < handB.handType()
	} else {
		return handA.highestCard() < handB.highestCard()
	}
}

func getValues(rankToOccurences map[Rank]int) (values []int) {
	for _, v := range rankToOccurences {
		values = append(values, v)
	}
	return values
}
