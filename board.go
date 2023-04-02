package main

type Board [120]Square;

func (b *Board) Copy() Board {
  var board Board;

  for index := range b {
    b[index] = b[index].Copy();
  }

  return board;
}

func (b *Board) Flip() {
    var flippedBoard Board;
    for index := range b {
        b[index].Flip()
        flippedBoard[119 - index] = b[index];
    }
    for index := range flippedBoard {
        b[index] = flippedBoard[index];
    }
}

func GenerateEmptyBoard() Board {
    var board Board;
    for index := range board {

        // Set the index for each square
        board[index].position = index;

        // For invalid, border squares, set
        // isPlayable to false and pieceVariant
        // to <space>
        if int(index / 10) > 1 &&
            int(index / 10) < 10 &&
            int(index % 10) > 0 &&
            int(index % 10) < 9 {
            board[index].isPlayable = true;
        }
    }
    return board;
}

func GenerateInitialPositionBoard() Board {
    board := GenerateEmptyBoard();

    // White pieces
    board[21].piece =  &Piece{color: 'W', variant: 'R'}
    board[22].piece =  &Piece{color: 'W', variant: 'N'}
    board[23].piece =  &Piece{color: 'W', variant: 'B'}
    board[24].piece =  &Piece{color: 'W', variant: 'Q'}
    board[25].piece =  &Piece{color: 'W', variant: 'K'}
    board[26].piece =  &Piece{color: 'W', variant: 'B'}
    board[27].piece =  &Piece{color: 'W', variant: 'N'}
    board[28].piece =  &Piece{color: 'W', variant: 'R'}

    // White pawns
    board[31].piece =  &Piece{color: 'W', variant: 'P'}
    board[32].piece =  &Piece{color: 'W', variant: 'P'}
    board[33].piece =  &Piece{color: 'W', variant: 'P'}
    board[34].piece =  &Piece{color: 'W', variant: 'P'}
    board[35].piece =  &Piece{color: 'W', variant: 'P'}
    board[36].piece =  &Piece{color: 'W', variant: 'P'}
    board[37].piece =  &Piece{color: 'W', variant: 'P'}
    board[38].piece =  &Piece{color: 'W', variant: 'P'}

    // Black pieces
    board[91].piece =  &Piece{color: 'B', variant: 'r'}
    board[92].piece =  &Piece{color: 'B', variant: 'n'}
    board[93].piece =  &Piece{color: 'B', variant: 'b'}
    board[94].piece =  &Piece{color: 'B', variant: 'q'}
    board[95].piece =  &Piece{color: 'B', variant: 'k'}
    board[96].piece =  &Piece{color: 'B', variant: 'b'}
    board[97].piece =  &Piece{color: 'B', variant: 'n'}
    board[98].piece =  &Piece{color: 'B', variant: 'r'}

    // Black pawns
    board[81].piece =  &Piece{color: 'B', variant: 'p'}
    board[82].piece =  &Piece{color: 'B', variant: 'p'}
    board[83].piece =  &Piece{color: 'B', variant: 'p'}
    board[84].piece =  &Piece{color: 'B', variant: 'p'}
    board[85].piece =  &Piece{color: 'B', variant: 'p'}
    board[86].piece =  &Piece{color: 'B', variant: 'p'}
    board[87].piece =  &Piece{color: 'B', variant: 'p'}
    board[88].piece =  &Piece{color: 'B', variant: 'p'}

    return board;
}
