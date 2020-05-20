package objects

import (
	"encoding/json"
)

type WarClan struct {
	*Clan
	Participants  uint16 `json:"participants,omitempty"`
	BattlesPlayed uint16 `json:"battlesPlayed,omitempty"`
	Wins          uint16 `json:"wins,omitempty"`
	Crowns        uint16 `json:"crowns,omitempty"`
}

type Standing struct {
	TrophyChange int16    `json:"trophyChange,omitempty"`
	Clan         *WarClan `json:"clan,omitempty"`
}

type Participant struct {
	*Player
	CardsEarned                uint16 `json:"cardsEarned,omitempty"`
	BattlesPlayed              uint8  `json:"battlesPlayed,omitempty"`
	Wins                       uint8  `json:"wins,omitempty"`
	CollectionDayBattlesPlayed uint8  `json:"collectionDayBattlesPlayed,omitempty"`
	NumberOfBattles            uint8  `json:"numberOfBattles,omitempty"`
}

type War struct {
	SeasonID     uint32         `json:"seasonId,omitempty"`
	CreatedDate  string         `json:"createdDate,omitempty"`
	Participants []*Participant `json:"participants,omitempty"`
	Standings    []*Standing    `json:"standings,omitempty"`
}

type WarsResponse struct {
	Wars []*War `json:"items,omitempty"`
}

/*
func (cr *WarResponse) Dump(locationName string) {
	tsvFile, err := os.Create(fmt.Sprintf("./output/clans/clan.%s.%d.tsv", locationName, time.Now().UnixNano()))
	if err != nil {
		fmt.Printf("failed to open tsv file: %+v\n", err)
		return
	}
	defer tsvFile.Close()
	for _, c := range cr.Clans {
		tsvFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%d\t%d\t%d\n", c.Name, c.Tag, c.Type, c.ClanScore, c.Members, c.DonationsPerWeek))
	}
}
*/

func ParseClanWarLogs(bytesArr []byte) (*WarsResponse, error) {
	result := WarsResponse{}
	if err := json.Unmarshal(bytesArr, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
