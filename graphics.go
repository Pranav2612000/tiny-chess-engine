package main

import (
    "fmt"
)

func DrawBoard(b *Board) {
    for i := 1; i <= 8; i = i + 1 {
        fmt.Print("        —— —— —— —— —— —— —— ——\n");

        fmt.Printf("%d   %d ",10 * (10 - i), i);
        for j := 0; j < 8; j = j + 1 {
            fmt.Print("|  ");
        }
        fmt.Print("|\n");
    }
    fmt.Print("        —— —— —— —— —— —— —— ——\n");
    fmt.Print("        a  b  c  d  e  f  g  h \n");
    fmt.Print("        1  2  3  4  5  6  7  8\n");
}
