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
