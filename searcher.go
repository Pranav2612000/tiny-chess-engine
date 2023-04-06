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

type PositionRaw struct {
  board string 
  score int
  wc [2]bool
  bc [2]bool
  ep Square
  kp Square
  turn bool
}

type Searcher struct {
  tp map[PositionRaw]entry // transposition table to keep track of best move at a position
  nodes int // Number of nodes to consider, to keep a check on the complexity of the calculations
}

func (s *Searcher) SearchMove(pos Position, maxNodes int) (m Move) {
  s.nodes = 0;

  for depth := 1; depth < 4; depth++ {
    alpha, beta := float64(-3 * MateValue), float64(3 * MateValue);
    gamma := int(( alpha + beta + 1 ) / 2);
    s.Search(pos, alpha, beta, gamma, depth);
  }

  posRaw := pos.CopyRaw();
  fmt.Printf("Current pos: %v\n\n", posRaw);
  if (GlobalIsDebugMode) {
    _, ok := s.tp[posRaw];
    fmt.Printf("Chosen TP %v \n Tp exists: %v", s.tp[posRaw], ok);
    fmt.Printf("Best Move| From: %v , To: %v\n", s.tp[posRaw].move.from, s.tp[posRaw].move.to);
    fmt.Printf("Move score: %d", s.tp[posRaw].score - pos.score);
  }

  return s.tp[posRaw].move;
}

func (s *Searcher) SearchNew(pos Position, depth int, alpha, beta float64) (searchScore float64) {
  score := float64(pos.score);
  if (pos.score < 0) {
    score = -1 * score;
  }

  if score >= float64(MateValue) {
    return score;
  }

  posRaw := pos.CopyRaw();
  e, ok := s.tp[posRaw];
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

  s.tp[posRaw] = entry{
    depth: depth,
    score: int(bestScore),
    alpha: alpha,
    beta: beta,
    gamma: 0,
    move: bestMove,
  };
	return alpha;
}

func (s *Searcher) Search(pos Position, alpha float64, beta float64, gamma, depth int) (score int) {

  s.nodes++;
  posRaw := pos.CopyRaw();

  if Abs(pos.score) > MateValue {
    return pos.score;
  }

  bestScore, bestMove := -3*MateValue, Move{};

  allMoves := pos.Moves()
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

      if score > bestScore {
        bestScore, bestMove = score, move
      }
    }
  }

  s.tp[posRaw] = entry{
    depth: depth,
    score: bestScore,
    alpha: alpha,
    beta: beta,
    gamma: gamma,
    move: bestMove,
  };

  return bestScore;
}
