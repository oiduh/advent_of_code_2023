package main

import (
	"fmt"
	"sort"
	"strconv"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Combo int
const (
    FiveKind Combo = iota
    FourKind
    FullHouse
    ThreeKind
    TwoPair
    OnePair
    HighCard
)

type Hand struct {
    cards []int
    bid int
    combo Combo
}


func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")


    value_map_1 := map[string]int{}
    value_map_1["A"] = 14
    value_map_1["K"] = 13
    value_map_1["Q"] = 12
    value_map_1["J"] = 11
    value_map_1["T"] = 10
    for idx := 2; idx < 10; idx++ {
        value_map_1[fmt.Sprint(idx)] = idx
    }

    value_map_2 := map[string]int{}
    value_map_2["A"] = 13
    value_map_2["K"] = 12
    value_map_2["Q"] = 11
    value_map_2["T"] = 10
    for idx := 2; idx < 10; idx++ {
        value_map_2[fmt.Sprint(idx)] = idx
    }
    value_map_2["J"] = 1

    hands_1 := []Hand{}
    hands_2 := []Hand{}

    for _, x := range lines {
        a := strings.Split(x, " ")
        hand_1 := Hand{}
        hand_2 := Hand{}
        bid, err := strconv.Atoi(a[1])
        check(err)
        hand_1.bid = bid
        hand_2.bid = bid
        occurences_1 := map[int]int{}
        occurences_2 := map[int]int{}
        for _, y := range a[0] {
            hand_1.cards = append(hand_1.cards, value_map_1[string(y)])
            hand_2.cards = append(hand_2.cards, value_map_2[string(y)])
            occurences_1[value_map_1[string(y)]] += 1
            occurences_2[value_map_2[string(y)]] += 1
        }
        tmp := 1
        for _, x := range occurences_1 {
            tmp *= x
        }
        if tmp == 1 {
            hand_1.combo = HighCard
        } else if tmp == 2 {
            hand_1.combo = OnePair
        } else if tmp == 3 {
            hand_1.combo = ThreeKind
        } else if tmp == 4 {
            if len(occurences_1) == 3 {
                hand_1.combo = TwoPair
            } else {
                hand_1.combo = FourKind
            }
        } else if tmp == 5 {
            hand_1.combo = FiveKind
        } else {
            hand_1.combo = FullHouse
        }

        if count, ok := occurences_2[1]; ok {
            delete(occurences_2, 1)
            if count == 1 {
                if len(occurences_2) == 4 {
                    hand_2.combo = OnePair
                } else if len(occurences_2) == 3 {
                    hand_2.combo = ThreeKind
                } else if len(occurences_2) == 2 {
                    res := 1
                    for _, x := range occurences_2 {
                        res *= x
                    }
                    if res == 3 {
                        hand_2.combo = FourKind
                    } else {
                        hand_2.combo = FullHouse
                    }
                } else {
                    hand_2.combo = FiveKind
                }
            } else if count == 2 {
                if len(occurences_2) == 3 {
                    hand_2.combo = ThreeKind
                } else if len(occurences_2) == 2 {
                    hand_2.combo = FourKind
                } else {
                    hand_2.combo = FiveKind
                }
            } else if count == 3 {
                if len(occurences_2) == 2 {
                    hand_2.combo = FourKind
                } else {
                    hand_2.combo = FiveKind
                }
            } else {
                hand_2.combo = FiveKind
            }
        } else {
            res := 1
            for _, x := range occurences_2 {
                res *= x
            }
            if res == 1 {
                hand_2.combo = HighCard
            } else if res == 2 {
                hand_2.combo = OnePair
            } else if res == 3 {
                hand_2.combo = ThreeKind
            } else if res == 4 {
                if len(occurences_2) == 3 {
                    hand_2.combo = TwoPair
                } else {
                    hand_2.combo = FourKind
                }
            } else if res == 5 {
                hand_2.combo = FiveKind
            } else {
                hand_2.combo = FullHouse
            }
        }
        hands_1 = append(hands_1, hand_1)
        hands_2 = append(hands_2, hand_2)
    }

    for x := 0; x < len(hands_1[0].cards); x++ {
        sort.Slice(hands_1[:], func(i, j int) bool {
            if hands_1[i].combo == hands_1[j].combo {
                for idx := 0; idx < len(hands_1[i].cards); idx++ {
                    if hands_1[i].cards[idx] != hands_1[j].cards[idx] {
                        return hands_1[i].cards[idx] > hands_1[j].cards[idx]
                    }
                }
            }
            return hands_1[i].combo < hands_1[j].combo
        })
        sort.Slice(hands_2[:], func(i, j int) bool {
            if hands_2[i].combo == hands_2[j].combo {
                for idx := 0; idx < len(hands_2[i].cards); idx++ {
                    if hands_2[i].cards[idx] != hands_2[j].cards[idx] {
                        return hands_2[i].cards[idx] > hands_2[j].cards[idx]
                    }
                }
            }
            return hands_2[i].combo < hands_2[j].combo
        })
    }
    result_1 := 0
    for x := len(hands_1)-1; x >= 0; x-- {
        bid := hands_1[len(hands_1)-x-1].bid
        result_1 += (x+1)*bid
    }
    fmt.Println(result_1)

    // result 2 experiment
    result_2 := 0
    for x := len(hands_2)-1; x >= 0; x-- {
        bid := hands_2[len(hands_2)-x-1].bid
        result_2 += (x+1)*bid
    }
    fmt.Println(result_2)
}
