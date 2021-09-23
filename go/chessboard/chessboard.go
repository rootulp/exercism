package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank = []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	for _, square := range cb[rank] {
		if square {
			ret += 1
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	if file > 8 || file < 1 {
		return 0
	}
	for _, rank := range cb {
		if rank[file-1] {
			ret += 1
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	for _, rank := range cb {
		for range rank {
			ret += 1
		}
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for _, rank := range cb {
		for _, square := range rank {
			if square {
				ret += 1
			}
		}
	}
	return ret
}
