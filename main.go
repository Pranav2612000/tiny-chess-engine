package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting a new game")

    board := GenerateInitialPositionBoard();
    position := Position{
        board: &board,
        score: 0,
        wc: [2]bool{false, false},
        bc:[2]bool{false, false},
        ep: nil,
        kp: nil,
        turn: true,
    };
    DrawBoard(position.board);
    var move string;
    for true {
        fmt.Scanln(&move);

        if move == "exit" {
            break;
        }

        err := position.MoveFromNotation(move);

        if err != nil {
            fmt.Printf("Error: %v\n", err);
            fmt.Println("Please enter a valid move");
            continue;
        }

        position.Flip();
        DrawBoard(position.board);
    }
    fmt.Println("Thank you for playing");
}
