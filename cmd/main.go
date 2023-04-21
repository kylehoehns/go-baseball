package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kylehoehns/go-baseball/internal"
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

	roster, err := internal.GenerateRoster(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating roster: %v\n", err)
		os.Exit(1)
	}

	roster = internal.AddBattingOrder(roster, innings)
	roster = internal.AddFieldingPositions(roster, innings)
	internal.PrintRoster(roster)
}
