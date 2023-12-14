package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	// "regexp"
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
    lines_2 := []string{}
    for _, line := range lines {
        split := strings.Split(line, " ")
        text, numbers := split[0], split[1]
        new_text, new_numbers := text, numbers
        for x := 0; x < 4; x++   {
            new_text += fmt.Sprintf("?%s", text)
            new_numbers += fmt.Sprintf(",%s", numbers)
        }
        lines_2 = append(lines_2, fmt.Sprintf("%s %s", new_text, new_numbers))
    }

    result_1 := 0
    for _, line := range lines {
        split := strings.Split(line, " ")
        text := split[0]
        nums := []int{}
        for _, num := range strings.Split(split[0], ",") {
            tmp, err := strconv.Atoi(num)
            check(err)
            nums = append(nums, tmp)
        }

        result_1 += f([]rune(text), nums, 0)
    }

    // qm_pattern := regexp.MustCompile(`\?+`)
    // hs_pattern := regexp.MustCompile(`\#+`)
    // num_pattern := regexp.MustCompile(`\d+`)
}

func f(chars []rune, nums []int, count int) int {
    return 0
}

