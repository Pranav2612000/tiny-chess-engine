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
        var color string;
        var oppositeColor string;
        if position.turn {
            color = "white";
            oppositeColor = "black"
        } else {
            color = "black";
            oppositeColor = "white"
        }

        // If the current player has no possible moves we end the game
        moves := position.Moves();
        if moves.GetNumberOfMoves() == 0 {
            fmt.Printf("Game over!! %s wins\n", oppositeColor);
            break;
        }

        fmt.Printf("%s's move: ", color);
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
        position.turn = !position.turn;

        DrawBoard(position.board);
    }
    fmt.Println("Thank you for playing");
}
