package main

type Position struct {
    board *Board
    score int
    wc [2]bool
    bc [2]bool
    ep *Square
    kp *Square
}

func (p *Position) Flip() {
    p.board.Flip();
    p.score = -1 * p.score;
    p.ep.Flip();
    p.kp.Flip();

    var temp [2]bool;
    temp[0] = p.wc[0];
    temp[1] = p.wc[1];
    p.wc = [2]bool{p.bc[0], p.bc[1]}
    p.bc = [2]bool{temp[0], temp[1]}
}

/* Function to return a list of our possible moves */
func (pos *Position) Moves() (moves Moves) {
    moves = make(map[Square][]Square);
    for _, sq := range pos.board {
        // Ignore the non-playable squares
        if !sq.isPlayable {
            continue;
        }

        // Return early if there is no piece on this square
        piece := sq.piece;
        if piece == nil {
            continue;
        }

        // Return early if piece does not belong to us
        if !piece.Ours() {
            continue;
        }

        currentPieceMoves := piece.GetMoves(sq.position, pos.board);
        moves[sq] = currentPieceMoves;
    }
    return moves;
}

func (pos *Position) Move(move Move) {
    from := *move.from;
    to := *move.to;

    pos.board[to.position].piece = pos.board[from.position].piece;
    pos.board[from.position].piece = nil;

    return;
}

func CreateStartPosition() Position {
    board := GenerateInitialPositionBoard();
    DrawBoard(&board);

    return Position{
        board: &board,
        score: 0,
        wc: [2]bool{false, false},
        bc: [2]bool{false, false},
        ep: nil,
        kp: nil,
    }
}
