package main

import (
    "strconv"
    "errors"
)

type Square struct {
    position int
    piece *Piece
    isPlayable bool
}

var columnLetterToNumberMap = map[byte]int{ 'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8 }

func (s *Square) Copy() Square {
  if s == nil {
    return Square{};
  }

  var piecePtr *Piece;
  piecePtr = nil;
  if s.piece != nil {
    piece := s.piece.Copy()
    piecePtr = &piece
  }
  return Square{ position: s.position, piece: piecePtr, isPlayable: s.isPlayable };
}

func (s *Square) Flip() {
    s.position = 119 - s.position;
    if s.piece != nil {
        s.piece.Flip()
    }
}

func (s *Square) PSTValue() int {
    if s.piece == nil {
        return 0;
    }
    return s.piece.PSTValue(s.position);
}

func GenerateSquareFromNotation(notation string, color byte) (Square, error) {
        var piece Piece;
        if !IsUpper(notation[0:1]) {
            piece = Piece{color: color, variant: 'P'}
        } else {
            piece = Piece{color: color, variant: PieceType(byte(notation[0]))}
        }

        row, err := strconv.Atoi(notation[len(notation) - 1:]);
        if err != nil {
            return Square{}, errors.New("Invalid row number");
        }

        columnChar := notation[len(notation) - 2];
        column := columnLetterToNumberMap[byte(columnChar)];

        if color == 'B' {
            piece.Flip();
            row = 9 - row;
            column = 9 - column;
        }

        position := (row + 1) * 10 + column;

        return Square{position: position, piece: &piece, isPlayable: true}, nil
}
