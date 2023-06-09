package internal

import (
	"fmt"
	"strings"
)

func PrintRoster(roster *Roster) {
	fmt.Println("## Batting Order:")
	printBattingOrderTable(roster)
	fmt.Println("\n## Fielding Positions:")
	printFieldingPositionsTable(roster)
}

func printBattingOrderTable(roster *Roster) {
	innings := len(roster.BattingOrder)
	header := "| Inning | " + strings.Join(makeRange(1, innings+1, "%-10d"), " | ")
	fmt.Println(header)
	separator := "|--------|" + strings.Repeat("------------|", innings)
	fmt.Println(separator)

	for i := 0; i < len(roster.PlayerNames); i++ {
		row := make([]string, innings)
		for inning, order := range roster.BattingOrder {
			row[inning-1] = fmt.Sprintf("%-10s", order[i])
		}
		fmt.Println("|        | " + strings.Join(row, " | "))
	}
}

func printFieldingPositionsTable(roster *Roster) {
	innings := len(roster.FieldingPositions)
	header := "| Inning | " + strings.Join(makeRange(1, innings+1, "%-10d"), " | ")
	fmt.Println(header)
	separator := "|--------|" + strings.Repeat("------------|", innings)
	fmt.Println(separator)

	positions := []string{"P", "1B", "2B", "3B", "SS", "LF", "RF"}
	if len(roster.PlayerNames) == 8 {
		positions = append(positions, "CF")
	}
	for _, position := range positions {
		row := make([]string, innings)
		for inning, posMap := range roster.FieldingPositions {
			row[inning-1] = fmt.Sprintf("%-10s", posMap[position])
		}
		fmt.Printf("| %-6s | %s\n", position, strings.Join(row, " | "))
	}
}

func makeRange(min, max int, format string) []string {
	arr := make([]string, max-min)
	for i := range arr {
		arr[i] = fmt.Sprintf(format, min+i)
	}
	return arr
}
