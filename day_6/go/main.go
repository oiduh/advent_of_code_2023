package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
    data_string := string(data)
    lines := strings.Split(data_string, "\n")
    pattern := regexp.MustCompile(`\d+`)
    times := pattern.FindAllString(lines[0], -1)
    distances := pattern.FindAllString(lines[1], -1)

    times_int := []int{}
    distances_int := []int{}

    result_1 := 1
    for i := 0; i < len(times); i++ {
        time_int, err := strconv.Atoi(times[i])
        check(err)
        distance_int, err := strconv.Atoi(distances[i])
        check(err)
        times_int = append(times_int, time_int)
        distances_int = append(distances_int, distance_int)
        var res int
        if time_int % 2 == 0 {
            res = (time_int >> 1) ^ 2
            if !(res > distance_int) {
                res = 0
            }
        } else {
            res = 0
        }
        for j := 1; j <= int(math.Floor(float64(time_int)/2)); j++ {
            x := (time_int-j)*j 
            if x > distance_int {
                res += time_int-1 - 2*(j-1)
                break
            }
        }
        result_1 *= res
    }
    fmt.Println(result_1)

    y := ""
    z := ""
    for i := 0; i < len(times); i++ {
        y = y + times[i]
        z = z + distances[i]
    }
    new_time, err := strconv.Atoi(y)
    check(err)
    new_distance, err := strconv.Atoi(z)
    check(err)
    var res2 int
    if new_time % 2 == 0 {
        res2 = (new_time >> 1) ^ 2
        if !(res2 > new_time) {
            res2 = 0
        }
    } else {
        res2 = 0
    }
    for j := 1; j <= int(math.Floor(float64(new_time)/2)); j++ {
        x := (new_time-j)*j 
        if x > new_distance {
            res2 += new_time-1 - 2*(j-1)
            break
        }
    }
    fmt.Println(res2)
}
