package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

type TravelPoint struct {
  row int
  col int
  direction int // 0 right 1 down 2 left 3 up
}

func next_mirror_points(current TravelPoint, currrenttype string, n_rows int, n_cols int) []TravelPoint {
  var points []TravelPoint
  switch currrenttype {
  case ".":
    switch current.direction {
    case 0:
      if current.col + 1 < n_cols { points = []TravelPoint{{current.row, current.col + 1, current.direction}} }
    case 1:
      if current.row + 1 < n_rows { points = []TravelPoint{{current.row + 1, current.col, current.direction}} }
    case 2:
      if current.col - 1 >= 0 { points = []TravelPoint{{current.row, current.col - 1, current.direction}} }
    case 3:
      if current.row - 1 >= 0 { points = []TravelPoint{{current.row - 1, current.col, current.direction}} }
    }
  case "/":
    switch current.direction {
    case 0:
      if current.row - 1 >= 0 { points = []TravelPoint{{current.row - 1, current.col, 3}} }
    case 1:
      if current.col - 1 >= 0 { points = []TravelPoint{{current.row, current.col - 1, 2}} }
    case 2:
      if current.row + 1 < n_rows { points = []TravelPoint{{current.row + 1, current.col, 1}} }
    case 3:
      if current.col + 1 < n_cols { points = []TravelPoint{{current.row, current.col + 1, 0}} }
    }
  case "\\":
    switch current.direction {
    case 0:
      if current.row + 1 < n_rows { points = []TravelPoint{{current.row + 1, current.col, 1}} }
    case 1:
      if current.col + 1 < n_cols { points = []TravelPoint{{current.row, current.col + 1, 0}} }
    case 2:
      if current.row - 1 >= 0 { points = []TravelPoint{{current.row - 1, current.col, 3}} }
    case 3:
      if current.col - 1 >= 0 { points = []TravelPoint{{current.row, current.col - 1, 2}} }
    }
  case "-":
    switch current.direction {
    case 0:
      if current.col + 1 < n_cols { points = []TravelPoint{{current.row, current.col + 1, current.direction}} }
    case 1:
      if current.col + 1 < n_cols { points = append(points, TravelPoint{current.row, current.col + 1, 0}) }
      if current.col - 1 >= 0 { points = append(points, TravelPoint{current.row, current.col - 1, 2}) }
    case 2:
      if current.col - 1 >= 0 { points = []TravelPoint{{current.row, current.col - 1, current.direction}} }
    case 3:
      if current.col + 1 < n_cols { points = append(points, TravelPoint{current.row, current.col + 1, 0}) }
      if current.col - 1 >= 0 { points = append(points, TravelPoint{current.row, current.col - 1, 2}) }
    }
  case "|":
    switch current.direction {
    case 0:
      if current.row + 1 < n_rows { points = append(points, TravelPoint{current.row + 1, current.col, 1}) }
      if current.row - 1 >= 0 { points = append(points, TravelPoint{current.row - 1, current.col, 3}) }
    case 1:
      if current.row + 1 < n_rows { points = []TravelPoint{{current.row + 1, current.col, current.direction}} }
    case 2:
      if current.row + 1 < n_rows { points = append(points, TravelPoint{current.row + 1, current.col, 1}) }
      if current.row - 1 >= 0 { points = append(points, TravelPoint{current.row - 1, current.col, 3}) }
    case 3:
      if current.row - 1 >= 0 { points = []TravelPoint{{current.row - 1, current.col, current.direction}} }
    }
  }
  return points
}

func traverse_mirror_grid(grid [][]string, start TravelPoint) int {
  n_rows := len(grid)
  n_cols := len(grid[0])

  visited := make([][][]bool, n_rows)
  for i := range visited {
    visited[i] = make([][]bool, n_cols)
    for j := range visited[i] { visited[i][j] = make([]bool, 4) }
  }
  
  queue := make([]TravelPoint, n_rows*n_cols*4) // all possible points * possible directions
  queue_start := 0
  queue_end := 0

  queue[queue_end] = start
  queue_end++
  visited[start.row][start.col][start.direction] = true

  for queue_start != queue_end {
    current := queue[queue_start]
    queue_start++
    next := next_mirror_points(current, grid[current.row][current.col], n_rows, n_cols)
    for n := range next {
      if !visited[next[n].row][next[n].col][next[n].direction] {
        visited[next[n].row][next[n].col][next[n].direction] = true
        queue[queue_end] = next[n]
        queue_end++
      }
    }
  }
  total := 0
  for i := range visited {
    for j := range visited[i] {
      if any_true(visited[i][j]) { total++ }
    }
  }
  return total
}

func find_best_mirror_start(grid [][]string) int {
  best := 0
  n_rows := len(grid)
  n_cols := len(grid[0])
  for row := range grid {
    if row == 0 {
      top_attempt := traverse_mirror_grid(grid, TravelPoint{0, 0, 1})
      if top_attempt > best { best = top_attempt }
    } else if row == n_rows - 1 {
      bottom_attempt := traverse_mirror_grid(grid, TravelPoint{n_rows - 1, n_cols - 1, 3})
      if bottom_attempt > best { best = bottom_attempt }
    }
    left_attempt := traverse_mirror_grid(grid, TravelPoint{row, 0, 0})
    if left_attempt > best { best = left_attempt }
    right_attempt := traverse_mirror_grid(grid, TravelPoint{row, n_cols - 1, 2})
    if right_attempt > best { best = right_attempt }
  }
  for col := 1; col < n_cols - 1; col++ {
    top_attempt := traverse_mirror_grid(grid, TravelPoint{0, col, 1})
    if top_attempt > best { best = top_attempt }
    bottom_attempt := traverse_mirror_grid(grid, TravelPoint{n_rows - 1, col, 3})
    if bottom_attempt > best { best = bottom_attempt }
  }
  return best
}

func Day16() (string, string) {
  file, _ := os.Open("input/day16.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  grid := make([][]string, len(input))
  for i := range input {
    grid[i] = make([]string, len(input[i]))
    chars := strings.Split(input[i], "")
    for c := range chars {
      grid[i][c] = chars[c]
    }
  }

  p1 := traverse_mirror_grid(grid, TravelPoint{0, 0, 0})
  p2 := find_best_mirror_start(grid)

  return strconv.Itoa(p1), strconv.Itoa(p2)
}
