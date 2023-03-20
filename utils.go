package main

import (
    "unicode"
)

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func flipByteCase(b byte) byte {
    if IsUpper(string(b)) {
        return byte(unicode.ToLower(rune(b)));
    }
    return byte(unicode.ToUpper(rune(b)));
}
