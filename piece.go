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
// <space> - Boundary Square 0
// Piece.color is stored as 'W' for White and 'B' for Black
// Opponent pieces are stored with the same symbols, but in lowercase
// Point values borrowed from https://zserge.com/posts/carnatus/
type PieceType byte;
type Piece struct {
    color byte 
    variant PieceType
};

const N, E, S, W = 10, 1, -10, -1;

var pieceValueMap = map[PieceType]int{'.': 0, '-': 0, 'P': 100, 'N': 280, 'B': 320, 'R': 479, 'Q': 929, 'K': 99999 }
var movesMap = map[PieceType][]int{
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

func (p *Piece) GetMoves(currentPosition int, b *Board) []Square {
    var moves []Square;

    // Iterate through movement in all possible direction
    for _, d := range movesMap[p.variant] {

        // Loop to search for all possible moves in an direction
        for j := currentPosition + d; ; j = j + d {
            newPositionSquare := b[j];

            // We stop searching for moves in this direction if we
            // reach the end of the board OR our movement is blocked
            // by another of our piece
            if !newPositionSquare.isPlayable ||
                ( newPositionSquare.piece != nil &&
                    newPositionSquare.piece.variant == ' ' ) ||
                ( newPositionSquare.piece != nil &&
                    newPositionSquare.piece.Ours() ) ||
                ( newPositionSquare.piece != nil &&
                    newPositionSquare.piece.variant != '.' &&
                    newPositionSquare.piece.Ours() ) {
                break;
            }

            moves = append(moves, Square{position: j, piece: p, isPlayable: true});

            variant := ToUpper(byte(p.variant));

            // If piece is one of pawn, night or king, it can move only once
            // in an direction, so we break for these pieces
            if variant == 'P' || variant == 'N' || variant == 'K' {
                break;
            }

            // If this is a capture move, we cannot proceed further in this
            // direction, so we stop the movement in this direction
            if ( newPositionSquare.piece != nil &&
                !newPositionSquare.piece.Ours() ) ||
                ( newPositionSquare.piece != nil &&
                    newPositionSquare.piece.variant != ' ' &&
                    newPositionSquare.piece.variant != '.' &&
                    !newPositionSquare.piece.Ours() ) {
                break;
            }
        }
    }
    return moves;
}
