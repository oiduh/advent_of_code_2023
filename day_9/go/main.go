package main

import (
	"fmt"
	"strconv"
	// "math"
	"os"
	"path/filepath"
	"regexp"

	// "strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")

    pattern := regexp.MustCompile(`-?\d+`)

    result_1 := 0
    result_2 := 0
    for _, line := range lines {
        numbers_str := pattern.FindAllString(line, -1)
        numbers_int := []int{}
        for _, number_str := range numbers_str {
            number_int, err := strconv.Atoi(number_str)
            check(err)
            numbers_int = append(numbers_int, number_int)
        }
        diffs := [][]int{}
        diffs = append(diffs, numbers_int)
        current_numbers := numbers_int
        for true {
            diff := []int{}
            for idx := 1; idx < len(current_numbers); idx++ {
                diff = append(diff, current_numbers[idx] - current_numbers[idx-1])
            }
            current_numbers = diff
            diffs = append(diffs, diff)
            if all_zero(diff) {
                for idx, x := range diffs {
                    x = append(x, 0)
                    x = append([]int{0}, x...)
                    diffs[idx] = x
                }
                for idx := 1; idx < len(diffs); idx++ {
                    target_arr := diffs[len(diffs)-idx-1]
                    source_arr := diffs[len(diffs)-idx]
                    target_arr[len(target_arr)-1] = target_arr[len(target_arr)-2] + source_arr[len(source_arr)-1]
                    target_arr[0] = target_arr[1] - source_arr[0]
                    diffs[len(diffs)-idx-1] = target_arr
                }
                result_1 += diffs[0][len(diffs[0])-1]
                result_2 += diffs[0][0]
                break
            }
        }

    }
    fmt.Println(result_1)
    fmt.Println(result_2)
}

func all_zero(arr []int) bool {
    for _, num := range arr {
        if num != 0 {
            return false
        }
    }
    return true
}
