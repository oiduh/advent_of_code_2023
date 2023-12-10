package main

import (
	"fmt"
	// "regexp"

	// "math"
	"os"
	"path/filepath"

	// "regexp"
	// "strconv"
	"errors"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Coordinate struct {
    x, y int
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")
    matrix := [][]rune{}
    for _, line := range lines {
        row := []rune{}
        for _, pipe := range line {
            row = append(row, pipe)
        }
        matrix = append(matrix, row)
    }

    start, err := find_S(matrix)
    check(err)

    history := []Coordinate{}
    go_next(start, matrix, &history)

    result_1 := len(history)/2
    fmt.Println(result_1)

    result_2 := 0

    replace_S := 'x'
    if start.y > history[0].y {
        if start.x > history[len(history)-2].x {
            replace_S = 'J'
        } else if start.x < history[len(history)-2].x {
            replace_S = '7'
        } else { 
            replace_S = '-'
        }

    } else if start.y < history[0].y{
        if start.x > history[len(history)-2].x {
            replace_S = 'L'
        } else if start.x < history[len(history)-2].x {
            replace_S = 'F'
        } else { 
            replace_S = '-'
        }
    } else {
        replace_S = '|'
    }
    matrix[start.x][start.y] = replace_S
    for idx, row := range matrix {
        for idy := range row {
            if !contains(history, Coordinate{x: idx, y: idy}) {
                matrix[idx][idy] = 'O'
            }
        }
    }
    
    for idx, row := range matrix {
        inn := false
        for idy, char := range row {
            if char != 'O' {
                cur := matrix[idx][idy]
                if cur == '|' || cur == 'F' || cur == '7' {
                    inn = !inn
                }
            } else {
                if inn {
                    result_2++
                }
            }
        }
    }

    fmt.Println(result_2)
}

func go_next(c Coordinate, m [][]rune, history *[]Coordinate) {
    if len(*history) > 0 && m[c.x][c.y] == 'S' {
        return
    }
    c_l := Coordinate{x: c.x,   y: c.y-1}
    c_u := Coordinate{x: c.x-1, y: c.y}
    c_r := Coordinate{x: c.x,   y: c.y+1}
    c_d := Coordinate{x: c.x+1, y: c.y}
    curr := m[c.x][c.y]
    if c_l.y > 0 && !visited(*history, c_l) && 
    (curr == '-' || curr == 'J' || curr == '7' || curr == 'S') {
        left := m[c_l.x][c_l.y]
        if left == 'F' || left == '-' || left == 'L' || left == 'S' {
            *history = append(*history, c_l)
            go_next(c_l, m, history)
        } 
    }
    if c_u.x > 0 && !visited(*history, c_u) &&
    (curr == '|' || curr == 'J' || curr == 'L' || curr == 'S') {
        up := m[c_u.x][c_u.y]
        if up == 'F' || up == '|' || up == '7' || up == 'S' {
            *history = append(*history, c_u)
            go_next(c_u, m, history)
        } 
    }
    if c_r.y < len(m[0]) && !visited(*history, c_r) &&
    (curr == '-' || curr == 'L' || curr == 'F' || curr == 'S') {
        right := m[c_r.x][c_r.y]
        if right == 'J' || right == '-' || right == '7' || right == 'S' {
            *history = append(*history, c_r)
            go_next(c_r, m, history)
        } 
    }
    if c_d.x < len(m) && !visited(*history, c_d) && 
    (curr == '|' || curr == 'F' || curr == '7' || curr == 'S') {
        down := m[c_d.x][c_d.y]
        if down == 'J' || down == '|' || down == 'L' || down == 'S' {
            *history = append(*history, c_d)
            go_next(c_d, m, history)
        } 
    }
}

func visited(history []Coordinate, c Coordinate) bool {
    for _, x := range history {
        if x == c {
            return true
        }
    }
    return false
}

func find_S(m [][]rune) (Coordinate, error) {
    for idx, row := range m {
        for idy, char := range row {
            if char == 'S' {
                return Coordinate{x: idx, y: idy}, nil
            }
        }
    }
    return Coordinate{x: -1, y: -1}, errors.New("Not found")
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

func contains(arr []Coordinate, c Coordinate) bool {
    for _, cmp := range arr {
        if cmp == c {
            return true
        }
    }
    return false
}
