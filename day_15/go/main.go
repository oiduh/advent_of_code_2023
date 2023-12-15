package main

import (
	"fmt"
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

type Item struct {
    label string
    focal_length int
}

func insert (hashmap map[int]([]Item), new_item Item, box int, operation rune) {
    items := hashmap[box]
    if operation == '-' {
        for i, item := range items {
            if item.label == new_item.label {
                if i == len(items)-1 {
                    items = items[:len(items)-1]
                } else if i == 0 {
                    items = items[1:]
                } else {
                    left, right := items[:i], items[i+1:]
                    items = append(left, right...)
                }
                hashmap[box] = items
                return
            }
        }
    } else if operation == '='{
        for i, item := range items {
            if item.label == new_item.label {
                items[i] = new_item
                hashmap[box] = items
                return
            }
        }
        items = append(items, new_item)
        hashmap[box] = items
        return
    } else {
        panic("invalid operation")
    }
}

func calculate_result(hashmap map[int]([]Item)) int {
    result := 0
    for box, items := range hashmap {
        for i, item := range items {
            result += (box+1)*(i+1)*item.focal_length
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
    codes := strings.Split(data_string, ",")
    pattern := regexp.MustCompile(`[a-z]+`)

    HASHMAP := map[int]([]Item){}
    result_1 := 0
    for _, code := range codes {
        // part 1
        tmp := 0
        for _, char := range code {
            tmp += int(char)
            tmp *= 17
            tmp %= 256
        }
        result_1 += tmp
        // part 2
        split := pattern.FindAllStringIndex(code, -1)[0]
        label := code[:split[1]]
        rest := code[split[1]:]
        operation := '-'
        number := -1
        if rune(rest[0]) == '=' {
            operation = '='
            num_str := code[split[1]+1:]
            number, err = strconv.Atoi(num_str)
            check(err)
        }
        box := 0
        for _, char := range label {
            box += int(char)
            box *= 17
            box %= 256
        }
        new_item := Item{label: label, focal_length: number}
        insert(HASHMAP, new_item, box, operation)
    }

    fmt.Printf("Result 1: %d\n", result_1)

    result_2 := calculate_result(HASHMAP)
    fmt.Printf("Result 2: %d\n", result_2)
}
