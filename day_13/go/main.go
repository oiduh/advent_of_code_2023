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

type SwapLines struct {
    first, second int
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    blocks := strings.Split(data_string, "\n\n")

    part := 2

    result_1 := 0
    result_2 := 0
    for i, block := range blocks {
        fmt.Println(i)
        rows := strings.Split(block, "\n")
        if part == 2 {
            original_rows := row_check(rows)
            original_cols := col_check(rows)
            potential_row_swaps := rows_fix_new(rows, "row")
            potential_col_swaps := cols_fix_new(rows, "col")
            potential_row_swaps = append(potential_row_swaps, rows_fix_old(rows)...)
            potential_col_swaps = append(potential_col_swaps, cols_fix_old(rows)...)
            fmt.Printf("new %d %d\n", potential_row_swaps, potential_col_swaps)
            var y []SwapLines
            a := 0
            for _, swap := range potential_row_swaps {
                cpy := make([]string, len(rows))
                copy(cpy, rows)
                cpy[swap.second] = cpy[swap.first]
                c := 0
                y = row_check(cpy)
                skip := false
                for _, yy := range y {
                    if yy.first == -1 {
                        continue
                    }
                    for _, xy := range original_rows {
                        if xy.first == yy.first {
                            skip = true
                            break
                        }
                    }
                    if skip {
                        continue
                    }
                    c = yy.second
                }
                a = c * 100
            }
            var x []SwapLines
            b := 0
            for _, swap := range potential_col_swaps {
                tmp := rows_to_cols(rows)
                cpy := make([]string, len(tmp))
                copy(cpy, tmp)
                cpy[swap.second] = cpy[swap.first]
                cpy = rows_to_cols(cpy)
                d := 0
                x = col_check(cpy)
                skip := false
                for _, xx := range x {
                    if xx.first == -1 {
                        continue
                    }
                    for _, xy := range original_cols {
                        if xy.first == xx.first {
                            skip = true
                            break
                        }
                    }
                    if skip {
                        continue
                    }
                    d = xx.second
                }
                b = d
            }
            fmt.Printf("row val: %d - col val: %d\n\n", a, b)
            result_2 += a + b
        } else {
            a := 0
            y := row_check(rows)
            for _, yy := range y {
                if yy.first >= 0 {
                    a = yy.second * 100
                }
            }
            b := 0
            x := col_check(rows)
            for _, xx := range x {
                if xx.first >= 0 {
                    b = xx.second
                }
            }
            result_1 += a + b
        }
    }

    fmt.Printf("Result 1: %d\n", result_1)
    fmt.Printf("Result 2: %d\n", result_2)
}

func col_check(lines []string) []SwapLines {
    rows_t := rows_to_cols(lines)
    return row_check(rows_t)
}

func row_check(rows []string) []SwapLines {
    res := []SwapLines{}
    // find 2 same rows
    var idx int
    found := false
    candidates := []int{}
    for idx = 0; idx < len(rows)-1; idx++ {
        if rows[idx] == rows[idx+1] {
            fmt.Printf("row %d and row %d are the same\n", idx, idx+1)
            found = true
            candidates = append(candidates, idx)
        }
    }
    if !found {
        res = append(res, SwapLines{first: -1, second: -1})
        return res
    }
    // count symmetry
    fmt.Printf("candidates %d\n", candidates)
    reflection_line := -1
    for _, candidate := range candidates {
        left, right := candidate, candidate+1
        for rows[left] == rows[right] {
            if left == 0 || right == len(rows)-1 {
                reflection_line = candidate
                res = append(res, SwapLines{first: candidate, second: candidate+1})
                break
            }
            left--; right++
        }
    }
    // fmt.Printf("start %d, end %d\n", left, right)
    if reflection_line == -1 {
        res = append(res, SwapLines{first: -1, second: -1})
        return res
    } 
    return res
}

func rows_fix_new(lines []string, type_ string) []SwapLines {
    candidates := []SwapLines{}
    for idx := 0; idx < len(lines)-1; idx++ {
        first, next := lines[idx], lines[idx+1]
        diff := 0
        for idy := 0; idy < len(first); idy++ {
            if first[idy] != next[idy] {
                diff++
            }
        }
        if diff == 1 {
            fmt.Printf("new %s found at %d\n", type_, idx)
            candidates = append(candidates, SwapLines{first: idx, second: idx+1})
        }
    }
    return candidates
}

func cols_fix_new(lines []string, type_ string) []SwapLines {
    lines_t := rows_to_cols(lines)
    return rows_fix_new(lines_t, type_)
}

func rows_fix_old(lines []string) []SwapLines {
    candidates := []SwapLines{}

    found := false
    tmp := []int{}
    var idx int
    for idx = 0; idx < len(lines)-1; idx++ {
        if lines[idx] == lines[idx+1] {
            found = true
            tmp = append(tmp, idx)
        }
    }
    if !found {
        return candidates
    }

    // count symmetry
    for _, candidate := range tmp {
        left, right := candidate, candidate+1
        skip := false
        for lines[left] == lines[right] {
            if left == 0 || right == len(lines)-1 {
                skip = true
                break
            } 
            left--; right++
        }
        if skip {
            continue
        }
        diff := 0
        for idx := 0; idx < len(lines[left]); idx++ {
            if lines[left][idx] != lines[right][idx] {
                diff++
            }
        }
        if diff == 1 {
            candidates = append(candidates, SwapLines{first: left, second: right})
        }
    }
    return candidates
}

func cols_fix_old(lines []string) []SwapLines {
    lines_t := rows_to_cols(lines)
    return rows_fix_old(lines_t)
}

func rows_to_cols(rows []string) []string {
    matrix := []([]rune){}
    for _, row := range rows {
        matrix = append(matrix, []rune(row))
    }

    cols := []string{}

    matrix_t := transpose(matrix)
    for _, chars := range matrix_t {
        cols = append(cols, string(chars))
    }

    return cols
}

func transpose(slice [][]rune) [][]rune {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]rune, xl)
    for i := range result {
        result[i] = make([]rune, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}
