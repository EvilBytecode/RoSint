package profile

import (
    "strings"
    "unicode"
)

func IsNumeric(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}

func CleanWord(word string) string {
    for _, r := range word {
        if unicode.IsPunct(r) {
            word = strings.ReplaceAll(word, string(r), " ")
        }
    }
    return word
}
