package objects

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Clan struct {
	Tag               string    `json:"tag,omitempty"`
	Name              string    `json:"name,omitempty"`
	BadgeId           uint32    `json:"badgeId,omitempty"`
	Type              string    `json:"type,omitempty"`
	ClanScore         uint32    `json:"clanScore,omitempty"`
	ClanWarTrophies   uint32    `json:"clanWarTrophies,omitempty"`
	RequiredTrophies  uint32    `json:"requiredTrophies,omitempty"`
	DonationsPerWeek  uint32    `json:"donationsPerWeek,omitempty"`
	ClanChestLevel    uint32    `json:"clanChestLevel,omitempty"`
	ClanChestMaxLevel uint32    `json:"clanChestMaxLevel,omitempty"`
	Members           uint16    `json:"members,omitempty"`
	Location          *Location `json:"location,omitempty"`
}

type ClansResponse struct {
	Clans []*Clan `json:"items,omitempty"`
}

func (cr *ClansResponse) Dump(locationName string) {
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

func ParseClans(bytesArr []byte) (*ClansResponse, error) {
	result := ClansResponse{}
	if err := json.Unmarshal(bytesArr, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
