package objects

type battleParticipant struct {
	Tag                     string   `json:"tag,omitempty"`
	Name                    string   `json:"name,omitempty"`
	StartingTrophies        uint16   `json:"startingTrophies,omitempty"`
	Crowns                  uint8    `json:"crowns,omitempty"`
	KingTowerHitPoints      uint16   `json:"kingTowerHitPoints,omitempty"`
	PrincessTowersHitPoints []uint16 `json:"princessTowersHitPoints,omitempty"`
	Clan                    *clan    `json:"clan,omitempty"`
	Cards                   []*card  `json:"cards,omitempty"`
}
