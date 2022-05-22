package poker

import (
	"fmt"
	"log"
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

func (h Hand) Ranks() (ranks []Rank) {
	for _, card := range h.cards {
		ranks = append(ranks, card.rank)
	}
	return ranks
}

// compare returns 0 if both hands have equal score
// compare returns a positive number if the score of h is less than the score of o
// compare returns a negative number if the score of h is greater than the score of o
func (h Hand) compare(o Hand) (int, error) {
	if h.handType() == o.handType() {
		return h.tiebreak(o)
	}
	return int(h.handType()) - int(o.handType()), nil
}

// compare returns 0 if both hands have equal score
// compare returns a positive number if the score of h is less than the score of o
// compare returns a negative number if the score of h is greater than the score of o
func (h Hand) tiebreak(o Hand) (int, error) {
	if h.handType() != o.handType() {
		log.Fatalf("tiebreak invoked with unequal hand types %v and %v", h, o)
	}

	switch h.handType() {
	case FourOfKind:
		comparison, err := h.compareByHighestQuad(o)
		if err != nil {
			return 0, err
		}
		if comparison != 0 {
			return comparison, nil
		}

		comparison, err = h.compareByHighestKicker(o)
		if err != nil {
			return 0, err
		}
		if comparison != 0 {
			return comparison, nil
		}
		return 0, nil
	case FullHouse:
		comparison, err := h.compareByHighestTriplet(o)
		if err != nil {
			return 0, err
		}
		if comparison != 0 {
			return comparison, nil
		}

		comparison, err = h.compareByHighestPair(o)
		if err != nil {
			return 0, err
		}
		if comparison != 0 {
			return comparison, nil
		}

		return 0, nil
	case Straight:
		if h.isStraightStartingWithAce() && o.isStraightStartingWithAce() {
			return 0, nil
		} else if h.isStraightStartingWithAce() {
			return -1, nil
		} else if o.isStraightStartingWithAce() {
			return 1, nil
		}
		return h.compareByHighestCard(o), nil
	default:
		return h.compareByHighestCard(o), nil
	}
}

func (h Hand) compareByHighestCard(o Hand) int {
	descendingRanksH := h.descendingRanks()
	descendingRanksO := o.descendingRanks()

	for i := range descendingRanksH {
		rankH := descendingRanksH[i]
		rankO := descendingRanksO[i]
		if rankH != rankO {
			return int(rankH) - int(rankO)
		}
	}
	return 0
}

func (h Hand) compareByHighestQuad(o Hand) (int, error) {
	quadH, err := h.quadRank()
	if err != nil {
		return 0, err
	}

	quadO, err := o.quadRank()
	if err != nil {
		return 0, err
	}

	return int(quadH) - int(quadO), nil
}

func (h Hand) compareByHighestKicker(o Hand) (int, error) {
	kickerH, err := h.kicker()
	if err != nil {
		return 0, err
	}

	kickerO, err := o.kicker()
	if err != nil {
		return 0, err
	}

	return int(kickerH) - int(kickerO), nil
}

func (h Hand) compareByHighestTriplet(o Hand) (int, error) {
	tripletH, err := h.tripletRank()
	if err != nil {
		return 0, err
	}

	tripletO, err := o.tripletRank()
	if err != nil {
		return 0, err
	}

	return int(tripletH) - int(tripletO), nil
}

func (h Hand) compareByHighestPair(o Hand) (int, error) {
	pairH, err := h.pairRank()
	if err != nil {
		return 0, err
	}

	pairO, err := o.pairRank()
	if err != nil {
		return 0, err
	}

	return int(pairH) - int(pairO), nil
}

func (h Hand) quadRank() (rank Rank, err error) {
	rankToOccurences := h.getRankToOccurences()
	for rank, occurences := range rankToOccurences {
		if occurences == 4 {
			return rank, nil
		}
	}
	return rank, fmt.Errorf("hand %v does not have a quad", h)
}

func (h Hand) tripletRank() (rank Rank, err error) {
	rankToOccurences := h.getRankToOccurences()
	for rank, occurences := range rankToOccurences {
		if occurences == 3 {
			return rank, nil
		}
	}
	return rank, fmt.Errorf("hand %v does not have a triplet", h)
}

func (h Hand) pairRank() (rank Rank, err error) {
	rankToOccurences := h.getRankToOccurences()
	for rank, occurences := range rankToOccurences {
		if occurences == 2 {
			return rank, nil
		}
	}
	return rank, fmt.Errorf("hand %v does not have a pair", h)
}

func (h Hand) kicker() (rank Rank, err error) {
	rankToOccurences := h.getRankToOccurences()
	for rank, occurences := range rankToOccurences {
		if occurences == 1 {
			return rank, nil
		}
	}
	return rank, fmt.Errorf("hand %v does not have a kicker", h)
}

func (h Hand) descendingRanks() (ranks []Rank) {
	for _, card := range h.cards {
		ranks = append(ranks, card.rank)
	}

	sort.Sort(sort.Reverse(ByRank(ranks)))
	return ranks
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
	return fmt.Sprintf("handType %v with cards %v", h.handType(), h.cards)
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
	sort.Ints(occurences)
	fullHouse := []int{2, 3}
	return reflect.DeepEqual(occurences, fullHouse)
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
	if h.isStraightStartingWithAce() {
		return true
	}

	ranks := h.Ranks()
	sort.Sort(ByRank(ranks))
	for i, rank := range ranks {
		if i == 0 {
			continue
		}
		if int(ranks[i-1]) != int(rank)-1 {
			return false
		}
	}
	return true
}

func (h Hand) isStraightStartingWithAce() bool {
	ranks := h.Ranks()
	sort.Sort(ByRank(ranks))
	straightStartingWithAce := []Rank{Two, Three, Four, Five, Ace}
	return reflect.DeepEqual(ranks, straightStartingWithAce)
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
	sort.Ints(occurences)
	twoPairs := []int{1, 2, 2}
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
	comparison, err := handA.compare(handB)
	if err != nil {
		log.Fatal(err)
	}
	return comparison < 0
}

func getValues(rankToOccurences map[Rank]int) (values []int) {
	for _, v := range rankToOccurences {
		values = append(values, v)
	}
	return values
}
