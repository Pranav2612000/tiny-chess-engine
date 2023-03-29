package main

type Move struct {
    from *Square
    to *Square
}

type Moves map[Square] []Square;

func (m Moves) GetNumberOfMoves() uint {
    var numberOfMoves uint = 0
    for _, moves := range m {
        numberOfMoves += uint(len(moves))
    }
    return numberOfMoves
}
