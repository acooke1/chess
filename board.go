package chess

type Board struct {
	bbWhiteKing     bitboard
	bbWhiteQueen    bitboard
	bbWhiteRook     bitboard
	bbWhiteKnight   bitboard
	bbWhiteBishop   bitboard
	bbWhitePawn     bitboard
	bbBlackKing     bitboard
	bbBlackQueen    bitboard
	bbBlackRook     bitboard
	bbBlackKnight   bitboard
	bbBlackBishop   bitboard
	bbBlackPawn     bitboard
	whiteOccupies   bitboard
	blackOccupies   bitboard
	emptySquares    bitboard
	whiteKingSquare Square
	blackKingSquare Square
}

func newBoard(m map[Square]Piece) *Board {
	b := &Board{}
	for _, p1 := range allPieces {
		bm := map[Square]bool{}
		for sq, p2 := range m {
			if p1 == p2 {
				bm[sq] = true
			}
		}
		bb := newBitBoard(bm)
		b.setForPiece(p1, bb)
	}
	b.calculateConvenienceBB(nil)
	return b
}

func (b *Board) generateMapping() map[Square]Piece {
	m := map[Square]Piece{}
	for i := 0; i < numberOfSquaresInBoard; i++ {
		sq := Square(i)
		m[sq] = b.getPiece(sq)
	}
	return m
}

func (b *Board) getPiece(sq Square) Piece {
	for _, p := range allPieces {
		bb := b.bbForPiece(p)
		if bb.Occupies(sq) {
			return p
		}
	}

	return NoPiece
}

func (b *Board) calculateConvenienceBB(move *Move) {
	whiteOccupies := b.bbWhiteKing | b.bbWhiteQueen | b.bbWhiteRook | b.bbWhiteKnight | b.bbWhiteBishop | b.bbWhitePawn
	blackOccupies := b.bbBlackKing | b.bbBlackQueen | b.bbBlackRook | b.bbBlackKnight | b.bbBlackBishop | b.bbBlackPawn
	emptySquares := ^(whiteOccupies | blackOccupies)
	b.whiteOccupies = whiteOccupies
	b.blackOccupies = blackOccupies
	b.emptySquares = emptySquares

	if move == nil {
		b.whiteKingSquare = NoSquare
		b.blackKingSquare = NoSquare

		for i := 0; i < numberOfSquaresInBoard; i++ {
			sq := Square(i)
			if b.whiteKingSquare == NoSquare && b.bbWhiteKing.Occupies(sq) {
				b.whiteKingSquare = sq
			} else if b.blackKingSquare == NoSquare && b.bbBlackKing.Occupies(sq) {
				b.blackKingSquare = sq
			}
		}
	} else if b.whiteKingSquare == move.sq1 {
		b.whiteKingSquare = move.sq2
	} else if b.blackKingSquare == move.sq1 {
		b.blackKingSquare = move.sq2
	}
}

func (b *Board) bbForPiece(p Piece) bitboard {
	switch p {
	case WhiteKing:
		return b.bbWhiteKing
	case WhiteQueen:
		return b.bbWhiteQueen
	case WhiteRook:
		return b.bbWhiteRook
	case WhiteKnight:
		return b.bbWhiteKnight
	case WhiteBishop:
		return b.bbWhiteBishop
	case WhitePawn:
		return b.bbWhitePawn
	case BlackKing:
		return b.bbBlackKing
	case BlackQueen:
		return b.bbBlackQueen
	case BlackRook:
		return b.bbBlackRook
	case BlackKnight:
		return b.bbBlackKnight
	case BlackBishop:
		return b.bbBlackBishop
	case BlackPawn:
		return b.bbBlackPawn
	default:
		panic("No bitboard for piece")
	}
}

func (b *Board) setForPiece(p Piece, bb bitboard) {
	switch p {
	case WhiteKing:
		b.bbWhiteKing = bb
	case WhiteQueen:
		b.bbWhiteQueen = bb
	case WhiteRook:
		b.bbWhiteRook = bb
	case WhiteKnight:
		b.bbWhiteKnight = bb
	case WhiteBishop:
		b.bbWhiteBishop = bb
	case WhitePawn:
		b.bbWhitePawn = bb
	case BlackKing:
		b.bbBlackKing = bb
	case BlackQueen:
		b.bbBlackQueen = bb
	case BlackRook:
		b.bbBlackRook = bb
	case BlackKnight:
		b.bbBlackKnight = bb
	case BlackBishop:
		b.bbBlackBishop = bb
	case BlackPawn:
		b.bbBlackPawn = bb
	default:
		panic("No bitboard for piece")
	}
}
