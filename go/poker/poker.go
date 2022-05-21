package poker

import (
	"fmt"
	"sort"
)

type Points int

const (
	Nothing Points = iota
	Pair
	TwoPairs
	ThreeOfKind
	Straight
	Flush
	FullHouse
	FourOfKind
	StraightFlush
)

func BestHand(rawHands []string) (bestCards []string, err error) {
	hands, err := parseHands(rawHands)
	if err != nil {
		return []string{}, err
	}

	// fmt.Printf("hands %v, err %v\n", hands, err)
	if len(hands) == 1 {
		return []string{hands[0].rawHand}, nil
	}

	sort.Sort(ByScore(hands))
	fmt.Printf("sortedHands %v\n", hands)
	bestHand := hands[len(hands)-1]

	return []string{bestHand.rawHand}, nil
}
