package days

import (
  "fmt"
  "io"
  "os"
  "slices"
  "strconv"
  "strings"
)

type LRNode struct {
  left string;
  right string;
}

func get_loop_indices(current []string, instructions []string, nodes map[string]LRNode) []int {
  loops := make([]int, len(current))
  n := 0
  inst_len := len(instructions)
  for {
    for i := range current {
      // if final position is reached as instructions run out, this section will loop over forever
      if n % inst_len == 0 && loops[i] == 0 && current[i][2] == 'Z' {
        loops[i] = n
        if !slices.Contains(loops, 0) { return loops }
      }
      // move
      if instructions[n % inst_len] == "R" {
        current[i] = nodes[current[i]].right
      } else {
        current[i] = nodes[current[i]].left
      }
    }
    n++
  }
}

func Day08() (string, string) {
  file, _ := os.Open("input/day08.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n\n")
  instructions := strings.Split(input[0], "")
  node_strings := strings.Split(input[1], "\n")
  nodes := make(map[string]LRNode)
  for n := range node_strings {
    var name string
    var left string
    var right string
    fmt.Sscanf(node_strings[n], "%3s = (%3s, %3s)", &name, &left, &right)
    nodes[name] = LRNode{left, right}
  }

  // part 1
  current := "AAA"
  p1 := 0
  inst_len := len(instructions)
  for current != "ZZZ" {
    if instructions[p1 % inst_len] == "R" {
      current = nodes[current].right
    } else {
      current = nodes[current].left
    }
    p1++
  }

  // part 2
  var p2_starts []string
  for k := range nodes { if k[2] == 'A' { p2_starts = append(p2_starts, k) } }
  p2 := get_loop_indices(p2_starts, instructions, nodes)

  return strconv.Itoa(p1), strconv.Itoa(lcm(p2...))
}
