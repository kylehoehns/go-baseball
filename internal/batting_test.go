package internal

import (
	"reflect"
	"testing"
)

func TestAddBattingOrder(t *testing.T) {
	testCases := []struct {
		name    string
		roster  *Roster
		innings int
	}{
		{
			name: "Sample roster with 3 innings",
			roster: &Roster{
				PlayerNames: []string{"Alice", "Bob", "Charlie", "David"},
			},
			innings: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rosterWithBattingOrder := AddBattingOrder(tc.roster, tc.innings)

			if len(rosterWithBattingOrder.BattingOrder) != tc.innings {
				t.Fatalf("Expected %d innings, got %d", tc.innings, len(rosterWithBattingOrder.BattingOrder))
			}

			for inning, order := range rosterWithBattingOrder.BattingOrder {
				expectedOrder := make([]string, len(tc.roster.PlayerNames))
				for i, name := range tc.roster.PlayerNames {
					newPosition := (i + 3*(inning-1)) % len(tc.roster.PlayerNames)
					expectedOrder[newPosition] = name
				}

				if !reflect.DeepEqual(order, expectedOrder) {
					t.Errorf("Inning %d: expected %v, got %v", inning, expectedOrder, order)
				}
			}
		})
	}
}
