package main

type Move struct {
    from *Square
    to *Square
}

type Moves map[Square] []Square;
