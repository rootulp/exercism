package dominoes

import "fmt"

type Domino [2]int

func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 {
		return input, true
	}
	return makeChain(input, []Domino{})
}

func makeChain(remaining []Domino, chain []Domino) (result []Domino, ok bool) {
	fmt.Printf("makeChain(%v, %v)\n", remaining, chain)
	if len(remaining) == 0 {
		return chain, isValid(chain)
	}
	if len(chain) >= 2 && !isValidConnections(chain) {
		return chain, false
	}
	for i, v := range remaining {
		candidateRemaining := removeIndex(remaining, i)

		candidateChain1 := append([]Domino{}, chain...)
		candidateChain2 := append([]Domino{}, chain...)

		candidateChain1 = append(candidateChain1, v)
		candidateChain2 = append(candidateChain2, flip(v))

		if result, ok := makeChain(candidateRemaining, candidateChain1); ok {
			return result, ok
		}
		if result, ok := makeChain(candidateRemaining, candidateChain2); ok {
			return result, ok
		}
	}
	return nil, false
}

// removeIndex removes the element at index from arr
func removeIndex(arr []Domino, index int) (result []Domino) {
	result = append([]Domino{}, arr...)
	return append(result[:index], result[index+1:]...)
}

func isValid(chain []Domino) bool {
	isValid := isValidEnds(chain) && isValidConnections(chain)
	fmt.Printf("isValid(%v)=%v\n", chain, isValid)
	return isValid
}

// isValidEnds returns whether the left side of first domino == right side of
// last domino
func isValidEnds(chain []Domino) bool {
	if len(chain) == 0 {
		return true
	} else if len(chain) == 1 {
		return chain[0][0] == chain[0][1]
	} else {
		return chain[0][0] == chain[len(chain)-1][1]
	}
}

// isValidEnds returns whether the the connections between dominos are valid
func isValidConnections(chain []Domino) bool {
	if len(chain) == 0 || len(chain) == 1 {
		return true
	}
	lastRightSide := chain[0][1]
	for _, v := range chain[1:] {
		if lastRightSide != v[0] {
			return false
		}
		lastRightSide = v[1]
	}
	return true
}

func flip(d Domino) (flipped Domino) {
	return Domino{d[1], d[0]}
}
