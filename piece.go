package main

// Piece.variant is stored as
// P - Pawn 100
// N - Knight 280
// B - Bishop 320
// R - Rook 479
// Q - Queen 929
// K - King 999999
// . - Empty Square 0
// - - Empty Square 0
// Piece.color is stored as 'W' for White and 'B' for Black
// Opponent pieces are stored with the same symbols, but in lowercase
// Point values borrowed from https://zserge.com/posts/carnatus/
type PieceType byte;
type Piece struct {
    color byte 
    variant PieceType
};

const N, E, S, W = -10, 1, 10, 1;

var pieceValueMap = map[PieceType]int{'.': 0, '-': 0, 'P': 100, 'N': 280, 'B': 320, 'R': 479, 'Q': 929, 'K': 99999 }
var movesMap = map[byte][]int{
    'P': {N, N + N, N + W, N + E},
    'N': {N + N + E, N + N + W, E + E + N, E + E + S, S + S + E, S + S + W, W + W + S, W + W + N},
    'B': {N + E, S + E, N + W, S + W},
    'R': {N, S, E, W},
    'Q': {N, S, E, W, N + E, S + E, N + W, S + W},
    'K': {N, S, E, W, N + E, S + E, N + W, S + W},
}

func (p Piece) Value() int {
    return pieceValueMap[p.variant];
}
func (p Piece) Ours() bool {
    if IsUpper(string(p.variant)) {
        return true;
    }
    return false;
}

func (p *Piece) Flip() {
    p.variant = PieceType(flipByteCase(byte(p.variant)))
}
