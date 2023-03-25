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

func TestPositionMoves(t *testing.T) {
    // Create a test board with
    // - white h2, Ke4, Nf3, Rh1,Bc5
    // - black b6, Ra8
    board := GenerateEmptyBoard();
    h2 := Piece{color: 'W', variant: 'P'};
    Ke4 := Piece{color: 'W', variant: 'K'}
    Nf3 := Piece{color: 'W', variant: 'N'}
    Rh1 := Piece{color: 'W', variant: 'R'}
    Bc5 := Piece{color: 'W', variant: 'B'}
    b6 :=  Piece{color: 'B', variant: 'p'}
    Ra8 := Piece{color: 'B', variant: 'r'}

    board[38].piece = &h2
    board[55].piece = &Ke4
    board[46].piece = &Nf3
    board[28].piece = &Rh1
    board[63].piece = &Bc5
    board[72].piece = &b6
    board[91].piece = &Ra8


    position := Position{
        board: &board,
        score: 0,
        wc: [2]bool{true, false},
        bc:[2]bool{false, true},
        ep: nil,
        kp: nil,
        turn: true,
    };
    allMoves := position.Moves();

    movesStringified := make(map[Piece] []int);
    movesStringified[h2] = []int{48, 58}
    movesStringified[Ke4] = []int{65, 45, 56, 54, 66, 64, 44}
    movesStringified[Nf3] = []int{67, 65, 58, 27, 25, 34, 54}
    movesStringified[Rh1] = []int{27, 26, 25, 24, 23, 22, 21}
    movesStringified[Bc5] = []int{74, 85, 96, 54, 45, 36, 27, 72, 52, 41}

    for start, moves := range allMoves {
        currentPiece := (*start.piece);
        currentPieceMoves := movesStringified[currentPiece];
        for index, sq := range moves {
            if (sq.position != currentPieceMoves[index]) {
                t.Errorf(`Piece: %s - Position mismatch for move %v`, string(currentPiece.variant), sq);
                break;
            }

            if (sq.isPlayable != true) {
                t.Errorf(`Piece: %s - isPlayable mismatch for move %v`, string(currentPiece.variant), sq);
                break;
            }

            if (sq.piece != nil && sq.piece.variant != currentPiece.variant) {
                t.Errorf(`Piece: %s - Piece mismatch for move %v`, string(currentPiece.variant), sq);
                break;
            }
        }
    }
}

func TestPositionMove(t *testing.T) {
    board := GenerateEmptyBoard();
    h2 := Piece{color: 'W', variant: 'P'};
    board[38].piece = &h2

    position := Position{
        board: &board,
        score: 0,
        wc: [2]bool{true, false},
        bc:[2]bool{false, true},
        ep: nil,
        kp: nil,
    };

    from := Square{position: 38, piece: &h2, isPlayable: true};
    to := Square{position: 48, piece: &h2, isPlayable: true};

    position.Move(Move{from: &from, to: &to});

    if position.board[48].piece == nil {
        t.Errorf(`Position update failed. Square is null`);
        return;
    }
    if position.board[48].piece.variant != 'P' || position.board[48].piece.color != 'W' {
        t.Errorf(`Position update failed. Incorrect piece at position`);
        return;
    }
}

func TestMoveFromNotation(t *testing.T) {
    board := GenerateEmptyBoard();
    a7 := Piece{color: 'B', variant: 'p'};
    board[81].piece = &a7

    position := Position{
        board: &board,
        score: 0,
        wc: [2]bool{true, false},
        bc:[2]bool{false, true},
        ep: nil,
        kp: nil,
        turn: false,
    };

    err := position.MoveFromNotation("a6");
    if err != nil {
        t.Errorf(`Position update failed. There was an error: %v`, err);
        return;
    }
    if position.board[71].piece == nil {
        t.Errorf(`Position update failed. Square is null`);
        return;
    }
    if position.board[71].piece.variant != 'p' || position.board[71].piece.color != 'B' {
        t.Errorf(`Position update failed. Incorrect piece at position`);
        return;
    }
}

func TestGetValueOfMove(t *testing.T) {
    board := GenerateEmptyBoard();
    h2 := Piece{color: 'W', variant: 'P'};
    g3 := Piece{color: 'B', variant: 'p'};
    board[38].piece = &h2;
    board[47].piece = &g3;

    position := Position{
        board: &board,
        score: 0,
        wc: [2]bool{true, false},
        bc:[2]bool{false, true},
        ep: nil,
        kp: nil,
    };

    from := Square{position: 38, piece: &h2, isPlayable: true};
    to := Square{position: 48, piece: &h2, isPlayable: true};

    score := position.GetValueOfMove(Move{from: &from, to: &to});
    if score != -83 {
        t.Errorf(`Incorrect score returned. Expected %d Actual %d`, -83, score);
    }

    toWithCapture := Square{position: 47, piece: &h2, isPlayable: true};
    scoreWithCapture := position.GetValueOfMove(Move{from: &from, to: &toWithCapture});
    if scoreWithCapture != 63 {
        t.Errorf(`Incorrect score returned. Expected %d Actual %d`, -83, scoreWithCapture);
    }
}
