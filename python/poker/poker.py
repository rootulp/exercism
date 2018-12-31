STRAIGHT_FLUSH = 800
FOUR_OF_A_KIND = 700
FULL_HOUSE = 600
FLUSH = 500
STRAIGHT = 400
THREE_OF_A_KIND = 300
TWO_PAIR = 200
ONE_PAIR = 100
HIGH_CARD = 0


class Poker:
    def __init__(self, hands):
        self.hands = [Hand(hand) for hand in hands]

    def best_hand(self):
        return [hand.hand for hand, score in list(self.scores().items()) if score ==
                self.best_score()]

    def best_score(self):
        return max(self.scores(), key=self.scores().get).score()

    def scores(self):
        return {hand: hand.score() for hand in self.hands}


class Hand:
    def __init__(self, hand):
        self.hand = hand
        self.cards = [Card(card) for card in hand]
        self.ranks = [card.rank for card in self.cards]
        self.suits = [card.suit for card in self.cards]

    def score(self):
        if self.straight_flush():
            return STRAIGHT_FLUSH + self.high_card()
        elif self.four_of_a_kind():
            return FOUR_OF_A_KIND + self.rank_of_card_with_highest_occurence()
        elif self.full_house():
            return FULL_HOUSE + self.rank_of_card_with_highest_occurence()
        elif self.flush():
            return FLUSH + self.high_card()
        elif self.straight():
            return STRAIGHT + self.high_card()
        elif self.three_of_a_kind():
            return THREE_OF_A_KIND + self.rank_of_card_with_highest_occurence()
        elif self.two_pair():
            return TWO_PAIR + self.high_card()
        elif self.one_pair():
            return ONE_PAIR + self.rank_of_card_with_highest_occurence()
        else:
            return self.high_card()

    def straight_flush(self):
        return self.straight() and self.flush()

    def straight(self):
        return (list(range(min(self.ranks), max(self.ranks) + 1)) ==
                sorted(self.ranks))

    def flush(self):
        return len(set(self.suits)) == 1

    def four_of_a_kind(self):
        return max(self.rank_occurences().values()) == 4

    def full_house(self):
        return self.three_of_a_kind() and self.one_pair()

    def three_of_a_kind(self):
        return 3 in list(self.rank_occurences().values())

    def two_pair(self):
        return 2 == list(self.rank_occurences().values()).count(2)

    def one_pair(self):
        return 2 in list(self.rank_occurences().values())

    def high_card(self):
        return max(self.ranks)

    def rank_occurences(self):
        return {rank: self.ranks.count(rank) for rank in self.ranks}

    def rank_of_card_with_highest_occurence(self):
        return max(self.rank_occurences(), key=self.rank_occurences().get)


class Card:
    FACE_CARDS = {"T": 10,
                  "J": 11,
                  "Q": 12,
                  "K": 13,
                  "A": 14}

    def __init__(self, inp):
        self.rank = self.numberify_face_cards(inp[0])
        self.suit = inp[1]

    def __str__(self):
        return str(self.rank) + self.suit

    @classmethod
    def numberify_face_cards(cls, rank):
        if rank in list(cls.FACE_CARDS.keys()):
            return cls.FACE_CARDS[rank]
        return int(rank)


def poker(inp):
    return Poker(inp).best_hand()
