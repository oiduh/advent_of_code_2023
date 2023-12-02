package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


type DigitString struct {
    digit string
    index int
    number int
}

type DigitInt struct {
    digit string
    index int
    number int
}

func main() {
    path := os.Args[1]
    absPath, err := filepath.Abs(path)
    check(err)
    dat, err := os.ReadFile(absPath)
    dat_string := string(dat)
    digits_str := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    digits_int := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

    sums_1 := 0
    sums_2 := 0
    lines := strings.Split(dat_string, "\n")
    for _, line := range lines {
        firsts := []int{}
        first_string := DigitString{digit: "", index: 999}
        for idx, digit := range digits_str {
            index := strings.Index(line, digit)
            if index > -1 && index < first_string.index {
                first_string.index = index
                first_string.digit = digit
                first_string.number = idx + 1
            }
        }
        first_int := DigitInt{digit: "", index: 999}
        for idx, digit := range digits_int {
            index := strings.Index(line, digit)
            if index > -1 && index < first_int.index {
                first_int.index = index
                first_int.digit = digit
                first_int.number = idx + 1
            }
        }
        if first_string.index < first_int.index {
            firsts = append(firsts, first_string.number)
        } else {
            firsts = append(firsts, first_int.number)
        }

        lasts := []int{}
        last_string := DigitString{digit: "", index: 999}
        for idx, digit := range digits_str {
            index := strings.Index(Reverse(line), Reverse(digit))
            if index > -1 && index < last_string.index {
                last_string.index = index
                last_string.digit = digit
                last_string.number = idx + 1
            }
        }
        last_int := DigitInt{digit: "", index: 999}
        for idx, digit := range digits_int {
            index := strings.Index(Reverse(line), Reverse(digit))
            if index > -1 && index < last_int.index {
                last_int.index = index
                last_int.digit = digit
                last_int.number = idx + 1
            }
        }
        if last_string.index < last_int.index {
            lasts = append(lasts, last_string.number)
        } else {
            lasts = append(lasts, last_int.number)
        }

        sums_1 += first_int.number*10 + last_int.number
        sums_2 += firsts[0]*10 + lasts[0]
    }
    fmt.Printf("part 1 result = %d\n", sums_1)
    fmt.Printf("part 2 result = %d\n", sums_2)
}
