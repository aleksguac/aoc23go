package days

import (
  "cmp"
  "io"
  "os"
  "slices"
  "strconv"
  "strings"
)

type Conversion struct {
  source int
  destination int
  range_ int
}

func cmpConversion (a, b Conversion) int {
  return cmp.Compare(a.source, b.source)
}

func get_gardening_conversions(input []string) [][]Conversion {
  /* Process conversion input strings into one array:
   *  [seed-to-soil, soil-to-fertiliser, fertiliser-to-water, water-to-light,
   *   light-to-temperature, temperature-to-humidity, humidity-to-location]
   * 
   * Each element is a list of `Conversion`s */
  var conversions [][]Conversion
  for stage := range input {
    conversions = append(conversions, []Conversion{})
    lines := strings.Split(input[stage], "\n")[1:]
    for i_line := range lines {
      vals := int_list(strings.Split(lines[i_line], " "))
      conversions[stage] = append(conversions[stage], Conversion{vals[1], vals[0], vals[2]})
    }
    slices.SortFunc(conversions[stage], cmpConversion)
  }
  return conversions
}

func get_index_between_sources(conversions *[]Conversion, n int) int {
  /* For a list of conversions and an integer n, return the index where the value of n is between
   * that Conversion's source value and the next, or -1 if n is lower than the first source value
   * e.g. conversions sources = {0, 10, 20, 30} and n = 25 returns 2 */
  for c := range *conversions {
    if (*conversions)[c].source <= n && (c == len(*conversions) - 1 || n < (*conversions)[c+1].source) { return c }
  }
  return -1
}

func day5part1(seeds []int, conversions [][]Conversion) int {
  values := make([]int, len(seeds))
  copy(values, seeds)
  for c := range conversions {
    // for each stage of the conversion, swap each value for its destination equivalent
    for v := range values {
      d := get_index_between_sources(&conversions[c], values[v])
      if d != -1 && conversions[c][d].source + conversions[c][d].range_ > values[v] {
        values[v] = conversions[c][d].destination + (values[v] - conversions[c][d].source)
      } // if no conversion found, the value stays the same
    }
  }
  return slices.Min(values)
}

func day5part2(seeds []int, conversions[][]Conversion) int {
  values := make([]int, len(seeds))
  copy(values, seeds)
  for c := range conversions {
    // for each stage of the conversion, process each range to its destination equivalent
    var next []int
    for v := 0; v < len(values)/2; v++ {
      current_start := values[2*v] // inclusive
      current_range := values[2*v+1]
      current_end := current_start + current_range // exclusive
      for {
        d := get_index_between_sources(&conversions[c], current_start)
        if d == -1 {
          if current_end <= conversions[c][0].source {
            /* 
             * source ranges:
             *       .......          .....      .........................
             * current range:
             *  ...
             */
            next = append(next, []int{current_start, current_range}...)
            break;
          } else {
            /* 
             * source ranges:
             *       .......          .....      .........................
             * current range:
             *  .........(or longer)
             */
            next = append(next, []int{current_start, conversions[c][0].source - current_start}...)
            current_start = conversions[c][0].source
            current_range = current_end - current_start
            d = 0
          }
        }
        conv_src_start := conversions[c][d].source
        conv_range := conversions[c][d].range_
        conv_src_end := conv_src_start + conv_range
        conv_dst_start := conversions[c][d].destination
        if current_start < conv_src_end {
          if current_end <= conv_src_end {
            /* 
             * source ranges:
             *       .......          .....      .........................
             * current range:
             *         ...
             */
            next = append(next, []int{conv_dst_start + current_start - conv_src_start, current_range}...)
            break
          } else {
            /* 
             * source ranges:
             *       .......          .....      .........................
             * current range:
             *          ............(or longer)
             */
            next = append(next, []int{conv_dst_start + current_start - conv_src_start, conv_src_end - current_start}...)
            current_start = conv_src_end
            current_range = current_end - current_start
          }
        } else if d == len(conversions[c]) - 1 || current_end <= conversions[c][d+1].source {
          /* 
           * source ranges:
           *       .......          .....      .........................
           * current range:
           *                  ...
           */
          next = append(next, []int{current_start, current_range}...)
          break
        } else {
          /* 
           * source ranges:
           *       .......          .....      .........................
           * current range:
           *                  ..........(or longer)
           */
          next = append(next, []int{current_start, conversions[c][d+1].source - current_start}...)
          current_start = conversions[c][d+1].source
          current_range = current_end - current_start
        }
      }
    }
    values = make([]int, len(next))
    copy(values, next)
  }
  
  min := values[0]
  for v := range values { if v % 2 == 0 && values[v] < min { min = values[v] }}
  return min
}

func Day05() (string, string) {
  file, _ := os.Open("input/day05.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n\n")
  
  seeds := int_list(strings.Split(strings.Split(input[0], ": ")[1], " "))
  conversions := get_gardening_conversions(input[1:])
 
  return strconv.Itoa(day5part1(seeds, conversions)), strconv.Itoa(day5part2(seeds, conversions))
}
