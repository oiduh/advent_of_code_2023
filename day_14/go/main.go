package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Direction int
const (
    NORTH Direction = iota
    WEST
    SOUTH
    EAST
)

type cycle struct {
    matrix [][]rune
    count int
    indices []int
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")
    matrix_og := [][]rune{}
    for _, line := range lines {
        matrix_og = append(matrix_og, []rune(line))
    }

    // result 1
    matrix_1 := make([][]rune, len(matrix_og))
    copy(matrix_1, matrix_og)
    matrix_1 = tilt(matrix_1, NORTH)
    result_1 := count_score(matrix_1)
    fmt.Printf("Result 1: %d\n", result_1)

    // result 2
    matrix_2 := make([][]rune, len(matrix_og))
    copy(matrix_2, matrix_og)
    iterations := 1000000000
    history := map[string]cycle{}
    idx := 0
    repeat := false
    for idx < iterations {
        key := generate_key(matrix_2)
        if found, ok := history[key]; ok && !repeat {
            if found.count >= 2 {
                repeat = true
                diff := found.indices[len(found.indices)-1] - found.indices[len(found.indices)-2]
                // spins left
                idx = iterations - (iterations-found.indices[len(found.indices)-1])%diff
                continue
            }
        }
        matrix_2 = tilt(matrix_2, NORTH)
        matrix_2 = tilt(matrix_2, WEST)
        matrix_2 = tilt(matrix_2, SOUTH)
        matrix_2 = tilt(matrix_2, EAST)
        if found, ok := history[key]; ok {
            found.count++
            found.indices = append(found.indices, idx)
            history[key] = found
        } else {
            history[key] = cycle{matrix: matrix_2, count: 1, indices: []int{idx}}
        }
        idx++
    }
    result_2 := count_score(matrix_2)
    fmt.Printf("Result 2: %d\n", result_2)
}

func generate_key(matrix [][]rune) string {
    key := ""
    for _, line := range matrix {
        key += string(line)
    }
    return key
}
func count_score(matrix [][]rune) int {
    result := 0
    for idx := 0; idx < len(matrix); idx++ {
        for idy := 0; idy < len(matrix[0]); idy++ {
            if matrix[idy][idx] == 'O' {
                result += len(matrix)-idy
            }
        }
    }
    return result
}

func tilt(matrix [][]rune, direction Direction) [][]rune {
    if direction == NORTH || direction == SOUTH {
        matrix = transpose(matrix)
    }
    pattern_1 := regexp.MustCompile(`(\.|O)+`)

    for idx := 0; idx < len(matrix); idx++ {
        tmp1 := pattern_1.FindAllStringIndex(string(matrix[idx]), -1)
        for _, x := range tmp1 {
            to_sort := strings.Split(string(matrix[idx][x[0]:x[1]]), "")
            if direction == NORTH || direction == WEST {
                sort.Sort(sort.Reverse(sort.StringSlice(to_sort)))
            } else {
                sort.Sort(sort.StringSlice(to_sort))
            }
            sorted := strings.Join(to_sort, "")
            copy(matrix[idx][x[0]:x[1]], []rune(sorted))
        }
    }

    if direction == NORTH || direction == SOUTH {
        matrix = transpose(matrix)
    }

    return matrix
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
