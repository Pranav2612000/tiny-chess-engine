package main

import (
    "testing"
)

func TestSquareFlip(t *testing.T) {
    piece := Piece{color: 'W', variant: 'P'};
    square := Square{position: 27, piece: &piece, isPlayable: true};
    square.Flip();
    if square.position != 92 {
        t.Error(`Failed to Flip Square. New position incorrect`);
    }
    if square.piece.variant != 'p' {
        t.Error(`Failed to Flip Square. Piece variant not updated`);
    }
}

func TestGenerateSquareFromNotation(t *testing.T) {
    h2, _ := GenerateSquareFromNotation("h2", 'W');
    if h2.position != 38 || h2.piece.variant != 'P' || h2.piece.color != 'W' {
        t.Errorf(`Failed to create square for %s`, "h2");
        return;
    }

    ed4, _ := GenerateSquareFromNotation("ed4", 'W');
    if ed4.position != 54 || ed4.piece.variant != 'P' || ed4.piece.color != 'W' {
        t.Errorf(`Failed to create square for %s`, "ed4");
        return;
    }

    Nf3, _ := GenerateSquareFromNotation("Nf3", 'B');
    if Nf3.position != 73 || Nf3.piece.variant != 'n' || Nf3.piece.color != 'B' {
        t.Errorf(`Failed to create square for %s`, "Nf3");
        return;
    }
}

func TestSquareCopy(t *testing.T) {
    piece := Piece{color: 'W', variant: 'P'};
    square := Square{position: 27, piece: &piece, isPlayable: true};
    copiedSquare := square.Copy();

    if copiedSquare == square {
      t.Error(`No new square created while copying`);
    }

    if copiedSquare.piece == square.piece {
      t.Error(`No new piece created while copying square`);
    }

    if square.position != copiedSquare.position ||
        square.isPlayable != copiedSquare.isPlayable {
      t.Error(`Data not copied correctly during piece.Copy`);
    }
}
