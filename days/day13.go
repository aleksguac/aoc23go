package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

func get_reflection(pattern []string, part2 bool) int {
  // vertical reflection
  reflections := make([]int, len(pattern[0]) - 1)
  seen_smudge := make([]bool, len(reflections)) // part 2
  for r := range reflections { reflections[r] = r }
  for row := range pattern {
    for i := len(reflections) - 1; i >= 0; i-- {
      r := reflections[i]
      lower := r
      higher := r + 1
      is_reflection := true
      for lower >= 0 && higher < len(pattern[row]) {
        if pattern[row][lower] != pattern[row][higher] {
          if !part2 || seen_smudge[i] {
            is_reflection = false
            break
          } else if part2 {
            seen_smudge[i] = true
          }
        }
        lower--
        higher++
      }
      if !is_reflection {
        reflections = append(reflections[:i], reflections[i+1:]...)
        seen_smudge = append(seen_smudge[:i], seen_smudge[i+1:]...)
      }
    }
  }

  if len(reflections) == 1 && (!part2 || seen_smudge[0]) {
    return reflections[0] + 1
  } else if len(reflections) == 2 {
    if seen_smudge[0] {
      return reflections[0] + 1
    } else {
      return reflections[1] + 1
    }
  }

  // horizontal reflection
  reflections = make([]int, len(pattern) - 1)
  seen_smudge = make([]bool, len(reflections))
  for r := range reflections { reflections[r] = r }
  for col := range pattern[0] {
    for i := len(reflections) - 1; i >= 0; i-- {
      r := reflections[i]
      lower := r
      higher := r + 1
      is_reflection := true
      for lower >= 0 && higher < len(pattern) {
        if pattern[lower][col] != pattern[higher][col] {
          if !part2 || seen_smudge[i] {
            is_reflection = false
            break
          } else if part2 {
            seen_smudge[i] = true
          }
        }
        lower--
        higher++
      }
      if !is_reflection {
        reflections = append(reflections[:i], reflections[i+1:]...)
        seen_smudge = append(seen_smudge[:i], seen_smudge[i+1:]...)
      }
    }
  }
  
  if len(reflections) == 1 {
    return 100*(reflections[0] + 1)
  } else if len(reflections) == 2 {
    if seen_smudge[0] {
      return 100*(reflections[0] + 1)
    } else {
      return 100*(reflections[1] + 1)
    }
  }
  // shouldn't ever reach here
  return 0
}

func Day13() (string, string) {
  file, _ := os.Open("input/day13.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n\n")
  p1 := 0
  p2 := 0
  for p := range input {
    pattern := strings.Split(input[p], "\n")
    p1 += get_reflection(pattern, false)
    p2 += get_reflection(pattern, true)
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
