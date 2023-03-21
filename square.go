package main

type Square int
func (s *Square) Flip() {
    *s = Square(119 - *s);
}
