package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Operator int
const(
    GREATER Operator = iota
    LESS
    NONE
)

type Workflow struct {
    variable string
    operator Operator
    comparison int
    next string
    is_end bool
}

type Rating struct {
    variable string
    rating int
}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    split := strings.Split(data_string, "\n\n")
    workflow_strings := split[0]
    ratings_strings := split[1]
    // fmt.Println(workflow_strings)
    // fmt.Println(ratings_strings)

    workflow_map := map[string][]Workflow{}

    for _, a := range strings.Split(workflow_strings, "\n") {
        test := strings.Split(a, "\n")
        test = strings.Split(test[0], "{")
        key := test[0]
        rest := test[1]
        tmp := strings.Split(rest[:len(rest)-1], ",")
        // fmt.Println(key)
        // fmt.Println(tmp)
        workflow_map[key] = []Workflow{}
        for _, x := range tmp {
            new_workflow := Workflow{}
            if strings.Contains(x, ">") {
                y := strings.Split(x, ">")
                l, r := y[0], y[1]
                z := strings.Split(r, ":")
                num, next := z[0], z[1]
                num_int, err := strconv.Atoi(num)
                check(err)
                // fmt.Printf("var: %s, operator: '>', comparison: %s, next: %s\n", l, num, next)
                new_workflow.is_end = false
                new_workflow.variable = l
                new_workflow.operator = GREATER
                new_workflow.comparison = num_int
                new_workflow.next = next
            } else if strings.Contains(x, "<") {
                y := strings.Split(x, "<")
                l, r := y[0], y[1]
                z := strings.Split(r, ":")
                num, next := z[0], z[1]
                num_int, err := strconv.Atoi(num)
                check(err)
                // fmt.Printf("var: %s, operator: '<', comparison: %s, next: %s\n", l, num, next)
                new_workflow.is_end = false
                new_workflow.variable = l
                new_workflow.operator = LESS
                new_workflow.comparison = num_int
                new_workflow.next = next
            } else {
                // fmt.Println(x)
                new_workflow.is_end = true
                new_workflow.variable = "NONE"
                new_workflow.operator = NONE
                new_workflow.comparison = -1
                new_workflow.next = x
            }
            workflow_map[key] = append(workflow_map[key], new_workflow)
        }
    }
    // fmt.Println(workflow_map)

    ratings := []map[string]int{}
    for _, x := range strings.Split(ratings_strings, "\n") {
        tmp := x[1:len(x)-1]
        a := strings.Split(tmp, ",")
        new_map := map[string]int{}
        for _, b := range a {
            split := strings.Split(b, "=")
            variable, rating := split[0], split[1]
            rating_int, err := strconv.Atoi(rating)
            check(err)
            new_map[variable] = rating_int
        }
        ratings = append(ratings, new_map)
    }
    // fmt.Println(ratings)


    result := 0
    for _, x := range ratings {
        workflow, ok := workflow_map["in"]
        if !ok {
            panic("error")
        }
        current := "START"
        for !(current == "A" || current == "R") {
            // fmt.Println(workflow)
            for _, tmp := range workflow {
                target_variable := tmp.variable
                // fmt.Println(tmp.comparison, x[target_variable])
                if tmp.operator == GREATER && tmp.comparison < x[target_variable] {
                    current = tmp.next
                    workflow = workflow_map[current]
                    break
                } else if tmp.operator == LESS && tmp.comparison > x[target_variable] {
                    current = tmp.next
                    workflow = workflow_map[current]
                    break
                } else if tmp.is_end {
                    current = tmp.next
                    if !(current == "A" || current == "R") {
                        workflow = workflow_map[current]
                    }
                    break
                }
            }
        }
        // fmt.Println(current)
        if current == "A" {
            for _, b := range x {
                result += b
            }
        }
    }
    fmt.Println(result)


    root := Tree{instruction: "in", value: 1}
    traverse(&root, workflow_map["in"], workflow_map)

    restrictions := map[string][]Restriction{}
    restrictions["x"] = []Restriction{}
    restrictions["m"] = []Restriction{}
    restrictions["a"] = []Restriction{}
    restrictions["s"] = []Restriction{}
    print_tree(&root, 0, &restrictions, &map[string][]Restriction{})

    res := 1
    for a, b := range restrictions {
        minimum, maximum := math.MaxInt, 0
        // fmt.Println(a)
        for _, c := range b {
            // fmt.Printf("    %+v\n", c)
            if c.operator == ">" && c.num < minimum {
                minimum = c.num
            } else if c.operator == ">=" && c.num <= minimum {
                minimum = c.num
            } else if c.operator == "<" && c.num > maximum {
                maximum = c.num
            } else if c.operator == "<=" && c.num >= maximum {
                maximum = c.num
            }
        }
        fmt.Printf("%s = (%d, %d)\n", a, minimum, maximum)
        res *= (maximum-minimum)
    }
    fmt.Println(res)
}

type Restriction struct {
    variable string
    operator string
    num int
}

func print_tree(root *Tree, indent int, restrictions *map[string][]Restriction, tmp *map[string][]Restriction) {
    if root.instruction == "A" || root.instruction == "R" {
        // for i := 0; i < indent; i++ {
        //     fmt.Printf(" ")
        // }
        // fmt.Print(root)
        // fmt.Println()
        if root.instruction == "A" {
            for key, value := range *tmp {
                (*restrictions)[key] = append((*restrictions)[key], value...)
            }
        }
        return
    }
    // for i := 0; i < indent; i++ {
    //     fmt.Printf(" ")
    // }
    // fmt.Print(root)
    // fmt.Println()
    indent += 2
    for _, child := range root.children {
        (*tmp)[child.variable] = append((*tmp)[child.variable], child.restrictions...)
        print_tree(child, indent, restrictions, tmp)
    }
}

func traverse(current *Tree, workflows []Workflow, workflow_map map[string][]Workflow) {
    var previous Workflow
    restrictions := []Restriction{}
    for _, workflow := range workflows {
        var sub_tree Tree
        if workflow.operator != NONE {
            if workflow.operator == GREATER {
                sub_tree = Tree{
                    value: workflow.comparison,
                    instruction: workflow.next,
                    variable: workflow.variable,
                    operator: ">",
                    restrictions: append(restrictions, Restriction{
                        num: workflow.comparison, operator: ">", variable: workflow.variable,
                    }),
                }
                restrictions = append(restrictions, Restriction{
                    num: workflow.comparison, operator: "<=", variable: workflow.variable,
                })
            } else {
                sub_tree = Tree{
                    value: workflow.comparison,
                    instruction: workflow.next,
                    variable: workflow.variable,
                    operator: "<",
                    restrictions: append(restrictions, Restriction{
                        num: workflow.comparison, operator: "<", variable: workflow.variable,
                    }),
                }
                restrictions = append(restrictions, Restriction{
                    num: workflow.comparison, operator: ">=", variable: workflow.variable,
                })
            }
            current.children = append(current.children, &sub_tree)
        } else {
            if previous.operator == GREATER {
                sub_tree = Tree{
                    value: previous.comparison,
                    instruction: workflow.next,
                    variable: previous.variable,
                    operator: "<=",
                    restrictions: restrictions,
                }
            } else {
                sub_tree = Tree{
                    value: previous.comparison,
                    instruction: workflow.next,
                    variable: previous.variable,
                    operator: ">=",
                    restrictions: restrictions,
                }
            }
            current.children = append(current.children, &sub_tree)
        }
        previous = workflow
    }
    for _, child := range current.children {
        traverse(child, workflow_map[child.instruction], workflow_map)
    }
}

type Tree struct {
    instruction string
    variable string
    operator string
    value int
    children []*Tree
    restrictions []Restriction
}
