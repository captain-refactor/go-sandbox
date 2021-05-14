package main

import (
	"fmt"
	"github.com/captain-refactor/go-sandbox/robot_boat"
	"sort"
)

func main() {
	solutions, iterations := robot_boat.SolvePuzzle()
	sort.Slice(solutions, func(i, j int) bool {
		return len(solutions[i]) > len(solutions[j])
	})
	fmt.Printf("Iterations: %d\n", iterations)
	for _, transitions := range solutions {
		fmt.Printf("\nSolution %+v\n", len(transitions)-1)
		for _, t := range transitions {
			fmt.Printf("%+v\n", t)
		}
	}
}
