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

type Coordinate struct {
    x, y int
}

type Expansion = int

const(
    SINGLE Expansion = iota
    MULTI
)

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")

    coordinates := []Coordinate{}

    col_count := map[int]Expansion{}
    row_count := map[int]Expansion{}
    for idx := 0; idx < len(lines); idx++ {
        found := false
        for idy := 0; idy < len(lines[0]); idy++ {
            if lines[idx][idy] == '#' {
                found = true
                coordinates = append(coordinates, Coordinate{x: idx, y: idy})
            }
        }
        if found {
            row_count[idx] = SINGLE
        } else {
            row_count[idx] = MULTI
        }
    }
    for idx := 0; idx < len(lines[0]); idx++ {
        found := false
        for idy := 0; idy < len(lines); idy++ {
            if lines[idy][idx] == '#' {
                found = true
                break
            }
        }
        if found {
            col_count[idx] = SINGLE
        } else {
            col_count[idx] = MULTI
        }
    }

    result_1 := 0
    result_2 := 0
    for idx := 0; idx < len(coordinates)-1; idx++ {
        for idy := idx+1; idy < len(coordinates); idy++ {
            min_x := min(coordinates[idx].x, coordinates[idy].x)
            max_x := max(coordinates[idx].x, coordinates[idy].x)
            for x := min_x; x < max_x; x++ {
                if row_count[x] == SINGLE {
                    result_1++
                    result_2++
                } else {
                    result_1 += 2
                    result_2 += 1000000
                }
            }
            min_y := min(coordinates[idx].y, coordinates[idy].y)
            max_y := max(coordinates[idx].y, coordinates[idy].y)
            for y := min_y; y < max_y; y++ {
                if col_count[y] == SINGLE {
                    result_1++
                    result_2++
                } else {
                    result_1 += 2
                    result_2 += 1000000
                }
            }
        }
    }

    fmt.Printf("result 1: %d\n", result_1)
    fmt.Printf("result 2: %d\n", result_2)
}

