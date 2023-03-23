package main

type Board [120]Square;

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
        if int(index / 12) > 0 &&
            int(index / 12) < 9 &&
            int(index % 12) > 1 &&
            int(index % 12) < 10 {
            board[index].isPlayable = true;
        }
    }
    return board;
}
