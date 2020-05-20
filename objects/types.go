package objects

const (
	EResponseTypePlayersList ResponseType = iota
	EResponseTypeBattles
	EResponseTypeLocations
	EResponseTypeClans
	EResponseTypeClansWarLog
)

type ResponseType uint8
