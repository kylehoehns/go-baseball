package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kylehoehns/go-baseball/internal/generate"
)

const fileName = "roster"
const defaultNumberOfInnings = 9

func main() {
	args := os.Args[1:]
	innings := defaultNumberOfInnings
	if len(args) > 0 {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "First argument must be an integer representing the number of innings: %v\n", err)
			os.Exit(1)
		} else {
			innings = val
		}
	}

	roster, err := generate.GenerateRoster(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating roster: %v\n", err)
		os.Exit(1)
	}

	roster = generate.AddBattingOrder(roster, innings)
	roster = generate.AddFieldingPositions(roster, innings)
	generate.PrintRoster(roster)
}
