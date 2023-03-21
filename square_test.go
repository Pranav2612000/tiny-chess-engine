package main

import (
    "testing"
)

func TestSquareFlip(t *testing.T) {
    var square Square = 45;
    square.Flip();
    if square != 74 {
        t.Error(`Failed to Flip Square`);
    }
}
