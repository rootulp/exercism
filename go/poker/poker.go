package poker

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func BestHand(rawHands []string) (bestCards []string, err error) {
	hands, err := parseHands(rawHands)
	fmt.Printf("hands %v, err %v\n", hands, err)
	if err != nil {
		return []string{}, err
	}

	return []string{}, nil
}

type Hand struct {
	rawHand string
	cards   []Card
}

type Card struct {
	rank Rank
	suit Suit
}

type Rank string

const (
	Two   Rank = "2"
	Three Rank = "3"
	// TODO
)

type Suit string

const (
	Heart   Suit = "♡"
	Diamond Suit = "♢"
	Club    Suit = "♧"
	Spade   Suit = "♤"
)

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

func parseCards(rawHand string) (cards []Card, err error) {
	rawCards := strings.Split(rawHand, " ")
	for _, rawCard := range rawCards {
		card, err := NewCard(rawCard)
		if err != nil {
			return cards, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func NewCard(rawCard string) (card Card, err error) {
	regex, err := regexp.Compile(`(?P<rank>\d+)(?P<suit>♢|♡|♧|♤)`)
	if err != nil {
		log.Fatal(err)
	}

	if !regex.MatchString(rawCard) {
		return card, fmt.Errorf("rawCard %v does not match regex %v", rawCard, regex.String())
	}

	matches := regex.FindStringSubmatch(rawCard)
	if len(matches) != 3 {
		return card, fmt.Errorf("unexpected number of matches %v for rawCard %v", len(matches), rawCard)
	}

	_, rawRank, rawSuit := matches[0], matches[1], matches[2]
	rank, err := parseRank(rawRank)
	if err != nil {
		return card, err
	}

	suit, err := parseSuit(rawSuit)
	if err != nil {
		return card, err
	}

	return Card{
		rank: rank,
		suit: suit,
	}, nil
}

func parseRank(rawRank string) (rank Rank, err error) {
	fmt.Printf("parsing rawRank %v\n", rawRank)
	switch rawRank {
	case "2":
		return Two, nil
	case "3":
		return Three, nil
	default:
		return Two, fmt.Errorf("invalid rawRank %v", rawRank)
	}

}

func parseSuit(rawSuit string) (suit Suit, err error) {
	switch rawSuit {
	case "♡":
		return Heart, nil
	case "♢":
		return Diamond, nil
	case "♧":
		return Club, nil
	case "♤":
		return Spade, nil
	default:
		return Heart, fmt.Errorf("invalid rawSuit %v", rawSuit)
	}
}
