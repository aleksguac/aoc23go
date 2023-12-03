package days

import (
  "io"
  "os"
  "slices"
  "strconv"
  "strings"
)

func get_whole_value(grid *[][]int, i int, j int) (int, []int) {
  if (*grid)[i][j] < 0 || (j > 0 && (*grid)[i][j-1] >= 0) { return 0, []int{} }
  val := 0
  symbol := 0
  j_current := j
  var gears []int
  for {
    if j_current >= len((*grid)[i]) { break }
    if (*grid)[i][j_current] < 0 { break }
  
    // get value
    val = val * 10 + (*grid)[i][j_current]
  
    // search for neighbouring symbols
    for i_search := i-1; i_search <= i+1; i_search++ {
      for j_search := j_current-1; j_search <= j_current+1; j_search++ {
        if i_search >= 0 && i_search < len(*grid) && j_search >= 0 && j_search < len((*grid)[i_search]) &&
           (*grid)[i_search][j_search] < -1 {
          symbol = (*grid)[i_search][j_search]
          // if symbol is a gear, add it to neighbour-gear list
          if symbol == -2 {
            gearloc := i_search*len((*grid)[i_search]) + j_search
            if !slices.Contains(gears, gearloc) {
              gears = append(gears, gearloc)
            }
          }
        }
      }
    }
    j_current++
  }
  if symbol < -1 {
    return val, gears
  } else {
    // value returned only if symbol was adjacent
    return 0, []int{}
  }
}

func Day03() (string, string) {
  file, _ := os.Open("input/day03.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  var grid [][]int
  for i := range input {
    grid = append(grid, []int{})
    chars := strings.Split(input[i], "")
    for c := range chars {
      val, err := strconv.Atoi(chars[c])
      if err == nil {
        grid[i] = append(grid[i], val)
      } else if chars[c] == "." {
        grid[i] = append(grid[i], -1)
      } else if chars[c] == "*" {
        grid[i] = append(grid[i], -2)
      } else {
        grid[i] = append(grid[i], -3)
      }
    }
  }
  p1 := 0
  gears := make(map[int][]int)
  for i := range grid {
    for j := range grid[i] {
      // search for values then note neighbour-gears
      val, ij_gears := get_whole_value(&grid, i, j)
      p1 += val
      for g := range ij_gears {
        _, ok := gears[ij_gears[g]]
        if ok {
          gears[ij_gears[g]] = append(gears[ij_gears[g]], val)
        } else {
          gears[ij_gears[g]] = []int{val}
        }
      }
    }
  }
  p2 := 0
  for _, v := range gears {
    if len(v) == 2 {
      p2 += v[0]*v[1]
    }
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
