package objects

type clan struct {
	Tag     string `json:"tag,omitempty"`
	Name    string `json:"name,omitempty"`
	BadgeId uint32 `json:"badgeId,omitempty"`
}
