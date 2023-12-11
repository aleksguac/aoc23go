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

func gcd(a, b int) int {
  /* Greatest common divisor */
  for b != 0 {
    tmp := b
    b = a % b
    a = tmp
  }
  return a
}

func lcm(nums ...int) int {
  /* Lowest common multiple */
  if len(nums) == 0 { return 0}
  if len(nums) == 1 { return nums[0] }
  result := nums[0] * nums[1] / gcd(nums[0], nums[1])
  for i := 2; i < len(nums); i++ {
    result = lcm(result, nums[i])
  }
  return result
}

func XYtoN(x, y, n_cols int) int {
  return y*n_cols + x
}

