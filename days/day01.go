package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

func word_to_digit(chars []string, c int) int {
  digit_words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
  for d := range digit_words {
    length := len(digit_words[d])
    if c >= length - 1 && strings.Join(chars[c-length+1:c+1], "") == digit_words[d] {
      return d
    }
  }
  return -1
}

func Day01() (string, string) {
  file, _ := os.Open("input/day01.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  p1 := 0
  p2 := 0
  for i := range input {
    p1digits := []int{0, 0}
    p2digits := []int{0, 0}
    chars := strings.Split(input[i], "")
    for c := range chars {
      if d := word_to_digit(chars, c); d != -1 {
        if p2digits[0] == 0 { p2digits[0] = d }
        p2digits[1] = d
      }
      if d, err := strconv.Atoi(chars[c]); err == nil {
        if p1digits[0] == 0 { p1digits[0] = d }
        if p2digits[0] == 0 { p2digits[0] = d }
        p1digits[1] = d
        p2digits[1] = d
      }
    }
    p1 += p1digits[0] * 10 + p1digits[1]
    p2 += p2digits[0] * 10 + p2digits[1]
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
