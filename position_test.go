package main

import (
    "testing"
)

func TestPositionFlip(t *testing.T) {
    eppiece := Piece{color: 'W', variant: 'P'};
    epsquare := Square{position: 27, piece: &eppiece, isPlayable: true};

    kppiece := Piece{color: 'W', variant: 'K'};
    kpsquare := Square{position: 34, piece: &kppiece, isPlayable: true};

    board := GenerateEmptyBoard();
    board[33].piece = &Piece{color: 'W', variant: 'P'}
    board[55].piece = &Piece{color: 'W', variant: 'P'}
    board[44].piece = &Piece{color: 'W', variant: 'N'}
    board[91].piece = &Piece{color: 'B', variant: 'P'}
    board[99].piece = &Piece{color: 'B', variant: 'R'}

    position := Position{board: &board, score: 5, wc: [2]bool{true, false}, bc:[2]bool{false, true}, ep: &epsquare, kp: &kpsquare}
    position.Flip();

    if position.score != -5 {
        t.Error(`Score not flipped when flipping position`);
    }

    if position.wc[0] != false || position.wc[1] != true {
        t.Error(`wc not flipped when flipping position`);
    }

    if position.bc[0] != true || position.bc[1] != false {
        t.Error(`bc not flipped when flipping position`);
    }
}
