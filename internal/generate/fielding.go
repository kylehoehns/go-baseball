package generate

import (
	"github.com/kylehoehns/go-baseball/internal/utils"
)

func AddFieldingPositions(roster *Roster, innings int) *Roster {

	positions := []string{"P", "1B", "2B", "3B", "SS", "LF", "CF", "RF"}
	fieldingPositions := make(map[int]map[string]string, innings)

	// Assign an initial position to each player
	initialPositions := make(map[string]string)
	for i, name := range roster.PlayerNames {
		initialPositions[name] = positions[i%len(positions)]
	}

	// Rotate positions for each inning
	for inning := 1; inning <= innings; inning++ {
		inningPositions := make(map[string]string)
		rotatedPositions := utils.Rotate(positions, 3*(inning-1))

		for i, name := range roster.PlayerNames {
			position := rotatedPositions[i%len(rotatedPositions)]
			inningPositions[position] = name
		}

		fieldingPositions[inning] = inningPositions
	}

	roster.FieldingPositions = fieldingPositions
	return roster
}
