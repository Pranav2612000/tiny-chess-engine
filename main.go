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
  fmt.Printf("Initial Score %v\n", position.score);

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

    // If the current player has no possible moves we end the game
    moves := position.Moves();
    if moves.GetNumberOfMoves() == 0 {
        fmt.Printf("Game over!! %s wins\n", oppositeColor);
        break;
    }

    if color == computerColor {
      cposition := position.Copy();
      fmt.Printf("%v\n", position);
      fmt.Printf("%v\n", cposition);
      fmt.Printf("%v\n", cposition.board);
      DrawBoard(cposition.board);
      moveRaw := s.SearchMove(cposition, 10000);
      fmt.Printf("%v\n", cposition.board);
      DrawBoard(cposition.board);
      fmt.Printf("\n%v\n", cposition);
      fmt.Printf("%v\n", position);
      position.Move(moveRaw);
      DrawBoard(position.board);
    } else {
      //fmt.Printf("The computer played %v\n", move);
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
    fmt.Printf("Position Score: %v\n", position.score);

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
