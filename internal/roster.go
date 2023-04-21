package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/kylehoehns/go-baseball/internal/utils"
)

type Roster struct {
	PlayerNames       []string
	BattingOrder      map[int][]string
	FieldingPositions map[int]map[string]string
}

func GenerateRoster(fileName string) (*Roster, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	playerNames := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			playerNames = append(playerNames, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Roster{
		PlayerNames: utils.Shuffle(playerNames),
	}, nil

}
