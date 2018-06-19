package game

const (
	Austria Country = "Austria"
	England         = "England"
	France          = "France"
	Germany         = "Germany"
	Italy           = "Italy"
	Russia          = "Russia"
	Turkey          = "Turkey"

	Army  Unit = "Army"
	Fleet      = "Fleet"
)

// Country is the name of a country in a game.
type Country string

// Unit is the name of a unit type.
type Unit string

func (u Unit) Short() string {
	return string(u[0])
}
