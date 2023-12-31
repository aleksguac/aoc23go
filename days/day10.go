package days

import (
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

type MetalNode struct {
  nodetype MetalBend;
  isloop bool;
  distance int;
  neighbours []int;
}

func get_start_type(grid *[]MetalNode, n int, n_cols int) {
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
    (*grid)[current].isloop = true
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

func get_insides(grid *[]MetalNode, n_cols int) int {
  leftcount := 0
  total := 0
  for i := range *grid {
    if i % n_cols == 0 { leftcount = 0 }
    if (*grid)[i].isloop {
      if (*grid)[i].nodetype == Vertical || (*grid)[i].nodetype == SouthEast || (*grid)[i].nodetype == SouthWest {
        leftcount++
      }
    } else if leftcount % 2 == 1 {
      total++
    }
  }

  return total
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
      n := XYtoN(j, i, len(chars))
      var neighbours []int
      switch chars[j] {
      case "|":
        nodetype = Vertical
        neighbours = []int{n-n_cols, n+n_cols}
      case "-":
        nodetype = Horizontal
        neighbours = []int{n-1, n+1}
      case "L":
        nodetype = NorthEast
        neighbours = []int{n-n_cols, n+1}
      case "J":
        nodetype = NorthWest
        neighbours = []int{n-n_cols, n-1}
      case "7":
        nodetype = SouthWest
        neighbours = []int{n+n_cols, n-1}
      case "F":
        nodetype = SouthEast
        neighbours = []int{n+n_cols, n+1}
      case ".":
        nodetype = NoPipe
      case "S":
        start = n
        nodetype = NoPipe
      default:
        continue
      }
      grid[n] = MetalNode{nodetype, false, 0, neighbours}
    }
  }

  get_start_type(&grid, start, n_cols)
  p1 := get_metal_loop(&grid, start, n_cols)
  p2 := get_insides(&grid, n_cols)

  return strconv.Itoa(p1), strconv.Itoa(p2)
}
