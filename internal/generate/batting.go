package generate

func AddBattingOrder(roster *Roster, innings int) *Roster {
	battingOrders := make(map[int][]string, innings)

	playersCount := len(roster.PlayerNames)
	for inning := 1; inning <= innings; inning++ {
		battingOrder := make([]string, playersCount)

		for i, name := range roster.PlayerNames {
			newPosition := (i + 3*(inning-1)) % playersCount
			battingOrder[newPosition] = name
		}

		battingOrders[inning] = battingOrder
	}

	roster.BattingOrder = battingOrders

	return roster
}
