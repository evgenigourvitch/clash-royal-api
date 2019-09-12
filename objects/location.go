package objects

import (
	"encoding/json"
)

type Location struct {
	Name        string `json:"name,omitempty"`
	ID          uint32 `json:"id,omitempty"`
	IsCountry   bool   `json:"isCountry,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type LocationsResponse struct {
	Locations []*Location `json:"items,omitempty"`
}

func ParseLocations(bytesArr []byte) (*LocationsResponse, error) {
	result := LocationsResponse{}
	if err := json.Unmarshal(bytesArr, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
