package main

type Square struct {
    position int
    piece *Piece
    isPlayable bool
}

func (s *Square) Flip() {
    s.position = 119 - s.position;
    if s.piece != nil {
        s.piece.Flip()
    }
}

func (s *Square) PSTValue() int {
    if s.piece == nil {
        return 0;
    }
    return s.piece.PSTValue(s.position);
}
