package days

import (
  "cmp"
  "fmt"
  "io"
  "os"
  "slices"
  "strconv"
  "strings"
)

type Hand struct {
  cards []int;
  bid int;
  hand_type int; // 7 five of a kind, 6 four of a kind, 5 full house, 4 three of a kind, 3 two pair, 2 one pair, 1 high card, 0 unset
}

func get_handints_from_handstr(hand string) []int {
  var handints []int
  handstrs := strings.Split(hand, "")
  for h := range handstrs {
    v, err := strconv.Atoi(handstrs[h])
    if err == nil {
      handints = append(handints, v)
    } else {
      switch handstrs[h] {
      case "T":
        handints = append(handints, 10)
      case "J":
        handints = append(handints, 11)
      case "Q":
        handints = append(handints, 12)
      case "K":
        handints = append(handints, 13)
      case "A":
        handints = append(handints, 14)
      default:
        fmt.Println("weird hand value")
      }
    }
  }
  return handints
}

func get_hand_type_1(hand *Hand) {
  counts := make(map[int]int)
  for h := range hand.cards {
    _, ok := counts[hand.cards[h]]
    if ok {
      counts[hand.cards[h]]++
    } else {
      counts[hand.cards[h]] = 1
    }
  }
  n_pairs := 0
  n_threes := 0
  for _, v := range counts {
    if v == 5 {
      hand.hand_type = 7
      return
    } else if v == 4 {
      hand.hand_type = 6
      return
    } else if v == 3 {
      n_threes++
    } else if v == 2 {
      n_pairs++
    }
  }
  if n_threes == 1 && n_pairs == 1 {
    hand.hand_type = 5
  } else if n_threes == 1 {
    hand.hand_type = 4
  } else if n_pairs == 2 {
    hand.hand_type = 3
  } else if n_pairs == 1 {
    hand.hand_type = 2
  } else {
    hand.hand_type = 1
  }
}

func get_hand_type_2(hand *Hand) {
  counts := make(map[int]int)
  for h := range hand.cards {
    _, ok := counts[hand.cards[h]]
    if ok {
      counts[hand.cards[h]]++
    } else {
      counts[hand.cards[h]] = 1
    }
  }
  n_collections := []int{0, 0, 0, 0} // pairs, threes, fours, fives
  for k, v := range counts {
    if k != 1 {
      switch v {
      case 2:
        n_collections[0]++
      case 3:
        n_collections[1]++
      case 4:
        n_collections[2]++
      case 5:
        n_collections[3]++
      }
    }
  }
  j, ok := counts[1]
  if !ok { j = 0 }

  switch j {
  case 5: hand.hand_type = 7
  case 4: hand.hand_type = 7
  case 3:
    if n_collections[0] == 1 {
      hand.hand_type = 7
    } else {
      hand.hand_type = 6
    }
  case 2:
    if n_collections[1] == 1 {
      hand.hand_type = 7
    } else if n_collections[0] == 1 {
      hand.hand_type = 6
    } else {
      hand.hand_type = 4
    }
  case 1:
    if n_collections[2] == 1 {
      hand.hand_type = 7
    } else if n_collections[1] == 1 {
      hand.hand_type = 6
    } else if n_collections[0] == 2 {
      hand.hand_type = 5
    } else if n_collections[0] == 1 {
      hand.hand_type = 4
    } else {
      hand.hand_type = 2
    }
  case 0:
    if n_collections[3] == 1 {
      hand.hand_type = 7
    } else if n_collections[2] == 1 {
      hand.hand_type = 6
    } else if n_collections[1] == 1 && n_collections[0] == 1 {
      hand.hand_type = 5
    } else if n_collections[1] == 1 {
      hand.hand_type = 4
    } else if n_collections[0] == 2 {
      hand.hand_type = 3
    } else if n_collections[0] == 1 {
      hand.hand_type = 2
    } else {
      hand.hand_type = 1
    }
  }
}

func cmpHand(a, b Hand) int {
  if a.hand_type == b.hand_type {
    for i := range a.cards {
      if a.cards[i] != b.cards[i] || i == len(a.cards) - 1 {
        return cmp.Compare(a.cards[i], b.cards[i])
      }
    }
  }
  return cmp.Compare(a.hand_type, b.hand_type)
}

func Day07() (string, string) {
  file, _ := os.Open("input/day07.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  var hands []Hand
  for h := range input {
    var handstr string
    var bid int
    fmt.Sscanf(input[h], "%5s %d", &handstr, &bid)
    hand := Hand{get_handints_from_handstr(handstr), bid, 0}
    get_hand_type_1(&hand)
    hands = append(hands, hand)
  }

  // part 1
  slices.SortFunc(hands, cmpHand)
  p1 := 0
  for i := range hands {
    p1 += hands[i].bid * (i+1)
  }

  // part 2
  for i := range hands {
    for j := range hands[i].cards {
      if hands[i].cards[j] == 11 { hands[i].cards[j] = 1 } // set J to lowest value
    }
    get_hand_type_2(&hands[i]) // find new hand types
  }
  slices.SortFunc(hands, cmpHand)
  p2 := 0
  for i := range hands {
    p2 += hands[i].bid * (i+1)
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
