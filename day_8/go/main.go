package main

import (
	"fmt"
	"regexp"
	// "math"
	"os"
	"path/filepath"

	// "regexp"
	// "strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Fork struct {
    left, right string
}

func end_z(a []string) bool {
    for idx := 0; idx < len(a); idx++ {
        last_rune := rune(a[idx][2])
        if last_rune != 'Z' {
            return false
        }
    }
    return true
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := string(data)
    tmp := strings.Split(data_string, "\n\n")
    instructions := strings.Trim(tmp[0], "\n")
    network := strings.Trim(tmp[1], "\n")

    network_map := map[string]Fork{}

    pattern := regexp.MustCompile(`\w+`)
    for _, x := range strings.Split(network, "\n") {
        y := pattern.FindAllString(x, -1)
        if len(y) != 3 {
            panic("not 3")
        }
        fork := Fork{left: y[1], right: y[2]}
        network_map[y[0]] = fork
    }

    current := "AAA"
    end := "ZZZ"
    steps := 0
    found := false
    for !found {
        for _, x := range instructions {
            choice := string(x)
            if choice == "L" {
                current = network_map[current].left
            } else {
                current = network_map[current].right
            }
            steps += 1
            if current == end {
                found = true
                break
            }
        }
    }
    fmt.Println(steps)

    currents := []string{}
    reached_z := []bool{}
    for k := range network_map {
        last_rune := rune(k[2])
        if last_rune == 'A' {
            currents = append(currents, k)
            reached_z = append(reached_z, false)
        }
    }

    found = false
    steps = 0
    steps_to_reach_z := []int{}
    for !found {
        for _, choice := range instructions {
            if choice == 'L' {
                for idx := 0; idx < len(currents); idx++ {
                    currents[idx] = network_map[currents[idx]].left
                }
            } else {
                for idx := 0; idx < len(currents); idx++ {
                    currents[idx] = network_map[currents[idx]].right
                }
            }
            steps++

            for idx := 0; idx < len(currents); idx++ {
                if rune(currents[idx][2]) == 'Z' {
                    reached_z[idx] = true
                    steps_to_reach_z = append(steps_to_reach_z, steps)
                }
            }

            if all_true(reached_z) {
                found = true
                break 
            }
        }
    }
    // after each starting point reached z at least once ->
    // least common multiple to calculate steps instead of brute forcing
    fmt.Println(LCM(steps_to_reach_z[0], steps_to_reach_z[1], steps_to_reach_z[2:]...))
}

func all_true(a []bool) bool {
    for _, x := range a {
        if !x {
            return false
        }
    }
    return true
}

func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)
    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}
