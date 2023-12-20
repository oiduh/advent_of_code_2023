package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Position struct {
    row, col int
}

type Vertex struct {
    key Position
    edges map[*Vertex]int
}

type Graph struct {
    vertices map[Position]*Vertex
}

type Direction int
const(
    NORTH Direction = iota
    EAST
    SOUTH
    WEST
)

func within_limit_row(history []Position) bool {
    fmt.Println(history)
    if len(history) >= 3 {
        last := history[len(history)-1]
        last_row := last.row
        // last_col := last.col
        // row check
        for _, position := range history[len(history)-3:] {
            if position.row != last_row {
                return true
            }
        }
        return false
    }
    return true
}

func within_limit_col(history []Position) bool {
    fmt.Println(history)
    if len(history) >= 3 {
        last := history[len(history)-1]
        last_col := last.col
        for _, position := range history[len(history)-3:] {
            if position.col != last_col {
                return true
            }
        }
        return false
    }
    return true
}

func (g *Graph) Dijkstra(start_pos Position) (distances map[Position]int, history map[Position][]Position, directions map[Position][]Direction, err error) {
    _, ok := g.vertices[start_pos]
    if !ok {
        return nil, nil, nil, fmt.Errorf("start vertex %d not found", start_pos)
    }

    history = make(map[Position][]Position)
    distances = make(map[Position]int)
    directions = make(map[Position][]Direction)
    for key := range g.vertices {
        distances[key] = math.MaxInt
        history[key] = []Position{}
        directions[key] = []Direction{}
    }
    distances[start_pos] = 0

    var vertices []*Vertex
    for _, vertex := range g.vertices {
        vertices = append(vertices, vertex)
    }

    for len(vertices) != 0 {
        sort.SliceStable(vertices, func(i, j int) bool {
            return distances[vertices[i].key] < distances[vertices[j].key]
        })

        vertex := vertices[0]
        vertices = vertices[1:]

        for adjacent, cost := range vertex.edges {
            alt := distances[vertex.key] + cost 
            if alt < distances[adjacent.key] {
                history[adjacent.key] = append(history[vertex.key], vertex.key)
                if len(history[adjacent.key]) >= 2 {
                    current := history[adjacent.key][len(history[adjacent.key])-1]
                    previous := history[adjacent.key][len(history[adjacent.key])-2]
                    direction := SOUTH
                    if previous.row - current.row == 1 {
                        direction = NORTH
                    } else if previous.row - current.row == -1 {
                        direction = SOUTH
                    } else if previous.col - current.col == 1 {
                        direction = WEST
                    } else {
                        direction = EAST
                    }
                    directions[adjacent.key] = append(directions[vertex.key], direction)
                }
                distances[adjacent.key] = alt
            } 
        }
    }
    return distances, history, directions, nil
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")
    fmt.Println(lines[0])

    matrix_num := [][]int{}
    for _, line := range lines {
        row_num := []int{}
        for _, num_str := range line {
            num_int, err := strconv.Atoi(string(num_str))
            check(err)
            row_num = append(row_num, num_int)
        }
        matrix_num = append(matrix_num, row_num)
    }

    graph := Graph{vertices: map[Position]*Vertex{}}
    for i, row := range matrix_num {
        for j, current_num := range row {
            current_pos := Position{i,j}
            var current_vertex *Vertex
            current_vertex, ok := graph.vertices[current_pos]
            if !ok {
                current_vertex = &Vertex{key: current_pos, edges: map[*Vertex]int{}}
                graph.vertices[current_pos] = current_vertex
            }
            if i-1 >= 0 {
                above_pos := Position{i-1,j}
                above_num := matrix_num[i-1][j]
                above_vertex, ok := graph.vertices[above_pos]
                if !ok {
                    new_vertex := Vertex{key: above_pos, edges: map[*Vertex]int{}}
                    above_vertex = &new_vertex
                    graph.vertices[above_pos] = above_vertex
                }
                current_vertex.edges[above_vertex] = above_num
                above_vertex.edges[current_vertex] = current_num
            }
            if i+1 < len(matrix_num) {
                below_pos := Position{i+1,j}
                below_num := matrix_num[i+1][j]
                below_vertex, ok := graph.vertices[below_pos]
                if !ok {
                    new_vertex := Vertex{key: below_pos, edges: map[*Vertex]int{}}
                    below_vertex = &new_vertex
                    graph.vertices[below_pos] = below_vertex
                }
                current_vertex.edges[below_vertex] = below_num
                below_vertex.edges[current_vertex] = current_num
            }
            if j-1 >= 0 {
                left_pos := Position{i,j-1}
                left_num := matrix_num[i][j-1]
                left_vertex, ok := graph.vertices[left_pos]
                if !ok {
                    new_vertex := Vertex{key: left_pos, edges: map[*Vertex]int{}}
                    left_vertex = &new_vertex
                    graph.vertices[left_pos] = left_vertex
                }
                current_vertex.edges[left_vertex] = left_num
                left_vertex.edges[current_vertex] = current_num
            }
            if j+1 < len(matrix_num[0]) {
                right_pos := Position{i,j+1}
                right_num := matrix_num[i][j+1]
                right_vertex, ok := graph.vertices[right_pos]
                if !ok {
                    new_vertex := Vertex{key: right_pos, edges: map[*Vertex]int{}}
                    right_vertex = &new_vertex
                    graph.vertices[right_pos] = right_vertex
                }
                current_vertex.edges[right_vertex] = right_num
                right_vertex.edges[current_vertex] = current_num
            }
        }
    }

    result, history, directions, err := graph.Dijkstra(Position{0,0})
    check(err)
    fmt.Println("RESULT")
    fmt.Println(result[Position{len(matrix_num)-1,len(matrix_num[0])-1}])
    fmt.Println(history[Position{len(matrix_num)-1,len(matrix_num[0])-1}])
    fmt.Println(directions[Position{len(matrix_num)-1,len(matrix_num[0])-1}])
    fmt.Println()

    for _, x := range matrix_num {
        fmt.Println(x)
    }
    fmt.Println()
    result, history, directions, err = graph.Dijkstra(Position{len(matrix_num)-1,len(matrix_num[0])-1})
    check(err)
    fmt.Println("RESULT")
    fmt.Println(result[Position{0,0}])
    fmt.Println(history[Position{0,0}])
    fmt.Println(directions[Position{0,0}])
    fmt.Println()

    for _, x := range matrix_num {
        fmt.Println(x)
    }
    fmt.Println()
}
