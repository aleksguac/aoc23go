package days

import (
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func Day02() (string, string) {
  file, _ := os.Open("input/day02.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  var game_id int
  var n int
  var colour string
  p1 := 0
  p2 := 0
  for i := range input {
    game_info := strings.Split(input[i], ": ")
    cube_info := strings.Split(game_info[1], "; ")
    fmt.Sscanf(game_info[0], "Game %d", &game_id)
    p1possible := true
    p2minimums := []int{0, 0, 0}
    for set := range cube_info {
      set_info := strings.Split(cube_info[set], ", ")
      for cube := range set_info {
        fmt.Sscanf(set_info[cube], "%d %s", &n, &colour)
        if (colour == "red" && n > 12) || (colour == "green" && n > 13) || (colour == "blue" && n > 14) { p1possible = false }
        if colour == "red" && n > p2minimums[0] {
          p2minimums[0] = n
        } else if colour == "green" && n > p2minimums[1] {
          p2minimums[1] = n
        } else if colour == "blue" && n > p2minimums[2] {
          p2minimums[2] = n
        }
      }
    }
    if p1possible { p1 += game_id }
    p2 += p2minimums[0] * p2minimums[1] * p2minimums[2]
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
