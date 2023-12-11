package days

import (
  "io"
  "os"
  "strconv"
  "strings"
)

func galaxies_in_row(grid *[][]bool, row int) int {
  total := 0
  for col := range (*grid)[row] {
    if (*grid)[row][col] { total++ }
  }
  return total
}

func galaxies_in_column(grid *[][]bool, col int) int {
  total := 0
  for row := range (*grid) {
    if (*grid)[row][col] { total++ }
  }
  return total
}

func distance(point1, point2 []int, empty_rows, empty_columns *[]bool) (int, int) {
  total1 := 0
  total2 := 0
  if point1[0] > point2[0] {
    total1 += point1[0] - point2[0]
    total2 += point1[0] - point2[0]
    for row := point2[0] + 1; row < point1[0]; row++ {
      if (*empty_rows)[row] { total1++; total2+=999999 }
    }
  } else {
    total1 += point2[0] - point1[0]
    total2 += point2[0] - point1[0]
    for row := point1[0] + 1; row < point2[0]; row++ {
      if (*empty_rows)[row] { total1++; total2+=999999 }
    }
  }
  if point1[1] > point2[1] {
    total1 += point1[1] - point2[1]
    total2 += point1[1] - point2[1]
    for col := point2[1] + 1; col < point1[1]; col++ {
      if (*empty_columns)[col] { total1++; total2+=999999 }
    }
  } else {
    total1 += point2[1] - point1[1]
    total2 += point2[1] - point1[1]
    for col := point1[1] + 1; col < point2[1]; col++ {
      if (*empty_columns)[col] { total1++; total2+=999999 }
    }
  }
  return total1, total2
}

func Day11() (string, string) {
  file, _ := os.Open("input/day11.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
  grid := make([][]bool, len(input))
  var galaxies [][]int
  for i := range input {
    chars := strings.Split(input[i], "")
    grid[i] = make([]bool, len(chars))
    for j := range chars {
      if chars[j] == "#" {
        galaxies = append(galaxies, []int{i, j})
        grid[i][j] = true
      } else {
        grid[i][j] = false
      }
    }
  }

  for i := range grid {
    for j := range grid {
      if grid[i][j] {
      }
    }
  }

  empty_rows := make([]bool, len(grid))
  for i := range grid {
    if galaxies_in_row(&grid, i) == 0 { empty_rows[i] = true }
  }
  empty_columns := make([]bool, len(grid[0]))
  for j := range grid[0] {
    if galaxies_in_column(&grid, j) == 0 { empty_columns[j] = true }
  }

  p1 := 0
  p2 := 0
  for g1 := range galaxies {
    for g2 := 0; g2 < g1; g2++ {
      p1part, p2part := distance(galaxies[g1], galaxies[g2], &empty_rows, &empty_columns)
      p1 += p1part
      p2 += p2part
    }
  }

  return strconv.Itoa(p1), strconv.Itoa(p2)
}
