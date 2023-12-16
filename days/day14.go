package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

type RockType int
const (
  NothingRock RockType = iota
  CubeRock
  RoundRock
)

func predict_nth_iter(n int, loop int, start int, loop_vals []int) int {
  return loop_vals[(n-start-1) % loop]
}

func north_load(rocks [][]RockType) int {
  total := 0
  for i := range rocks {
    for j := range rocks[i] {
      if rocks[i][j] == RoundRock {
        total += len(rocks) - i
      }
    }
  }
  return total
}

func spin_cycle(rocks [][]RockType, iters int) []int {
  results := make([]int, iters)
  for n := 0; n < iters; n++ {
    changed := true
    for changed {
      changed = false
      for i := len(rocks) - 1; i >= 0; i-- {
        for j := range rocks[i] {
          if rocks[i][j] == RoundRock && i > 0 && rocks[i-1][j] == NothingRock {
            rocks[i][j] = NothingRock
            rocks[i-1][j] = RoundRock
            changed = true
          }
        }
      }
    }
    changed = true
    for changed {
      changed = false
      for j := len(rocks[0]) - 1; j >= 0; j-- {
        for i := range rocks {
          if rocks[i][j] == RoundRock && j > 0 && rocks[i][j-1] == NothingRock {
            rocks[i][j] = NothingRock
            rocks[i][j-1] = RoundRock
            changed = true
          }
        }
      }
    }
    changed = true
    for changed {
      changed = false
      for i := range rocks {
        for j := range rocks[i] {
          if rocks[i][j] == RoundRock && i < len(rocks) - 1 && rocks[i+1][j] == NothingRock {
            rocks[i][j] = NothingRock
            rocks[i+1][j] = RoundRock
            changed = true
          }
        }
      }
    }
    changed = true
    for changed {
      changed = false
      for j := range rocks[0] {
        for i := range rocks {
          if rocks[i][j] == RoundRock && j < len(rocks[i]) - 1 && rocks[i][j+1] == NothingRock {
            rocks[i][j] = NothingRock
            rocks[i][j+1] = RoundRock
            changed = true
          }
        }
      }
    }
    results[n] = north_load(rocks)
  }
  return results
}

func Day14() (string, string) {
  file, _ := os.Open("input/day14.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  rocks := make([][]RockType, len(input))
  for i := range input {
    rocks[i] = make([]RockType, len(input[i]))
    chars := strings.Split(input[i], "")
    for c := range chars {
      if chars[c] == "O" {
        rocks[i][c] = RoundRock
      } else if chars[c] == "#" {
        rocks[i][c] = CubeRock
      } else {
        rocks[i][c] = NothingRock
      }
    }
  }
  move_north := make([][]RockType, len(rocks))
  for i := range rocks { move_north[i] = make([]RockType, len(rocks[i])) }
  copy(move_north, rocks)
  changed := true
  for changed {
    changed = false
    for i := len(move_north) - 1; i >= 0; i-- {
      for j := range move_north[i] {
        if move_north[i][j] == RoundRock && i > 0 && move_north[i-1][j] == NothingRock {
          move_north[i][j] = NothingRock
          move_north[i-1][j] = RoundRock
          changed = true
        }
      }
    }
  }
  p1 := north_load(move_north)
  p2results := spin_cycle(rocks, 200)
  // (from prior knowledge to speed up the search)
  // guess that repeat cycle is between 30 and 40 iterations long
  // and already in place by iteration 150 and will be unique for
  // a sequence of 3 matching values
  loop := 0
  start := 150
  for i := 30; i <= 40; i++ {
    if p2results[start] == p2results[start+i] &&
       p2results[start+1] == p2results[start+1+i] &&
       p2results[start+2] == p2results[start+2+i] {
      loop = i;
      break
    }
  }
  p2 := predict_nth_iter(1000000000, loop, start, p2results[start:start+loop])
  return strconv.Itoa(p1), strconv.Itoa(p2)
}