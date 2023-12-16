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

type Position struct {
    row, col int
}

type Direction int
const(
    NORTH Direction = iota
    EAST
    SOUTH
    WEST
)

func go_next(matrix [][]rune, visited [][]rune, pos Position, direction Direction,
traversed_forks *[]Position) {
    current := visited[pos.row][pos.col]
    loop := (current == '>' && direction == EAST) ||
    (current == '<' && direction == WEST) ||
    (current == '^' && direction == NORTH) ||
    (current == 'v' && direction == SOUTH)
    if loop {
        return
    }

    if matrix[pos.row][pos.col] == '.' {
        switch direction {
        case NORTH:
            visited[pos.row][pos.col] = '^'
        case EAST:
            visited[pos.row][pos.col] = '>'
        case WEST:
            visited[pos.row][pos.col] = '<'
        case SOUTH:
            visited[pos.row][pos.col] = 'v'
        }
    } else {
        *traversed_forks = append(*traversed_forks, pos)
    }
    var next_pos Position
    switch direction {
    case NORTH:
        next_pos = Position{row: pos.row-1, col: pos.col}
        if next_pos.row < 0 {
            return
        }
        switch matrix[next_pos.row][next_pos.col] {
        case '/':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
        case '\\':
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        case '|':
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
        case '-':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        case '.':
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
        }
    case EAST:
        next_pos = Position{row: pos.row, col: pos.col+1}
        if next_pos.col >= len(matrix[0]) {
            return
        }
        switch matrix[next_pos.row][next_pos.col] {
        case '/':
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
        case '\\':
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
        case '|':
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
        case '-':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
        case '.':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
        }
    case SOUTH:
        next_pos = Position{row: pos.row+1, col: pos.col}
        if next_pos.row >= len(matrix) {
            return
        }
        switch matrix[next_pos.row][next_pos.col] {
        case '/':
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        case '\\':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
        case '|':
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
        case '-':
            go_next(matrix, visited, next_pos, EAST, traversed_forks)
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        case '.':
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
        }
    case WEST:
        next_pos = Position{row: pos.row, col: pos.col-1}
        if next_pos.col < 0 {
            return
        }
        switch matrix[next_pos.row][next_pos.col] {
        case '/':
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
        case '\\':
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
        case '|':
            go_next(matrix, visited, next_pos, NORTH, traversed_forks)
            go_next(matrix, visited, next_pos, SOUTH, traversed_forks)
        case '-':
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        case '.':
            go_next(matrix, visited, next_pos, WEST, traversed_forks)
        }
    }
}

func count_lava(visited [][]rune, forks []Position, matrix [][]rune) int {
    for _, pos := range forks {
        visited[pos.row][pos.col] = matrix[pos.row][pos.col]
    }
    result := 0
    for _, row := range visited {
        for _, char := range row {
            if char != '.' {
                result++
            }
        }
    }
    return result
}

func traverse(matrix [][]rune) int {
    traversed_forks := []Position{}
    direction := EAST
    switch matrix[0][0] {
    case '\\':
        direction = SOUTH
    case '|':
        direction = SOUTH
    }
    visited := [][]rune{}
    for _, line := range matrix {
        empty_line := []rune{}
        for range line {
            empty_line = append(empty_line, '.')
        }
        visited = append(visited, empty_line)
    }
    go_next(matrix, visited, Position{row: 0, col: 0}, direction, &traversed_forks)
    return count_lava(visited, traversed_forks, matrix)
}

func clear_visited(visited [][]rune) {
    for idx, row := range visited {
        for idy := range row {
            visited[idx][idy] = '.'
        }
    }
}

func traverse2(matrix [][]rune) int {
    visited := [][]rune{}
    for _, line := range matrix {
        empty_line := []rune{}
        for range line {
            empty_line = append(empty_line, '.')
        }
        visited = append(visited, empty_line)
    }

    result := 0
    cpy1 := make([][]rune, len(matrix))
    cpy2 := make([][]rune, len(matrix))
    copy(cpy1, visited)
    copy(cpy2, visited)
    for i, char := range matrix[0] {
        clear_visited(cpy1)
        clear_visited(cpy2)
        tf_cpy1 := []Position{}
        tf_cpy2 := []Position{}
        if char == '/' {
            go_next(matrix, cpy1, Position{row: 0, col: i}, WEST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '\\' {
            go_next(matrix, cpy1, Position{row: 0, col: i}, EAST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '-' {
            go_next(matrix, cpy1, Position{row: 0, col: i}, EAST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
            go_next(matrix, cpy2, Position{row: 0, col: i}, WEST, &tf_cpy2)
            result = max(result, count_lava(cpy2, tf_cpy2, matrix))
        } else {
            go_next(matrix, cpy1, Position{row: 0, col: i}, SOUTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        }
    }
    last_row := len(matrix)-1
    for i, char := range matrix[last_row] {
        clear_visited(cpy1)
        clear_visited(cpy2)
        tf_cpy1 := []Position{}
        tf_cpy2 := []Position{}
        if char == '/' {
            go_next(matrix, cpy1, Position{row: last_row, col: i}, EAST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '\\' {
            go_next(matrix, cpy1, Position{row: last_row, col: i}, WEST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '-' {
            go_next(matrix, cpy1, Position{row: last_row, col: i}, EAST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
            go_next(matrix, cpy2, Position{row: last_row, col: i}, WEST, &tf_cpy2)
            result = max(result, count_lava(cpy2, tf_cpy2, matrix))
        } else {
            go_next(matrix, cpy1, Position{row: last_row, col: i}, NORTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        }
    }
    for i:= 0; i< len(matrix); i++ {
        char := matrix[i][0]
        clear_visited(cpy1)
        clear_visited(cpy2)
        tf_cpy1 := []Position{}
        tf_cpy2 := []Position{}
        if char == '/' {
            go_next(matrix, cpy1, Position{row: i, col: 0}, NORTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '\\' {
            go_next(matrix, cpy1, Position{row: i, col: 0}, SOUTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '-' {
            go_next(matrix, cpy1, Position{row: i, col: 0}, NORTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
            go_next(matrix, cpy2, Position{row: i, col: 0}, SOUTH, &tf_cpy2)
            result = max(result, count_lava(cpy2, tf_cpy2, matrix))
        } else {
            go_next(matrix, cpy1, Position{row: i, col: 0}, EAST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        }
    }
    last_col := len(matrix[0])-1
    for i:= 0; i< len(matrix); i++ {
        char := matrix[i][last_col]
        clear_visited(cpy1)
        clear_visited(cpy2)
        tf_cpy1 := []Position{}
        tf_cpy2 := []Position{}
        if char == '/' {
            go_next(matrix, cpy1, Position{row: i, col: last_col}, SOUTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '\\' {
            go_next(matrix, cpy1, Position{row: i, col: last_col}, NORTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        } else if char == '-' {
            go_next(matrix, cpy1, Position{row: i, col: last_col}, NORTH, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
            go_next(matrix, cpy2, Position{row: i, col: last_col}, SOUTH, &tf_cpy2)
            result = max(result, count_lava(cpy2, tf_cpy2, matrix))
        } else {
            go_next(matrix, cpy1, Position{row: i, col: last_col}, WEST, &tf_cpy1)
            result = max(result, count_lava(cpy1, tf_cpy1, matrix))
        }
    }

    return result
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
        matrix = append(matrix, []rune(line))
    }

    result_1 := traverse(matrix)
    fmt.Printf("Result 1: %d\n", result_1)

    result_2 := traverse2(matrix)
    fmt.Printf("Result 2: %d\n", result_2)
}
