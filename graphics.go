package main

import (
    "fmt"
    "golang.org/x/exp/slices"
)

func DrawBoard(b *Board) {
    for i := 1; i <= 8; i = i + 1 {
        fmt.Print("        ——— ——— ——— ——— ——— ——— ——— ———\n");

        fmt.Printf("%d   %d ",10 * (10 - i), (9 - i));
        for j := 1; j <= 8; j = j + 1 {
            var ch rune;

            if b[10 * (10 - i) + j].piece == nil {
                ch = ' ';
            } else {
                ch = rune(b[10 * (10 - i) + j].piece.variant);
            }
            fmt.Printf("| %c ", ch);
        }
        fmt.Print("|\n");
    }
    fmt.Print("        ——— ——— ——— ——— ——— ——— ——— ———\n");
    fmt.Print("         a   b   c   d   e   f   g   h \n");
    fmt.Print("         1   2   3   4   5   6   7   8\n");
}

func DrawBoardWithMoves(b *Board, moves []Square) {
    for i := 1; i <= 8; i = i + 1 {
        fmt.Print("        ——— ——— ——— ——— ——— ——— ——— ———\n");

        fmt.Printf("%d   %d ",10 * (10 - i), (9 - i));
        for j := 1; j <= 8; j = j + 1 {
            var ch rune;
            linearIndex := 10 * (10 - i) + j;

            if slices.IndexFunc(moves, func(sq Square) bool { return sq.position == linearIndex }) != -1 {
                ch = 'x';
            } else if b[linearIndex].piece == nil {
                ch = ' ';
            } else {
                ch = rune(b[linearIndex].piece.variant);
            }
            fmt.Printf("| %c ", ch);
        }
        fmt.Print("|\n");
    }
    fmt.Print("        ——— ——— ——— ——— ——— ——— ——— ———\n");
    fmt.Print("         a   b   c   d   e   f   g   h \n");
    fmt.Print("         1   2   3   4   5   6   7   8\n");
}
