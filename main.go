package main

import (
    "fmt"
    "flag"
)

var GlobalIsDebugMode bool;
func startTwoPlayerGame () {
  fmt.Println("Starting a new 2 player game")
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

      if GlobalIsDebugMode {
        fmt.Printf("Score %v\n", position.score);
      }

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

func startGameWithComputer() {
  computerColor := "white";
  s := &Searcher{ tp: map[PositionRaw]entry{} }

  fmt.Println("Starting a new game with Computer")
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

  var move string;
  for true {
    var color string;
    var oppositeColor string;
    if position.turn {
      color = "white";
      oppositeColor = "black";
    } else {
      color = "black";
      oppositeColor = "white";
    }

    if GlobalIsDebugMode {
      fmt.Printf("Score %v\n", position.score);
    }

    // If the current player has no possible moves we end the game
    moves := position.Moves();
    if moves.GetNumberOfMoves() == 0 {
        fmt.Printf("Game over!! %s wins\n", oppositeColor);
        break;
    }

    if color == computerColor {
      cposition := position.Copy();
      moveRaw := s.SearchMove(cposition, 10000);
      position.Move(moveRaw, false);
      DrawBoard(position.board);
    } else {
      fmt.Printf("Your move: ");
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
    }

    position.turn = !position.turn;

  }
  fmt.Println("Thank you for playing");
}

func main() {
  isTwoPlayerGame := flag.Bool("two-player", false, "Starts a two player game");
  isDebugMode := flag.Bool("debug", false, "Shows debug logs");
  flag.Parse();
  GlobalIsDebugMode = *isDebugMode;

  if (*isTwoPlayerGame) {
    startTwoPlayerGame();
  } else {
    startGameWithComputer();
  }
}
