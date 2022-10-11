package chess

type Color int8

const (
	NoColor Color = iota
	White
	Black
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

var allPieces = []Piece{
	WhiteKing, WhiteQueen, WhiteRook, WhiteKnight, WhiteBishop, WhitePawn,
	BlackKing, BlackQueen, BlackRook, BlackKnight, BlackBishop, BlackPawn,
}
