package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

type FocalLens struct {
  label string;
  focal_length int;
}

func lensboxindex(elem string, list []FocalLens) int {
  for i := range list {
    if elem == list[i].label {
      return i
    }
  }
  return -1
}

func Day15() (string, string) {
  file, _ := os.Open("input/day15.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), ",")
  p1 := 0
  for i := range input {
    hash := 0
    for c := range input[i] {
      hash = (hash + int(input[i][c])) * 17 % 256
    }
    p1 += hash
  }

  boxes := make(map[int][]FocalLens)
  for i := range input {
    hash := 0
    label := ""
    flen := -1
    if strings.Contains(input[i], "-") {
      label = strings.Split(input[i], "-")[0]
    } else {
      s := strings.Split(input[i], "=")
      label = s[0]
      flen, _ = strconv.Atoi(s[1])
    }
    for c := range label { hash = (hash + int(input[i][c])) * 17 % 256 }
    box, ok := boxes[hash]
    if flen == -1 {
      if ok {
        if ind := lensboxindex(label, box); ind != -1 {
          boxes[hash] = append(boxes[hash][:ind], boxes[hash][ind+1:]...)
        }
      }
    } else {
      lens := FocalLens{label, flen}
      if ok {
        if ind := lensboxindex(label, box); ind != -1 {
          boxes[hash][ind] = lens
        } else {
          boxes[hash] = append(boxes[hash], lens)
        }
      } else {
        boxes[hash] = []FocalLens{lens}
      }
    }
  }
  p2 := 0
  for b, v := range boxes {
    for i := range v {
      p2 += (b+1) * (i+1) * v[i].focal_length
    }
  }
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
