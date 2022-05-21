package poker

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Rank int
type Suit string
type Card struct {
	rank Rank
	suit Suit
}

const (
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

const (
	Heart   Suit = "♡"
	Diamond Suit = "♢"
	Club    Suit = "♧"
	Spade   Suit = "♤"
)

func NewCard(rawCard string) (card Card, err error) {
	regex, err := regexp.Compile(`(?P<rank>\d+|J|Q|K)(?P<suit>♢|♡|♧|♤)$`)
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

func parseCards(rawHand string) (cards []Card, err error) {
	rawCards := strings.Split(rawHand, " ")
	for _, rawCard := range rawCards {
		card, err := NewCard(rawCard)
		if err != nil {
			return cards, err
		}
		cards = append(cards, card)
	}
	if len(cards) != 5 {
		return cards, fmt.Errorf("incorrect number of cards %v", len(cards))
	}
	return cards, nil
}

func parseRank(rawRank string) (rank Rank, err error) {
	switch rawRank {
	case "2":
		return Two, nil
	case "3":
		return Three, nil
	case "4":
		return Four, nil
	case "5":
		return Five, nil
	case "6":
		return Six, nil
	case "7":
		return Seven, nil
	case "8":
		return Eight, nil
	case "9":
		return Nine, nil
	case "10":
		return Ten, nil
	case "J":
		return Jack, nil
	case "Q":
		return Queen, nil
	case "K":
		return King, nil
	case "A":
		return Ace, nil
	default:
		return 0, fmt.Errorf("invalid rawRank %v", rawRank)
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

type ByRank []Rank

func (r ByRank) Len() int               { return len(r) }
func (r ByRank) Swap(a int, b int)      { r[a], r[b] = r[b], r[a] }
func (r ByRank) Less(a int, b int) bool { return r[a] < r[b] }
