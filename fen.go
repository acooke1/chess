package chess

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	startingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func FENError(message string) error {
	return fmt.Errorf("ERROR: invalid FEN string... " + message)
}

func DecodeFen(fen string) (*Position, error) {
	components := strings.Split(fen, " ")
	if len(components) != 6 {
		return nil, FENError("FEN string has more than 6 parts")
	}
	board, err := decodeBoard(components[0])
	if err != nil {
		return nil, err
	}
	turn, ok := fenTurnMap[components[1]]
	if !ok {
		return nil, FENError("turn string incorrectly formatted")
	}
	rights, err := decodeRights(components[2])
	if err != nil {
		return nil, err
	}
	square, err := decodeEnPassant(components[3])
	if err != nil {
		return nil, err
	}
	halfmoves, err := decodeHalfmoves(components[4])
	if err != nil {
		return nil, err
	}
	fullmoves, err := decodeFullmoves(components[5])
	if err != nil {
		return nil, err
	}

	position := &Position{
		board:          board,
		turn:           turn,
		castlingRights: rights,
		enPassantSq:    square,
		halfmoveClock:  halfmoves,
		fullmoves:      fullmoves,
	}
	return position, nil
}

func decodeBoard(boardString string) (*Board, error) {
	rankStrings := strings.Split(boardString, "/")
	if len(rankStrings) != 8 {
		return nil, FENError("board string does not have exactly 8 ranks")
	}
	m := map[Square]Piece{}
	for i, rankStr := range rankStrings {
		rank := Rank(7 - i)
		fileCntr := 0
		for _, squareStr := range rankStr {
			piece, ok := fenPieceMap[fmt.Sprintf("%c", squareStr)]
			if ok {
				file := File(fileCntr)
				m[newSquare(rank, file)] = piece
				fileCntr++
			} else {
				emptyCount, err := strconv.Atoi(fmt.Sprintf("%c", squareStr))
				if err != nil {
					return nil, FENError("board string conatins a character which is not either a piece symbol or an integer")
				}
				if fileCntr+emptyCount > 8 || emptyCount < 1 {
					return nil, FENError("board string contains an integer which is ill formatted")
				}
				fileCntr += emptyCount
			}
		}
	}
	return newBoard(m), nil
}

func decodeRights(rightsString string) (CastlingRights, error) {
	if rightsString == "-" {
		return CastlingRights(rightsString), nil
	}
	if len(rightsString) > 4 {
		return "", fmt.Errorf("rights string longer than 4 characters")
	}
	K, k, Q, q := "", "", "", ""
	for _, side := range rightsString {
		switch side {
		case 'K':
			K += "K"
		case 'k':
			k += "k"
		case 'Q':
			Q += "Q"
		case 'q':
			q += "q"
		default:
			return "", fmt.Errorf("rights string contains a character which is not one of 'KQkq'")
		}
	}
	if len(K) > 1 || len(k) > 1 || len(Q) > 1 || len(q) > 1 {
		return "", fmt.Errorf("rights string contains more than one of 'KQkq'")
	}
	return CastlingRights(K + Q + k + q), nil
}

func decodeEnPassant(squareString string) (Square, error) {
	if squareString == "-" {
		return NoSquare, nil
	}
	sq, ok := strToSquareMap[squareString]
	if !ok {
		return NoSquare, fmt.Errorf("invalid given en passant square")
	}
	return sq, nil
}

func decodeHalfmoves(halfmoveString string) (int, error) {
	halfmoves, err := strconv.Atoi(halfmoveString)
	if err != nil || halfmoves < 0 {
		return 0, fmt.Errorf("halfmoves is not an integer >= 0")
	}
	return halfmoves, nil
}

func decodeFullmoves(fullmoveString string) (int, error) {
	fullmoves, err := strconv.Atoi(fullmoveString)
	if err != nil || fullmoves < 1 {
		return 0, fmt.Errorf("fullmoves is not an integer >= 1")
	}
	return fullmoves, nil
}

var fenTurnMap = map[string]Color{
	"b": Black,
	"w": White,
}

var fenPieceMap = map[string]Piece{
	"K": WhiteKing,
	"Q": WhiteQueen,
	"R": WhiteRook,
	"N": WhiteKnight,
	"B": WhiteBishop,
	"P": WhitePawn,
	"k": BlackKing,
	"q": BlackQueen,
	"r": BlackRook,
	"n": BlackKnight,
	"b": BlackBishop,
	"p": BlackPawn,
}
