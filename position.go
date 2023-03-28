package main

import (
    "errors"
)

type Position struct {
    board *Board
    score int
    wc [2]bool
    bc [2]bool
    ep *Square
    kp *Square
    turn bool // true for white, false for black
}

func (p *Position) Flip() {
    p.board.Flip();
    p.score = -1 * p.score;

    if p.ep != nil {
        p.ep.Flip();
    }

    if p.kp != nil {
        p.kp.Flip();
    }

    var temp [2]bool;
    temp[0] = p.wc[0];
    temp[1] = p.wc[1];
    p.wc = [2]bool{p.bc[0], p.bc[1]}
    p.bc = [2]bool{temp[0], temp[1]}
}

/* Function to return a list of our possible moves */
func (pos *Position) Moves() (moves Moves) {
    // If its black's turn we want the moves from his perspective,
    // so we flip the board for our internal calculations.
    if (!pos.turn) {
        pos.Flip();
    }

    moves = pos.RawMoves()

    // if our king is under threat, our move should make our king safe
    movesThreateningKing := pos.MovesThreateningKing();
    if len(movesThreateningKing) != 0 {
        // to check for valid moves when under check, we perform the move,
        // then check if the king has been made safe, and undo the move
        for start, thisPieceMoves := range moves {
            var validMoves []Square;
            for _, sq := range thisPieceMoves {
                var temp Square;
                // if the current move is a capture, we need to store the captured piece details
                // so that we can restore it back when we undo the move
                if (pos.board[sq.position].piece != nil) {
                    newPiece := pos.board[sq.position].piece.Copy()
                    temp = Square{position:sq.position, piece: &newPiece, isPlayable: true}
                }
                // perform the move
                pos.Move(Move{from: &start, to: &sq});

                // check if the king is safe
                movesThreateningKingAfterThisMove := pos.MovesThreateningKing()
                if len(movesThreateningKingAfterThisMove) == 0 {
                    validMoves = append(validMoves, sq);
                }

                // undo the move
                pos.Move(Move{from: &sq, to: &start});

                pos.board[temp.position] = temp

            }
            moves[start] = validMoves
        }
    }
    // If we had previously flipped the board, we flip it back
    // to the original position
    if (!pos.turn) {
        pos.Flip();
    }
    return moves;
}

func (pos *Position) RawMoves() (moves Moves) {
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

    return moves
}

func (pos *Position) Move(move Move) {
    from := *move.from;
    to := *move.to;

    pos.board[to.position].piece = pos.board[from.position].piece;
    pos.board[from.position].piece = nil;

    return;
}

func (pos *Position) MoveFromNotation(notation string) (error) {
    allMoves := pos.Moves();
    var color byte;
    if pos.turn {
        color = 'W';
    } else {
        color = 'B';
    }
    toSquare, err := GenerateSquareFromNotation(notation, color);

    if err != nil {
        return errors.New("Invalid notation");
    }

    for start, moves := range allMoves {
        currentPiece := (*start.piece);

        if toSquare.piece.variant != currentPiece.variant {
            continue;
        }

        movePossibleFromThisPiece := false;
        for _, sq := range moves {
            if sq.position == toSquare.position {
                movePossibleFromThisPiece = true;
                break;
            }
        }

        if movePossibleFromThisPiece {
            // Since the move positions are calculated from White's POV,
            // if its black's move we first flip the board before we apply
            // the move
            if (!pos.turn) {
                pos.Flip();
            }

            pos.Move(Move{from: &start, to: &toSquare});

            // Once the move has been successful, we flip the board back to
            // the original
            if (!pos.turn) {
                pos.Flip();
            }
            return nil;
        }
    }
    return errors.New("No piece satisfies the condition");
}

func (pos *Position) GetValueOfMove(move Move) int {
    from := *move.from;
    to := *move.to;
    fromPiece := from.piece;
    toPiece := to.piece;

    score := toPiece.PSTValue(to.position) - fromPiece.PSTValue(from.position);

    // if the move is a capture
    if pos.board[to.position].piece != nil {
        // add the PST value of the captured piece, from the opponents POV
        capturedSquare := Square{
            piece: &Piece{color: pos.board[to.position].piece.color, variant: pos.board[to.position].piece.variant},
            position: to.position,
            isPlayable: true,
        }

        capturedSquare.Flip();

        score += capturedSquare.piece.PSTValue(capturedSquare.position);
    }

    return score;
}

func (pos *Position) GetKingPosition() (Square, error)  {
    for _, sq := range pos.board {
        if sq.piece == nil {
            continue;
        }

        if sq.piece.variant == 'K' {
            return sq, nil
        }
    }
    return Square{}, errors.New("No king on the board")
}

func (pos *Position) GetOpponentKingPosition() (Square, error)  {
    for _, sq := range pos.board {
        if sq.piece == nil {
            continue;
        }

        if sq.piece.variant == 'k' {
            return sq, nil
        }
    }
    return Square{}, errors.New("No king on the board")
}

func (pos *Position) MovesThreateningKing() []Move {
    // Flip the board to get the moves from opponents perspective
    pos.Flip();

    // Get the position of the opponent's opponent's (i.e current player's) king
    king, _ := pos.GetOpponentKingPosition()

    var movesThreateningKing []Move;

    // Get all moves for opponent
    allMoves := pos.RawMoves();

    // and iterate through them to search for moves which are a check
    for start, moves := range allMoves {
        for _, sq := range moves {
            if sq.position == king.position {
                movesThreateningKing = append(movesThreateningKing, Move{
                    from: &Square{position: 119 - start.position, piece: start.piece, isPlayable: true},
                    to: &Square{position: 119 - sq.position, piece: sq.piece, isPlayable: true},
                })
                break;
            }
        }
    }

    // Unflip the flipped board
    pos.Flip();

    return movesThreateningKing
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
