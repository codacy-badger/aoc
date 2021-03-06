package main

import (
	"flag"
	"fmt"

	"github.com/afarbos/aoc/pkg/mathematic"
	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/str"
	"github.com/afarbos/aoc/pkg/utils"
)

var flagInput string

const (
	separator      = "\n"
	distanceFormat = "%s to %s = %d"

	resShortest = 117
	resLongest  = 909
)

func init() {
	utils.Init(&flagInput)
}

func fDistances(directions []string, f func(...int) int) int {
	cityGraph, cities := make(map[string]map[string]int), make(map[string]struct{})

	for _, direction := range directions {
		var (
			src, dst string
			distance int
		)

		if n, err := fmt.Sscanf(direction, distanceFormat, &src, &dst, &distance); err == nil && n == 3 {
			if _, ok := cityGraph[dst]; !ok {
				cityGraph[dst] = make(map[string]int)
			}

			if _, ok := cityGraph[src]; !ok {
				cityGraph[src] = make(map[string]int)
			}

			cities[dst], cities[src] = struct{}{}, struct{}{}
			cityGraph[dst][src], cityGraph[src][dst] = distance, distance
		}
	}

	var res int = f()

	str.Permutations(cities, func(cities []string) {
		var sum int
		for i, city := range cities[1:] {
			sum += cityGraph[cities[i]][city]
		}
		res = f(res, sum)
	})

	return res
}

func main() {
	flag.Parse()

	directions := read.Strings(flagInput, separator)
	utils.AssertEqual(fDistances(directions, mathematic.MinInt), resShortest)
	utils.AssertEqual(fDistances(directions, mathematic.MaxInt), resLongest)
}
