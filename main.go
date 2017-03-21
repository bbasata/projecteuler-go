package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var solvers = map[string]func(){
	"1": solveProblem1,
	"2": solveProblem2,
	"4": solveProblem4,
	"5": solveProblem5,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <problem>\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	problem := os.Args[1]

	if _, ok := solvers[problem]; !ok {
		fmt.Fprintf(os.Stderr, "Sorry, I do not know how to solve problem %q\n", problem)
		os.Exit(2)
	}

	fmt.Printf("Solving problem %s\n", problem)
	solvers[problem]()
}
