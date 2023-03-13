package main

import "fmt"

// Piece.variant is stored as
// P - Pawn
// N - Knight
// B - Bishop
// R - Rook
// Q - Queen
// K - King
// Piece.color is stored as 'W' for White and 'B' for Black
type Piece struct {
    color byte 
    variant byte 
}

func (p Piece) Value() int {
    return 0;
}
func (p Piece) Ours() bool {
    return true;
}
func (p Piece) Flip() Piece {
}

type Board [120]Piece
func (b Board) Flip() Board {
}

type Square int
func (s Square) Flip() Square {
}



