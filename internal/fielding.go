package internal

import (
	"github.com/kylehoehns/go-baseball/internal/utils"
)

func AddFieldingPositions(roster *Roster, innings int) *Roster {
	group1 := []string{"LF", "RF"}
	if len(roster.PlayerNames) == 8 {
		group1 = append(group1, "CF")
	}
	group2 := []string{"P", "1B", "2B", "SS", "3B"}
	fieldingPositions := make(map[int]map[string]string, innings)

	initialPositions := make(map[string]string)
	for i, name := range roster.PlayerNames {
		initialPositions[name] = group1[i%len(group1)]
	}

	firstGroupRotationAmount := 1
	secondGroupRotationAmount := 4
	if len(roster.PlayerNames) != 8 {
		firstGroupRotationAmount = 2
		secondGroupRotationAmount = 5
	}

	for inning := 1; inning <= innings; inning++ {
		inningPositions := make(map[string]string)

		// Rotate positions within each group
		rotatedGroup1 := utils.Rotate(group1, firstGroupRotationAmount*(inning-1))
		rotatedGroup2 := utils.Rotate(group2, secondGroupRotationAmount*(inning-1))

		// Switch between groups each inning
		if inning%2 == 0 {
			rotatedGroup1, rotatedGroup2 = rotatedGroup2, rotatedGroup1
		}

		// Assign positions to players
		for i, name := range roster.PlayerNames {
			if i < len(rotatedGroup1) {
				inningPositions[rotatedGroup1[i]] = name
			} else {
				inningPositions[rotatedGroup2[i-len(rotatedGroup1)]] = name
			}
		}

		fieldingPositions[inning] = inningPositions
	}

	roster.FieldingPositions = fieldingPositions
	return roster
}
