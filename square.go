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
