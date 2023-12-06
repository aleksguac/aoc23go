package days

import "strconv"

func int_list(strlist []string) []int {
  /* Convert list of strings to list of ints, ignore errors */
  intlist := make([]int, len(strlist))
  for i := range strlist {
    v, _ := strconv.Atoi(strlist[i])
    intlist[i] = v
  }
  return intlist
}

func float64_list(strlist []string) []float64 {
  /* Convert list of strings to list of float64s, ignore errors */
  flist := make([]float64, len(strlist))
  for i := range strlist {
    v, _ := strconv.Atoi(strlist[i])
    flist[i] = float64(v)
  }
  return flist
}