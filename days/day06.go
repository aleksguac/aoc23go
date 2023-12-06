package days

import (
  "io"
  "math"
  "os"
  "strconv"
  "strings"
)

func combine_string_digits(input string) float64 {
  result := 0
  input_list := strings.Split(input, "")
  for s := range input_list {
    v, err := strconv.Atoi(input_list[s])
    if err == nil {
      result = result * 10 + v
    }
  }
  return float64(result)
}

func get_roots(tmax float64, dmax float64) int {
  // find roots of inequality and return number of integers that satisfy equation
  // t*t - tmax*t + dmax < 0
  sqrt := math.Sqrt(tmax*tmax/4 - dmax)
  root1 := tmax/2 - sqrt
  root2 := tmax/2 + sqrt
  return int(math.Floor(root2) - math.Ceil(root1) + 1)
}

func Day06() (string, string) {
  file, _ := os.Open("input/day06.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  
  times1 := float64_list(strings.Fields(input[0])[1:])
  distances1 := float64_list(strings.Fields(input[1])[1:])
  p1 := 1
  for race := range times1 { p1 *= get_roots(times1[race], distances1[race]) }
  
  time2 := combine_string_digits(input[0])
  distance2 := combine_string_digits(input[1])
  p2 := get_roots(time2, distance2)

  return strconv.Itoa(p1), strconv.Itoa(p2)
}
