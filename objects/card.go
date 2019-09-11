package objects

type card struct {
	Name     string            `json:"name,omitempty"`
	ID       uint32            `json:"id,omitempty"`
	Level    uint8             `json:"level,omitempty"`
	MaxLevel uint8             `json:"maxLevel,omitempty"`
	IconUrls map[string]string `json:"iconUrls,omitempty"`
}
