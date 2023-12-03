package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    path := os.Args[1]
    absPath, err := filepath.Abs(path)
    check(err)
    dat, err := os.ReadFile(absPath)
    check(err)
    dat_string := strings.Trim(string(dat), "\n")
    lines := strings.Split(dat_string, "\n")

    pattern := regexp.MustCompile(`\d+`)

    symbols := map[string]bool{}
    numeric := regexp.MustCompile(`\d`)
    for _, x := range dat_string {
        if !numeric.MatchString(string(x)) {
            symbols[string(x)] = true
        }
    }
    delete(symbols, "\n")
    delete(symbols, ".")

    symbols_list := []string{}
    for x := range symbols {
        symbols_list = append(symbols_list, x)
    }

    matrix := [140][140]([]int){}

    lines = strings.Split(dat_string, "\n")
    for idx, line := range lines {
        for _, span := range pattern.FindAllStringIndex(line, -1) {
            start, end := span[0], span[1]
            number, err := strconv.Atoi(line[start:end])
            check(err)
            if (start > 0 && slices.Contains(symbols_list, string(line[start-1]))) {
                matrix[idx][start-1] = append(matrix[idx][start-1], number)
            }
            if (end < len(line)-1 && slices.Contains(symbols_list, string(line[end]))) {
                matrix[idx][end] = append(matrix[idx][end], number)
            }
            for col := start-1; col < end+1; col++ {
                if !(col > 0 && col < len(line)-1) {
                    continue
                }
                if idx > 0 && slices.Contains(symbols_list, string(lines[idx-1][col])) {
                    matrix[idx-1][col] = append(matrix[idx-1][col], number)
                }
                if idx < len(lines)-1 && slices.Contains(symbols_list, string(lines[idx+1][col])) {
                    matrix[idx+1][col] = append(matrix[idx+1][col], number)
                }
            }
        }
    }

    result_1 := 0
    result_2 := 0
    for ix, x := range matrix {
        for iy, y := range x {
            a := 0
            if len(y) == 2 {
                a = 1
            }
            for _, z := range y {
                result_1 += z
                if string(lines[ix][iy]) == "*" {
                    a *= z
                }
            }
            result_2 += a
        }
    }
    fmt.Println(result_1)
    fmt.Println(result_2)
}
