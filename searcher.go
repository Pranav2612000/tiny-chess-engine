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
  s.nodes = 0;

  for depth := 1; depth < 99; depth++ {
    alpha, beta := 3 * MateValue, -3 * MateValue;
    score := 0
    for beta < alpha - 13 {
      gamma := int(( alpha + beta + 1 ) / 2);
      score := s.Search(pos, alpha, beta, gamma, depth);

      if score >= gamma {
        beta = score;
      }

      if score < gamma {
        alpha = score;
      }
    }

    if Abs(score) >= MateValue || s.nodes >= maxNodes {
      break;
    }
  }

  if (GlobalIsDebugMode) {
    fmt.Printf("Best Move| From: %v , To: %v\n", s.tp[pos].move.from, s.tp[pos].move.to);
    fmt.Printf("Move score: %d", s.tp[pos].score - pos.score);
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

func (s *Searcher) Search(pos Position, alpha int, beta int, gamma int, depth int) (score int) {
  s.nodes++;

  e, ok := s.tp[pos];
  if ok && e.depth >= depth && ((e.score < e.gamma && e.score < gamma) ||
      (e.score >= e.gamma && e.score >= gamma)) {
    return e.score;
  }

  if Abs(pos.score) > MateValue {
    return pos.score;
  }

  nullScore := pos.score;
  if depth > 0 {
    flippedPos := pos.Copy();
    flippedPos.Flip();

    nullScore = -1 * s.Search(flippedPos, alpha, beta, 1 - gamma, depth - 3);
  }

  if nullScore >= gamma {
    return nullScore
  }

  bestScore, bestMove := -3*MateValue, Move{};

  allMoves := pos.Moves()
  for start, moves := range allMoves {
    for _, sq := range moves {
      startCp := start.Copy();
      sqCp := sq.Copy();
      move := Move{ from: &startCp, to: &sqCp} 

      
      if depth <= 0 && pos.GetValueOfMove(move) < 150 {
        break;
      }

      flippedPos := pos.Copy();
      flippedPos.Move(move);
      flippedPos.Flip();

      score := -1 * s.Search(flippedPos, alpha, beta, 1 - gamma, depth - 1);

      if score > bestScore {
        bestScore, bestMove = score, move
      }
      if score >= gamma {
        break;
      }
    }
  }

  // if we've exhausted the depth set, then we want to return the score at this position
  // i.e the nullScore without saving it in transposition table
  if depth <= 0 && bestScore < nullScore {
    return nullScore;
  }

  // Stalemate check: best move loses king + null move is better
  if depth > 0 && bestScore <= -MateValue && nullScore > -MateValue {
    bestScore = 0
  }

  if !ok || depth >= e.depth && bestScore >= gamma {
    s.tp[pos] = entry{
      depth: depth,
      score: bestScore,
      alpha: float64(alpha),
      beta: float64(beta),
      gamma: gamma,
      move: bestMove,
    };
    if len(s.tp) > MaxTableSize {
      s.tp = map[Position]entry{};
    }
  }

  return bestScore;
}
