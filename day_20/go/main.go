package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Pulse int
const(
    HIGH Pulse = iota
    LOW
    NONE
)

type State int
const(
    ON State = iota
    OFF
)

type FlipFlop struct {
    state State
}

func (f *FlipFlop) receive_pulse(pulse Pulse) (bool, Pulse) {
    if pulse == LOW {
        if f.state == ON {
            f.state = OFF
            return true, LOW
        } else {
            f.state = ON
            return true, HIGH
        }
    }
    return false, NONE
}

type Conjunction struct {
    input_pulses map[string]Pulse
}

func (c *Conjunction) next_pulse() Pulse {
    for _, x := range c.input_pulses {
        if x != HIGH {
            return HIGH
        }
    }
    return LOW
}

type Instruction struct {
    name string
    flip_flop *FlipFlop
    conjunction *Conjunction
    receivers []string
    pulse Pulse
}

var INSTRUCTION_MAP = map[string]Instruction{}

func main() {
    input_file := os.Args[1]
    abs_path, err := filepath.Abs(input_file)
    check(err)
    data, err := os.ReadFile(abs_path)
    data_string := strings.Trim(string(data), "\n")
    lines := strings.Split(data_string, "\n")
    var key string
    for _, x := range lines {
        new_instruction := Instruction{pulse: LOW}
        split := strings.Split(x, " -> ")
        left, right := split[0], split[1]
        receivers := strings.Split(right, ", ")
        new_instruction.receivers = receivers
        if left[0] == '&' {
            conjunction := Conjunction{input_pulses: map[string]Pulse{}}
            new_instruction.conjunction = &conjunction
            key = left[1:]
        } else if left[0] == '%' {
            new_instruction.flip_flop = &FlipFlop{state: OFF}
            key = left[1:]
        } else {
            // broadcast
            key = left
        }
        new_instruction.name = key
        INSTRUCTION_MAP[key] = new_instruction
    }

    // initialize low input pulse for each conjuction input

    target := "rx"
    source := Instruction{name: "not found"}
    for key, x := range INSTRUCTION_MAP {
        for _, y := range x.receivers {
            if y == target {
                source = x
            }
            receiver, _ := INSTRUCTION_MAP[y]
            if receiver.conjunction != nil {
                INSTRUCTION_MAP[receiver.name].conjunction.input_pulses[key] = LOW
            }
        }
        // fmt.Printf("%s: <%s, %+v, %+v>\n", key, x.receivers, x.flip_flop, x.conjunction)
    }
    fmt.Println(source.conjunction.input_pulses)
    targets := []string{}
    for x := range source.conjunction.input_pulses {
        targets = append(targets, x)
    }

    check := map[string][]int{}
    low, high := 0, 0
    for i := range [100000]int{} {
        low++
        start, _ := INSTRUCTION_MAP["broadcaster"]
        queue := []Instruction{start}
        for len(queue) > 0 {
            current := queue[0]
            queue = queue[1:]
            if slices.Contains(targets, current.name) && current.pulse == HIGH {
                check[current.name] = append(check[current.name], i)
            }
            for _, receiver := range current.receivers {
                if current.pulse == LOW {
                    low++
                } else {
                    high++
                }
                x, _ := INSTRUCTION_MAP[receiver]
                if x.flip_flop != nil {
                    ok, next_pulse := x.flip_flop.receive_pulse(current.pulse)
                    if ok {
                        x.pulse = next_pulse
                        queue = append(queue, x)
                    }
                } else if x.conjunction != nil {
                    x.conjunction.input_pulses[current.name] = current.pulse
                    x.pulse = x.conjunction.next_pulse()
                    queue = append(queue, x)
                }
            }
        }
    }

    nums := []int{}
    for key, x := range check {
        fmt.Println(key, " > ", x[len(x)-1]-x[len(x)-2])
        nums = append(nums, x[len(x)-1]-x[len(x)-2])
    }

    fmt.Println(low, high)
    fmt.Println("Result 1: ", low*high)
    fmt.Println("Result 2: ", LCM(nums[0], nums[1], nums[2:]...))
}

func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}
