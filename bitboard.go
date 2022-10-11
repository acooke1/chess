package chess

import (
	"math/bits"
	"strconv"
)

type bitboard uint64

func newBitBoard(m map[Square]bool) bitboard {
	s := ""
	for sq := 0; sq < numberOfSquaresInBoard; sq++ {
		if m[Square(sq)] {
			s += "1"
		} else {
			s += "0"
		}
	}
	bb, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return bitboard(bb)
}

func (b bitboard) Reverse() bitboard {
	return bitboard(bits.Reverse64(uint64(b)))
}

func (b bitboard) Occupies(sq Square) bool {
	return bits.RotateLeft(uint(b), int(sq)+1)&1 == 1
}
