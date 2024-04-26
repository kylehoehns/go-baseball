package internal

import (
	"github.com/kylehoehns/go-baseball/internal/utils"
)

func AddFieldingPositions(roster *Roster, innings int) *Roster {

	fieldingPositions := []string{"LF", "RF", "CF", "P", "1B", "2B", "SS", "3B"}

	finalPositions := make(map[int]map[string]string, innings)

	for inning := 1; inning <= innings; inning++ {
		inningPositions := make(map[string]string)

		// Rotate positions within each group
		rotatedGroup := utils.Rotate(fieldingPositions, 3*(inning-1))

		// Assign positions to players
		for i, name := range roster.PlayerNames {
			inningPositions[rotatedGroup[i]] = name
		}

		finalPositions[inning] = inningPositions
	}

	roster.FieldingPositions = finalPositions
	return roster
}
