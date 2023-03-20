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
