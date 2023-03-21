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
