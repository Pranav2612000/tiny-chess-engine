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

    movesIndex := []int{ 23, 43, 53, 63, 73, 83, 93, 32, 31, 34, 35, 36, 37, 38,
                    22, 42, 51, 24, 44, 55, 66, 77, 88 };
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

func TestGetMovesForPawn(t *testing.T) {
    // Create a empty test Board with e2
    board := GenerateEmptyBoard();

    pawn := Piece{color: 'W', variant: 'P'};
    board[35].piece = &pawn;

    moves := pawn.GetMoves(35, &board);

    movesIndex := []int{ 45, 55, 44, 46 };
    for index, sq := range moves {
        if (sq.position != movesIndex[index]) {
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
}
