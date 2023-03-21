package main

import (
    "testing"
)

func TestBoardFlip(t *testing.T) {
    // Create a test board with
    // - white h2, e4, Nf3
    // - black e7, Ra8
    var board Board;
    for index := range board {
        board[index].position = index;

        if int(index / 12) > 0 &&
            int(index / 12) < 9 &&
            int(index % 12) > 1 &&
            int(index % 12) < 10 {
            board[index].isPlayable = true;
        }
    }
    board[33].piece = &Piece{color: 'W', variant: 'P'}
    board[55].piece = &Piece{color: 'W', variant: 'P'}
    board[44].piece = &Piece{color: 'W', variant: 'N'}
    board[91].piece = &Piece{color: 'B', variant: 'P'}
    board[99].piece = &Piece{color: 'B', variant: 'R'}

    board.Flip();

    if board[86].piece.variant != 'p' ||
        board[86].position != 86 {
        t.Error(`Board Flip failed for h2`);
    }
    if board[64].piece.variant != 'p' ||
        board[64].position != 64 {
        t.Error(`Board Flip failed for e4`);
    }
    if board[75].piece.variant != 'n' ||
        board[75].position != 75 {
        t.Error(`Board Flip failed for Nf3`);
    }
    if board[28].piece.variant != 'p' ||
        board[28].position != 28 {
        t.Error(`Board Flip failed for e7`);
    }
    if board[20].piece.variant != 'r' ||
        board[20].position != 20 {
        t.Error(`Board Flip failed for Ra8`);
    }
}
