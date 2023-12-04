package days

import (
  "io"
  "math"
  "os"
  "slices"
  "strconv"
  "strings"
)

func add_to_total(winners_per_card *[]int, i int, all_cards *[]int) {
  *all_cards = append(*all_cards, i)
  for j := i + 1; j <= i + (*winners_per_card)[i]; j++ {
    add_to_total(winners_per_card, j, all_cards)
  }
}

func Day04() (string, string) {
  file, _ := os.Open("input/day04.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  p1 := 0
  var p2winnercounts []int
  for i := range input {
    nums := strings.Split(input[i], " | ")
    winner_strings := strings.Split(strings.Split(nums[0], ": ")[1], "")
    mynum_strings := strings.Split(nums[1], "")
    var winner_ints []int
    var mynum_ints []int
    for c := 0; c < len(winner_strings); c += 3 {
      val, _ := strconv.Atoi(strings.Trim(strings.Join(winner_strings[c:c+2], ""), " "))
      winner_ints = append(winner_ints, val)
    }
    for c := 0; c < len(mynum_strings); c += 3 {
      val, _ := strconv.Atoi(strings.Trim(strings.Join(mynum_strings[c:c+2], ""), " "))
      mynum_ints = append(mynum_ints, val)
    }
    
    n_wins := 0
    for n := range mynum_ints {
      if slices.Contains(winner_ints, mynum_ints[n]) { n_wins++ }
    }
    if n_wins > 0 { p1 += int(math.Pow(2, float64(n_wins - 1))) }
    p2winnercounts = append(p2winnercounts, n_wins)
  }

  var p2cards []int
  for i := range p2winnercounts {
    add_to_total(&p2winnercounts, i, &p2cards)
  }

  return strconv.Itoa(p1), strconv.Itoa(len(p2cards))
}
