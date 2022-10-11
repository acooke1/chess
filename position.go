package chess

type CastlingRights string

type Position struct {
	board          *Board
	turn           Color
	castlingRights CastlingRights
	enPassantSq    Square
	halfmoveClock  int
	fullmoves      int
	inCheck        bool
	validMoves     []*Move
}
