package objects

import (
	"encoding/json"
)

type Battle struct {
	Type               string               `json:"type,omitempty"`
	BattleTime         string               `json:"battleTime,omitempty"`
	IsLadderTournament bool                 `json:"isLadderTournament,omitempty"`
	DeckSelection      string               `json:"deckSelection,omitempty"`
	Arena              *arena               `json:"arena,omitempty"`
	GameMode           *gameMode            `json:"gameMode,omitempty"`
	Team               []*battleParticipant `json:"team,omitempty"`
	Opponent           []*battleParticipant `json:"opponent,omitempty"`
}

func ParseBattles(bytesArr []byte) ([]*Battle, error) {
	result := []*Battle{}
	if err := json.Unmarshal(bytesArr, &result); err != nil {
		return nil, err
	}
	return result, nil
}
