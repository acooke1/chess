package chess

type Color int8

const (
	NoColor Color = iota
	White
	Black
)

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	}
	return "-"
}

type PieceType int8

const (
	NoPieceType PieceType = iota
	King
	Queen
	Rook
	Knight
	Bishop
	Pawn
)

type Piece int8

const (
	NoPiece Piece = iota
	WhiteKing
	WhiteQueen
	WhiteRook
	WhiteKnight
	WhiteBishop
	WhitePawn
	BlackKing
	BlackQueen
	BlackRook
	BlackKnight
	BlackBishop
	BlackPawn
)

func (p Piece) Type() PieceType {
	switch p {
	case WhiteKing, BlackKing:
		return King
	case WhiteQueen, BlackQueen:
		return Queen
	case WhiteRook, BlackRook:
		return Rook
	case WhiteKnight, BlackKnight:
		return Knight
	case WhitePawn, BlackPawn:
		return Pawn
	}
	return NoPieceType
}

func (p Piece) Color() Color {
	switch p {
	case WhiteKing, WhiteQueen, WhiteRook, WhiteKnight, WhiteBishop, WhitePawn:
		return White
	case BlackKing, BlackQueen, BlackRook, BlackKnight, BlackBishop, BlackPawn:
		return Black
	}
	return NoColor
}

var allPieces = []Piece{
	WhiteKing, WhiteQueen, WhiteRook, WhiteKnight, WhiteBishop, WhitePawn,
	BlackKing, BlackQueen, BlackRook, BlackKnight, BlackBishop, BlackPawn,
}
