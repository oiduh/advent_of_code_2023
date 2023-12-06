package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type RangeMap struct {
    src, dst, range_ int
}

type A_to_B_Transformer struct {
    range_maps []RangeMap
}

type Full_Transformer struct {
    transformers []A_to_B_Transformer
}


func main() {
    path := os.Args[1]
    absPath, err := filepath.Abs(path)
    check(err)
    dat, err := os.ReadFile(absPath)
    check(err)
    dat_string := strings.Trim(string(dat), "\n")
    blocks := strings.Split(dat_string, "\n\n")
    seed_block := blocks[0]
    pattern := regexp.MustCompile(`\d+`)
    seeds := pattern.FindAllString(seed_block, -1)
    range_maps := blocks[1:]

    transformers := Full_Transformer{}
    for idx, range_map := range range_maps {
        split := strings.Split(range_map, "\n")
        transformers.transformers = append(transformers.transformers, A_to_B_Transformer{})
        for _, x := range split[1:] {
            numbers_str := pattern.FindAllString(x, -1)
            numbers_int := []int{}
            for _, z := range numbers_str {
                num, err := strconv.Atoi(z)
                check(err)
                numbers_int = append(numbers_int, num)
            }
            y := RangeMap{src: numbers_int[1], dst: numbers_int[0], range_: numbers_int[2]}
            transformers.transformers[idx].range_maps = append(transformers.transformers[idx].range_maps, y)
        }
    }
    seeds_int_1 := []int{}
    for _, x := range seeds {
        num, err := strconv.Atoi(x)
        check(err)
        seeds_int_1 = append(seeds_int_1, num)
    }
    fmt.Println(len(seeds_int_1))

    min_val_1 := math.MaxInt
    for idx := 0; idx < len(seeds_int_1); idx++ {
        for _, transformer := range transformers.transformers {
            for _, mapper := range transformer.range_maps {
                if mapper.src <= seeds_int_1[idx] && seeds_int_1[idx] < mapper.src + mapper.range_ {
                    new_val := seeds_int_1[idx] - mapper.src + mapper.dst
                    seeds_int_1[idx] = new_val
                    break
                }
            }
        }
        if seeds_int_1[idx] < min_val_1 {
            min_val_1 = seeds_int_1[idx]
        }
    }
    fmt.Printf("Result 1: %d\n", min_val_1)

    // potential candidates
    seeds_int_2 := []int{}
    ranges := []int{}
    seeds_int_3 := []int{}
    for x := 0; x < len(seeds); x++ {
        num, err := strconv.Atoi(seeds[x])
        check(err)
        if x % 2 == 0 {
            seeds_int_2 = append(seeds_int_2, num)
            seeds_int_3 = append(seeds_int_3, num)
        } else {
            ranges = append(ranges, num)
        }
    }
    fmt.Println(len(seeds_int_2))
    fmt.Println(seeds_int_2)
    fmt.Println(ranges)

    x := transformers.transformers[0].range_maps 
    sort.Slice(x[:], func (i, j int) bool {
        return x[i].src < x[j].src
    })
    fmt.Println(len(x))
    for idx := 0; idx < len(seeds_int_2); idx++ {
        seed := seeds_int_2[idx]
        range_ := ranges[idx]
        for _, y := range x {
            if seed + range_ < y.src || seed > y.src + y.range_ {
                continue
            }
            if seed < y.src {
                if seed + range_ >= y.src {
                    seeds_int_3 = append(seeds_int_3, y.src)
                }
                if seed + range_ >= y.src + y.range_ {
                    seeds_int_3 = append(seeds_int_3, y.src + y.range_)
                }
                continue
            }
            if seed > y.src && seed < y.src + y.range_ {
                if seed + range_ >= y.src + y.range_ {
                    seeds_int_3 = append(seeds_int_3, y.src + y.range_)
                }
            }
        }
    }
    fmt.Println(len(seeds_int_3))

    min_val_2 := math.MaxInt
    for idx := 0; idx < len(seeds_int_3); idx++ {
        for _, transformer := range transformers.transformers {
            for _, mapper := range transformer.range_maps {
                if mapper.src <= seeds_int_3[idx] && seeds_int_3[idx] < mapper.src + mapper.range_ {
                    seeds_int_3[idx] = mapper.dst
                    break
                }
            }
        }
        if seeds_int_3[idx] < min_val_2 {
            min_val_2 = seeds_int_3[idx]
        }
    }
    fmt.Printf("Result 2: %d\n", min_val_2)


    h := transformers.transformers[len(transformers.transformers)-1].range_maps 
    sort.Slice(h[:], func (i, j int) bool {
        return h[i].dst < h[j].dst
    })
    fmt.Println(h[0].dst)
    fmt.Println(h[0].range_)

    a := h[0].dst
    b := h[0].dst + h[0].range_ - 1
    arr := []int{a, b}
    for idy := 0; idy < 2; idy++ {
        for idx := len(transformers.transformers) - 1; idx >= 0; idx-- {
            for _, mapper := range transformers.transformers[idx].range_maps {
                if mapper.dst <= arr[idy] && arr[idy] < mapper.dst + mapper.range_ {
                    new_val := arr[idy] - mapper.dst + mapper.src
                    arr[idy] = new_val
                    break
                }
            }
        }
    }

    for idx := 0; idx < len(seeds_int_2); idx++ {
        if arr[0] >= seeds_int_2[idx] && arr[0] < seeds_int_2[idx] + ranges[idx] {
            fmt.Println("the lowest destination can be found")
        }
    }
    
    fmt.Println(arr)

}
