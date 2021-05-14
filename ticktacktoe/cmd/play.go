package main

import (
	"fmt"
	"github.com/captain-refactor/go-sandbox/ticktacktoe"
)

func main() {
	state := ticktacktoe.NewState(ticktacktoe.Circle, [3][3]ticktacktoe.Side{})
	solver := ticktacktoe.NewSolver(ticktacktoe.Cross)
	for {
		if end, _ := state.WhoWon(); end {
			break
		}
		fmt.Print(state)
		var x, y int
		var err error
		if state.NextTurn() == solver.Side() {
			fmt.Println("I am thinking...")
			x, y = solver.NextMove(state)
			state, err = state.CheckNWrite(x, y, solver.Side())
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Print("Your turn: ")
			_, err = fmt.Scan(&x, &y)
			if err != nil {
				panic(err)
			}
			fmt.Println()
			state, err = state.CheckNWrite(x, y, solver.Side().Flip())
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Print(state)
	_, side := state.WhoWon()
	switch side {
	case ticktacktoe.Empty:
		fmt.Println("Draw")
	case solver.Side().Flip():
		fmt.Println("You Win")
	case solver.Side():
		fmt.Println("You Lose")
	}
}
