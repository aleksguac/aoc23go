package days

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func Day09() (string, string) {
  file, _ := os.Open("input/day09.txt")
  input_bytes, _ := io.ReadAll(file)
  input := strings.Split(string(input_bytes), "\n")
	p1 := 0
  p2 := 0
	for i := range input {
		var layers [][]int
		layers = append(layers, int_list(strings.Split(input[i], " ")))
		for {
			n_layers := len(layers)
			len_last := len(layers[n_layers-1])
			layers = append(layers, make([]int, len_last-1))
			allzero := true
			for j := range layers[n_layers] {
				layers[n_layers][j] = layers[n_layers-1][j+1] - layers[n_layers-1][j]
				if layers[n_layers][j] != 0 { allzero = false }
			}
			if allzero { break }
		}
    p1bit := 0
    p2bit := 0
    for l := len(layers) - 1; l >= 0; l-- {
      p1bit = layers[l][len(layers[l])-1] + p1bit
      p2bit = layers[l][0] - p2bit
    }
    p1 += p1bit
    p2 += p2bit
	}
  return strconv.Itoa(p1), strconv.Itoa(p2)
}
