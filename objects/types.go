package objects

const (
	EResponseTypePlayersList ResponseType = iota
	EResponseTypeBattles
	EResponseTypeLocations
	EResponseTypeClans
)

type ResponseType uint8
