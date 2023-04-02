package main

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

  for depth := 1; depth < 2; depth++ {
    alpha, beta := float64(-3 * MateValue), float64(3 * MateValue);
    gamma := int(( alpha + beta + 1 ) / 2);
    s.Search(pos, alpha, beta, gamma, depth);
  }

  return s.tp[pos].move;
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

  return bestScore;
}
