package main

import (
    "testing"
)

func TestValue(t *testing.T) {
    pawn := Piece{color: 'W', variant: 'P'}
    if pawn.Value() != 100 {
        t.Error(`Piece of type Pawn should have value`, pawn.Value());
    }

    knight := Piece{color: 'B', variant: 'N'}
    if knight.Value() != 280 {
        t.Error(`Piece of type Knight should have value`, pawn.Value());
    }

    bishop:= Piece{color: 'W', variant: 'B'}
    if bishop.Value() != 320 {
        t.Error(`Piece of type Bishop should have value`, pawn.Value());
    }

    rook := Piece{color: 'B', variant: 'R'}
    if rook.Value() != 479 {
        t.Error(`Piece of type Rook should have value`, pawn.Value());
    }

    queen := Piece{color: 'B', variant: 'Q'}
    if queen.Value() != 929 {
        t.Error(`Piece of type Queen should have value`, pawn.Value());
    }

    king := Piece{color: 'W', variant: 'K'}
    if king.Value() != 99999 {
        t.Error(`Piece of type White should have value`, pawn.Value());
    }
}

func TestCopy(t *testing.T) {
  piece := Piece{color: 'B', variant: 'K'}
  copiedPiece := piece.Copy();
  if &piece == &copiedPiece {
    t.Error(`No new piece created while copying`);
  }

  if piece.variant != copiedPiece.variant || piece.color != copiedPiece.color {
    t.Error(`Data not copied correctly during piece.Copy`);
  }
}

func TestOurs(t *testing.T) {
    ourPiece := Piece{color: 'W', variant: 'P'}
    if ourPiece.Ours() != true {
        t.Error(`Piece with capital variant should be our piece`);
    }

    theirPiece := Piece{color: 'B', variant: 'k'}
    if theirPiece.Ours() != false {
        t.Error(`Piece with lowercase variant should not be our piece`);
    }
}

func TestFlip(t *testing.T) {
    ourPiece := Piece{color: 'W', variant: 'P'}
    ourPiece.Flip();
    if rune(ourPiece.variant) != 'p' {
        t.Error(`Failed to flip piece`);
    }
}

func TestGetMovesForQueen(t *testing.T) {
    // Create a empty test Board with Qc2
    board := GenerateEmptyBoard();

    queen := Piece{color: 'W', variant: 'Q'};
    board[33].piece = &queen;

    moves := queen.GetMoves(33, &board);

    movesIndex := []int{ 43, 53, 63, 73, 83, 93, 23, 34, 35, 36, 37, 38, 32, 31,
                    44, 55, 66, 77, 88, 24, 42, 51, 22 };
    for index, sq := range moves {
        if (sq.position != movesIndex[index]) {
            t.Errorf(`Position mismatch for move %v`, sq);
            break;
        }

        if (sq.isPlayable != true) {
            t.Errorf(`isPlayable mismatch for move %v`, sq);
            break;
        }

        if (sq.piece != nil && sq.piece.variant != 'Q') {
            t.Errorf(`Piece mismatch for move %v`, sq);
            break;
        }
    }
}

func TestPSTValue(t *testing.T) {
    king := Piece{color: 'W', variant: 'K'}
    value := king.PSTValue(45);
    if value != 59933 {
        t.Errorf(`Incorrect PST Value Expected: %d, Actual: %d`, 59933, value);
    }
}

func TestGetMovesForPawn(t *testing.T) {
    // Create a empty test Board with
    // W - e2, d4, a4
    // B - f3, a5
    board := GenerateEmptyBoard();

    pawnE2 := Piece{color: 'W', variant: 'P'};
    pawnD4 := Piece{color: 'W', variant: 'P'};
    pawnA4 := Piece{color: 'W', variant: 'P'};
    board[35].piece = &pawnE2;
    board[54].piece = &pawnD4;
    board[51].piece = &pawnA4;
    board[46].piece = &Piece{color: 'B', variant: 'p'};
    board[61].piece = &Piece{color: 'B', variant: 'p'};

    movesE2 := pawnE2.GetMoves(35, &board);
    movesD4 := pawnD4.GetMoves(54, &board);
    movesA4 := pawnA4.GetMoves(61, &board);

    movesE2Index := []int{ 45, 55, 46 };
    for index, sq := range movesE2 {
        if (sq.position != movesE2Index[index]) {
            t.Errorf(`Position mismatch for move %v`, sq);
            break;
        }

        if (sq.isPlayable != true) {
            t.Errorf(`isPlayable mismatch for move %v`, sq);
            break;
        }

        if (sq.piece != nil && sq.piece.variant != 'P') {
            t.Errorf(`Piece mismatch for move %v`, sq);
            break;
        }
    }

    movesD4Index := []int{ 64 };
    for index, sq := range movesD4 {
        if (sq.position != movesD4Index[index]) {
            t.Errorf(`Position mismatch for move %v`, sq);
            break;
        }

        if (sq.isPlayable != true) {
            t.Errorf(`isPlayable mismatch for move %v`, sq);
            break;
        }

        if (sq.piece != nil && sq.piece.variant != 'P') {
            t.Errorf(`Piece mismatch for move %v`, sq);
            break;
        }
    }

    if len(movesA4) == 0 {
        t.Errorf(`Invalid moves suggested for %v`, pawnA4);
    }
}
