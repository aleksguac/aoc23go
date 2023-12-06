package main

import (
  "fmt"
  "os"
  "github.com/aleksguac/aoc23go/days"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Run day N with `go run . N`")
    return
  }
  day := os.Args[1]
  pt1 := ""
  pt2 := ""
  switch day {
    case "1":
      pt1, pt2 = days.Day01()
    case "2":
      pt1, pt2 = days.Day02()
    case "3":
      pt1, pt2 = days.Day03()
    case "4":
      pt1, pt2 = days.Day04()
    case "5":
      pt1, pt2 = days.Day05()
    case "6":
      pt1, pt2 = days.Day06()
    case "7":
    //   pt1, pt2 = days.Day07()
    case "8":
    //   pt1, pt2 = days.Day08()
    case "9":
    //   pt1, pt2 = days.Day09()
    case "10":
    //   pt1, pt2 = days.Day10()
    case "11":
    //   pt1, pt2 = days.Day11()
    case "12":
    //   pt1, pt2 = days.Day12()
    case "13":
    //   pt1, pt2 = days.Day13()
    case "14":
    //   pt1, pt2 = days.Day14()
    case "15":
    //   pt1, pt2 = days.Day15()
    case "16":
    //   pt1, pt2 = days.Day16()
    case "17":
    //   pt1, pt2 = days.Day17()
    case "18":
    //   pt1, pt2 = days.Day18()
    case "19":
    //   pt1, pt2 = days.Day19()
    case "20":
    //   pt1, pt2 = days.Day20()
    case "21":
    //   pt1, pt2 = days.Day21()
    case "22":
    //   pt1, pt2 = days.Day22()
    case "23":
    //   pt1, pt2 = days.Day23()
    case "24":
    //   pt1, pt2 = days.Day24()
    case "25":
    //   pt1, pt2 = days.Day25()
    default:
      fmt.Println("Invalid day specified")
      return
  }
  if pt1 == "" && pt2 == "" {
    fmt.Println("Not yet implemented")
  }

  if (pt1 != "") {
    fmt.Println("day", day, "part 1:", pt1)
  }
  if (pt2 != "") {
    fmt.Println("day", day, "part 2:", pt2)
  }
}
