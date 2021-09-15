// Package chessboard provides Chessboard solution.
package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[byte]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank byte) int {
	n := 0
	for _, occupied := range cb[rank] {
		if occupied {
			n++
		}
	}
	return n
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) int {
	n := 0
	for _, rank := range cb {
		if file > len(rank) {
			break
		}
		if rank[file-1] {
			n++
		}
	}
	return n
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() int {
	return len(cb) * len(cb['A'])
}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() int {
	n := 0
	for _, rank := range cb {
		for _, occupied := range rank {
			if occupied {
				n++
			}
		}
	}
	return n
}
