package main

import (
    "testing"
)

func TestBoardFlip(t *testing.T) {
    // Create a test board with
    // - white h2, e4, Nf3
    // - black e7, Ra8
    board := GenerateEmptyBoard();
    board[38].piece = &Piece{color: 'W', variant: 'P'}
    board[55].piece = &Piece{color: 'W', variant: 'P'}
    board[46].piece = &Piece{color: 'W', variant: 'N'}
    board[85].piece = &Piece{color: 'B', variant: 'p'}
    board[91].piece = &Piece{color: 'B', variant: 'r'}

    board.Flip();

    if board[81].piece.variant != 'p' ||
        board[81].position != 81 {
        t.Error(`Board Flip failed for h2`);
    }
    if board[64].piece.variant != 'p' ||
        board[64].position != 64 {
        t.Error(`Board Flip failed for e4`);
    }
    if board[73].piece.variant != 'n' ||
        board[73].position != 73 {
        t.Error(`Board Flip failed for Nf3`);
    }
    if board[34].piece.variant != 'P' ||
        board[34].position != 34 {
        t.Error(`Board Flip failed for e7`);
    }
    if board[28].piece.variant != 'R' ||
        board[28].position != 28 {
        t.Error(`Board Flip failed for Ra8`);
    }
}
