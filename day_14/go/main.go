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
    iterations_full := 1000000000
    iterations :=  5000
    history := map[int]([]int){}
    for idx := 0; idx < iterations; idx++ {
        matrix_2 = tilt(matrix_2, NORTH)
        matrix_2 = tilt(matrix_2, WEST)
        matrix_2 = tilt(matrix_2, SOUTH)
        matrix_2 = tilt(matrix_2, EAST)
        tmp := count_score(matrix_2)
        history[tmp] = append(history[tmp], idx)
        skip := false
        for _, y := range history {
            if len(y) > 100 {
                skip = true
            }
        }
        if skip {
            break
        }
    }
    for x, y := range history {
        if len(y) > 40 {
            fmt.Printf("%d %d\n", x, y[len(y)-10:])
            calc_next(y[len(y)-10:])
            fmt.Println((iterations_full-y[len(y)-1])%11)
        }
    }
    fmt.Println(iterations_full%11)

    result_2 := count_score(matrix_2)
    fmt.Printf("Result 2: %d\n", result_2)

}

func calc_next(arr []int) int {
    tmp := make([]int, len(arr))
    copy(tmp, arr)
    x := [][]int{}
    for !all_zero(tmp) {
        y := []int{}
        for idx := 0; idx < len(tmp)-1; idx++ {
            y = append(y, tmp[idx+1]-tmp[idx])
        }
        x = append(x, y)
        tmp = y
    }

    for _, a := range x {
        fmt.Println(a)
    }


    return 0
}

func all_zero(arr []int) bool {
    for _, item := range arr {
        if item != 0 {
            return false
        }
    }
    return true
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
