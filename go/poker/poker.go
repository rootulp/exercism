package poker

import (
	"fmt"
	"sort"
)

type HandType int

const (
	Nothing HandType = iota
	Pair
	TwoPairs
	ThreeOfKind
	Straight
	Flush
	FullHouse
	FourOfKind
	StraightFlush
)

func BestHand(rawHands []string) (bestHands []string, err error) {
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
	hands, bestHand := hands[:len(hands)-1], hands[len(hands)-1]
	bestHands = append(bestHands, bestHand.rawHand)

	for _, hand := range hands {
		if bestHand.compare(hand) == 0 {
			bestHands = append(bestHands, hand.rawHand)
		}
	}

	return bestHands, nil
}
