package main

// Piece.variant is stored as
// P - Pawn 100
// N - Knight 280
// B - Bishop 320
// R - Rook 479
// Q - Queen 929
// K - King 999999
// Piece.color is stored as 'W' for White and 'B' for Black
// Point values borrowed from https://zserge.com/posts/carnatus/
type Piece struct {
    color byte 
    variant byte 
}

var pieceValueMap = map[byte]int{'P': 100, 'N': 280, 'B': 320, 'R': 479, 'Q': 929, 'K': 99999 }

func (p Piece) Value() int {
    return pieceValueMap[p.variant];
}
func (p Piece) Ours() bool {
    return true;
}
func (p Piece) Flip() Piece {
}

type Board [120]Piece
func (b Board) Flip() Board {
}
