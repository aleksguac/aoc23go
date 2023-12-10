package days

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MetalBend int
const (
  NoPipe MetalBend = iota
  Horizontal
  Vertical
  NorthWest
  NorthEast
  SouthWest
  SouthEast
)

type LoopState int
const (
  UnknownState LoopState = iota
  IsLoop
  OutsideLoop
  InsideLoop
)

type MetalNode struct {
  x int;
  y int;
  nodetype MetalBend;
  where LoopState;
  distance int;
  neighbours []int;
}

func get_start_type_and_neighbours(grid *[]MetalNode, n int, n_cols int) {
  dirs := make([]bool, 4) // NESW
  if n / n_cols < len(*grid) / n_cols - 1 && ((*grid)[n+n_cols].nodetype == Vertical ||
                                              (*grid)[n+n_cols].nodetype == NorthEast ||
                                              (*grid)[n+n_cols].nodetype == NorthWest) {
    (*grid)[n].neighbours = append((*grid)[n].neighbours, n+n_cols)
    dirs[2] = true
  }
  if n / n_cols > 0 && ((*grid)[n-n_cols].nodetype == Vertical ||
                        (*grid)[n-n_cols].nodetype == SouthEast ||
                        (*grid)[n-n_cols].nodetype == SouthWest) {
    (*grid)[n].neighbours = append((*grid)[n].neighbours, n-n_cols)
    dirs[0] = true
  }
  if n % n_cols > 0 && ((*grid)[n-1].nodetype == Horizontal ||
                        (*grid)[n-1].nodetype == SouthEast ||
                        (*grid)[n-1].nodetype == NorthEast) {
    (*grid)[n].neighbours = append((*grid)[n].neighbours, n-1)
    dirs[3] = true
  }
  if n % n_cols < n_cols-1 && ((*grid)[n+1].nodetype == Horizontal ||
                               (*grid)[n+1].nodetype == SouthWest ||
                               (*grid)[n+1].nodetype == NorthWest) {
    (*grid)[n].neighbours = append((*grid)[n].neighbours, n+1)
    dirs[1] = true
  }

  if dirs[0] {
    if dirs[1] {
      (*grid)[n].nodetype = NorthEast
    } else if dirs[2] {
      (*grid)[n].nodetype = Vertical
    } else if dirs[3] {
      (*grid)[n].nodetype = NorthWest
    }
  } else if dirs[1] {
    if dirs[2] {
      (*grid)[n].nodetype = SouthEast
    } else if dirs[3] {
      (*grid)[n].nodetype = Horizontal
    }
  } else if dirs[2] {
    if dirs[3] {
      (*grid)[n].nodetype = SouthWest
    }
  }

  for i := range *grid {
    switch (*grid)[i].nodetype {
    case Horizontal:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i-1, i+1}...)
    case Vertical:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i-n_cols, i+n_cols}...)
    case NorthEast:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i+1, i-n_cols}...)
    case NorthWest:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i-1, i-n_cols}...)
    case SouthEast:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i+1, i+n_cols}...)
    case SouthWest:
      (*grid)[i].neighbours = append((*grid)[i].neighbours, []int{i-1, i+n_cols}...)
    }
  }
}

func get_metal_loop(grid *[]MetalNode, start int, n_cols int) int {
  // BFS from the start node
  visited := make([]bool, len(*grid))
  queue := make([]int, len(*grid))
  queue_start := 0
  queue_end := 0

  queue[queue_end] = start
  visited[start] = true
  queue_end++

  for (queue_start != queue_end) {
    current := queue[queue_start]
    queue_start++
    (*grid)[current].where = IsLoop
    for i := range (*grid)[current].neighbours {
      n_neigh := (*grid)[current].neighbours[i]
      if !visited[n_neigh] {
        visited[n_neigh] = true
        queue[queue_end] = n_neigh
        queue_end++
        (*grid)[n_neigh].distance = (*grid)[current].distance + 1
      }
    }
  }
  return (*grid)[queue[queue_end-1]].distance
}

func get_inside_outside(grid *[]MetalNode, n_cols int) int {
  leftcount := 0
  for i := range *grid {
    if i % n_cols == 0 { leftcount = 0 }
    if (*grid)[i].where == UnknownState {
      if leftcount % 2 == 0 {
        (*grid)[i].where = OutsideLoop
      } else {
        (*grid)[i].where = InsideLoop
      }
    } else if (*grid)[i].where == IsLoop {
      if (*grid)[i].nodetype == Vertical || (*grid)[i].nodetype == SouthEast || (*grid)[i].nodetype == SouthWest {
        leftcount++
      }
    }
  }

  total_in := 0
  for i := range *grid {
    if (*grid)[i].where == InsideLoop {
      total_in++
    }
  }
  return total_in
}

func draw_metal_loop(grid *[]MetalNode, n_cols int) {
  for i := range *grid {
    if i % n_cols == 0 { fmt.Println() }
    if (*grid)[i].where == IsLoop {
      switch (*grid)[i].nodetype {
      case Horizontal:
        fmt.Print("─")
      case Vertical:
        fmt.Print("│")
      case NorthEast:
        fmt.Print("└")
      case NorthWest:
        fmt.Print("┘")
      case SouthEast:
        fmt.Print("┌")
      case SouthWest:
        fmt.Print("┐")
      }
    } else if (*grid)[i].where == OutsideLoop {
      fmt.Print(" ")
    } else if (*grid)[i].where == InsideLoop {
      fmt.Print(".")
    } else {
      fmt.Print("o")
    }
  }
  fmt.Println()
}

func Day10() (string, string) {
  file, _ := os.Open("input/day10.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")

  n_cols := len(input[0])
  grid := make([]MetalNode, len(input)*n_cols)
  start := 0
  for i := range input {
    chars := strings.Split(input[i], "")
    for j := range chars {
      var nodetype MetalBend
      switch chars[j] {
      case "|":
        nodetype = Vertical
      case "-":
        nodetype = Horizontal
      case "L":
        nodetype = NorthEast
      case "J":
        nodetype = NorthWest
      case "7":
        nodetype = SouthWest
      case "F":
        nodetype = SouthEast
      case ".":
        nodetype = NoPipe
      case "S":
        start = XYtoN(j, i, len(chars))
        nodetype = NoPipe
      default:
        continue
      }
      grid[XYtoN(j, i, len(chars))] = MetalNode{j, i, nodetype, UnknownState, 0, []int{}}
    }
  }

  get_start_type_and_neighbours(&grid, start, n_cols)
  p1 := get_metal_loop(&grid, start, n_cols)
  p2 := get_inside_outside(&grid, n_cols)

  return strconv.Itoa(p1), strconv.Itoa(p2)
}
