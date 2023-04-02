package main

/* THIS IMPLEMENTATION OF ALPHA-BETA PRUNING MIN-MAX IS RIFE WITH BUGS
   PLEASE USE YOUR OWN IMPLEMENTATION.
 */
import (
  "fmt"
)

var (
  MateValue = pieceValueMap['K'] + 10 * pieceValueMap['Q'] // A very large value to denote mate score
  MaxTableSize = 10000000 // length of the transposition table
)

type entry struct {
  depth int // depth of the search
  score int // score at the current move

  // Values used for alpha beta pruning -
  alpha float64 // alpha value at the current node
  beta float64 // beta value at the current node

  // Value used for windowing
  gamma int

  // best move at this node
  move Move
}

type Searcher struct {
  tp map[Position]entry // transposition table to keep track of best move at a position
  nodes int // Number of nodes to consider, to keep a check on the complexity of the calculations
}

func (s *Searcher) SearchMove(pos Position, maxNodes int) (m Move) {
  fmt.Printf("Starting score %v, turn: %v\n", pos.score, pos.turn);
  s.nodes = 0;

  fmt.Printf("Mate value %v\n", MateValue);
  for depth := 1; depth < 2; depth++ {
    alpha, beta := float64(-3 * MateValue), float64(3 * MateValue);
    //gamma := int(( alpha + beta + 1 ) / 2);
    s.SearchNew(pos, 6, alpha, beta);
    //fmt.Printf("%v", score);
  }

  return s.tp[pos].move;
}

func (s *Searcher) SearchNew(pos Position, depth int, alpha, beta float64) (searchScore float64) {
  score := float64(pos.score);
  if (pos.score < 0) {
    score = -1 * score;
  }

  if score >= float64(MateValue) {
    return score;
  }

  e, ok := s.tp[pos];
  if ok && e.depth >= depth  {
    return float64(e.score);
  }

	if depth == 0 {
		return float64(pos.score);
	}

  allMoves := pos.Moves()

  var bestMove Move
  var bestScore float64;
  for start, moves := range allMoves {
    for _, sq := range moves {
      startCp := start.Copy();
      sqCp := sq.Copy();
      move := Move{ from: &startCp, to: &sqCp} 

      flippedPos := pos.Copy();
      flippedPos.Move(move);
      flippedPos.Flip();

      tempScore := -1 * s.SearchNew(flippedPos, depth - 1, -beta, -alpha);

      if tempScore > alpha {
        alpha = tempScore;
        bestMove = move;

        if beta <= alpha {
          bestScore = beta;
          break;
        }
      }
    }

    if bestScore == beta {
      break;
    }
  }

  s.tp[pos] = entry{
    depth: depth,
    score: int(bestScore),
    alpha: alpha,
    beta: beta,
    gamma: 0,
    move: bestMove,
  };
	return alpha;
}

func (s *Searcher) Search(pos Position, alpha float64, beta float64, gamma int, depth int) (score int) {
  s.nodes++;

  e, ok := s.tp[pos];
  if ok && e.depth >= depth  {
    return e.score;
  }

  if Abs(pos.score) > MateValue {
    return pos.score;
  }

  nullScore := pos.score;

  /*
  if depth < 0 {
    // If we are in the maximizing part
  }
  */
  bestScore, bestMove := -3*MateValue, Move{};

  allMoves := pos.Moves()
  //fmt.Printf("All moves: %v\n", allMoves);
  for start, moves := range allMoves {
    for _, sq := range moves {
      
      if depth <= 0 {
        break;
      }

      startCp := start.Copy();
      sqCp := sq.Copy();
      move := Move{ from: &startCp, to: &sqCp} 

      flippedPos := pos.Copy();
      flippedPos.Move(move);
      flippedPos.Flip();

      score := -1 * s.Search(flippedPos, beta, alpha, gamma, depth - 1);

      // -ve score denotes it's an iteration for them
      if (score < 0 ) {
        //fmt.Printf("-ve score observed\n");
        if ( float64(score) > alpha ) {
          alpha = float64(score);
          break;
        }
      } else {
        //fmt.Printf("+ve score observed\n");
        if ( float64(score) > beta ) {
          beta = float64(score);
          break;
        }
      }

      if score > bestScore {
        bestScore, bestMove = score, move
      }
    }
  }

  // if we've exhausted the depth set, then we want to return the score at this position
  // i.e the nullScore without saving it in transposition table
  if depth <= 0 && bestScore < nullScore {
    return nullScore;
  }

  s.tp[pos] = entry{
    depth: depth,
    score: bestScore,
    alpha: alpha,
    beta: beta,
    gamma: gamma,
    move: bestMove,
  };
  if len(s.tp) > MaxTableSize {
    s.tp = map[Position]entry{};
  }

  //fmt.Printf("%v %v %v", nullScore, bestScore, bestMove);
  return bestScore;
}
