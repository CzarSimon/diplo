package province

// Types of provinces
const (
	Land   Type = "LAND"
	Costal      = "COSTAL"
	Water       = "WATER"
)

// Type name of a prvince type.
type Type string

// Province represents a teritory on the diplomacy map.
type Province struct {
	Name            string   `json:"name"`
	ShortName       string   `json:"shortName"`
	HasSupplyCenter bool     `json:"hasSupplyCenter"`
	Type            Type     `json:"type"`
	Neighbours      []string `json:"neighbours"`
}
