package main

type Position struct {
    board *Board
    score int
    wc [2]bool
    bc [2]bool
    ep *Square
    kp *Square
}

func (p *Position) Flip() {
    p.board.Flip();
    p.score = -1 * p.score;
    p.ep.Flip();
    p.kp.Flip();

    var temp [2]bool;
    temp[0] = p.wc[0];
    temp[1] = p.wc[1];
    p.wc = [2]bool{p.bc[0], p.bc[1]}
    p.bc = [2]bool{temp[0], temp[1]}
}

