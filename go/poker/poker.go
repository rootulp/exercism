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
		return bestHands, err
	}

	bestHand, err := getBestHand(rawHands)
	fmt.Printf("bestHand %v\n", bestHand)
	if err != nil {
		return bestHands, err
	}

	for _, hand := range hands {
		comparison, err := bestHand.compare(hand)
		if err != nil {
			return bestHands, err
		}
		if comparison == 0 {
			bestHands = append(bestHands, hand.rawHand)
		}
	}

	return bestHands, nil
}

func getBestHand(rawHands []string) (bestHand Hand, err error) {
	hands, err := parseHands(rawHands)
	if err != nil {
		return Hand{}, err
	}

	sort.Sort(ByScore(hands))
	fmt.Printf("sortedHands %v\n", hands)

	bestHand = hands[len(hands)-1]
	return bestHand, nil
}
