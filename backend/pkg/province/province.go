package province

// Types of provinces
const (
	Land   Type = "Land"
	Costal      = "Costal"
	Water       = "Water"
)

// Type name of a prvince type.
type Type string

// Name province short name.
type ShortName string

// Province represents a teritory on the diplomacy map.
type Province struct {
	Name            string      `json:"name"`
	ShortName       ShortName   `json:"shortName"`
	Parent          ShortName   `json:"parent"`
	ChildProvinces  []ShortName `json:"childProvinces"`
	HasSupplyCenter bool        `json:"hasSupplyCenter"`
	Type            Type        `json:"type"`
	Neighbours      []ShortName `json:"neighbours"`
}

// HasParent checks if the province has a parent.
func (p *Province) HasParent() bool {
	return p.Parent != ""
}
