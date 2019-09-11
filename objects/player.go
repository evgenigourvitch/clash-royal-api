package objects

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Tag               string `json:"tag,omitempty"`
	Name              string `json:"name,omitempty"`
	Role              string `json:"role,omitempty"`
	LastSeen          string `json:"lastSeen,omitempty"`
	ExpLevel          uint8  `json:"expLevel,omitempty"`
	Trophies          uint16 `json:"trophies,omitempty"`
	Arena             *arena `json:"arena,omitempty"`
	ClanRank          uint8  `json:"clanRank,omitempty"`
	PreviousClanRank  uint8  `json:"previousClanRank,omitempty"`
	Donations         uint16 `json:"donations,omitempty"`
	DonationsReceived uint16 `json:"donationsReceived,omitempty"`
	ClanChestPoints   uint16 `json:"clanChestPoints,omitempty"`
}

func (p *Player) print() {
	if p == nil {
		return
	}
	fmt.Printf("%s\t%s\t%s\t%d\n", p.Name, p.Tag, p.Role, p.Trophies)
}

type PlayersResponse struct {
	Players []*Player `json:"items,omitempty"`
}

func (pr *PlayersResponse) Print() {
	if pr == nil {
		return
	}
	for _, player := range pr.Players {
		player.print()
	}
}

func ParsePlayers(bytesArr []byte) (*PlayersResponse, error) {
	result := PlayersResponse{}
	if err := json.Unmarshal(bytesArr, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
