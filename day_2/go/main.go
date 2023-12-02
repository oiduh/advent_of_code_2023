package main

import (
    "fmt"
	"os"
	"path/filepath"
	"strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Game struct {
    id int
    red int
    green int
    blue int
}

func main() {
    path := os.Args[1]
    absPath, err := filepath.Abs(path)
    check(err)
    dat, err := os.ReadFile(absPath)
    dat_string := strings.Trim(string(dat), "\n")
    lines := strings.Split(dat_string, "\n")
    games := []Game{}
    for _, x := range lines {
        game_colors_split := strings.Split(x, ": ")
        game_id := strings.Split(game_colors_split[0], " ")[1]
        id_int, err := strconv.Atoi(game_id)
        check(err)
        colors := strings.Split(game_colors_split[1], "; ")
        new_game := Game{id_int, 0, 0, 0}
        for _, y := range colors {
            color_groups_split := strings.Split(y, ", ")
            for _, z := range color_groups_split {
                color_amount_split := strings.Split(z, " ")
                color := color_amount_split[1]
                amount, err := strconv.Atoi(color_amount_split[0])
                check(err)

                switch color {
                case "red":
                    new_game.red = max(new_game.red, amount)
                case "green":
                    new_game.green = max(new_game.green, amount)
                case "blue":
                    new_game.blue = max(new_game.blue, amount)
                }
            }
        }
        games = append(games, new_game)
    }
    target := Game{-1, 12, 13, 14}
    result_1 := 0
    for _, game := range games {
        if game.red <= target.red && game.green <= target.green && game.blue <= target.blue {
            result_1 += game.id
        }
    }
    fmt.Println(result_1)

    result_2 := 0
    for _, game := range games {
        result_2 += game.red * game.green * game.blue
    }
    fmt.Println(result_2)
}
