package profile

import (
    "strings"
)

func ValidWord(word string, requirements map[string]interface{}) bool {
    length, ok := requirements["length"].(int)
    if !ok {
        return false
    }
    if len(word) == length || length == 999 {
        if end, exists := requirements["end"].(map[string]interface{}); exists {
            return ValidEnd(word, end)
        }
        if has, exists := requirements["has"].([]string); exists {
            return ValidHas(word, has)
        }
        number, _ := requirements["number"].(bool)
        return IsNumeric(word) == number
    }
    return false
}

func ValidEnd(word string, endDict map[string]interface{}) bool {
    length, ok := endDict["length"].(int)
    if !ok || len(word) <= length {
        return false
    }
    number, _ := endDict["number"].(bool)
    return IsNumeric(word[len(word)-length:]) == number
}

func ValidHas(word string, hasList []string) bool {
    word = CleanWord(word)

    for _, w := range strings.Split(word, " ") {
        for _, need := range hasList {
            if strings.EqualFold(need, w) {
                return true
            }
        }
    }
    return false
}

func CheckWord(name, word string, requirement map[string]interface{}) map[string]interface{} {
    if ValidWord(word, requirement) {
        return map[string]interface{}{name: requirement["value"]}
    }
    return nil
}
