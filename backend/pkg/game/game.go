package game

import (
	"errors"

	"github.com/CzarSimon/diplo/backend/pkg/province"
)

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

	BuildPhase   Phase = "Build"
	FallPhase          = "Fall"
	RetreatPhase       = "Retreat"
	SpringPhase        = "Spring"

	StartingYear int = 1901
)

var (
	ErrPhaseOutOfSync = errors.New("PhaseCycle out of sync")
)

// Country is the name of a country in a game.
type Country string

// Unit is the name of a unit type.
type Unit string

func (u Unit) Short() string {
	return string(u[0])
}

// Phase is a part of a year in the game.
type Phase string

// Board represents the current state of the game board.
type Board struct {
	Year      int                                      `json:"year"`
	Phase     PhaseCycle                               `json:"phase"`
	Positions map[Country][]Position                   `json:"positions"`
	Provinces map[province.ShortName]province.Province `json:"provinces"`
}

// NewBoard returns a game board in its inital state.
func NewBoard() Board {
	return Board{
		Year:      StartingYear,
		Phase:     NewPhaseCycle(),
		Provinces: province.GetProvinces(),
	}
}

// Position represents a countrys position in a province.
type Position struct {
	Country   Country            `json:"country"`
	Unit      Unit               `json:"unit"`
	Province  province.ShortName `json:"province"`
	Dislodged bool               `json:"dislodged"`
}

// PhaseCycle contains an ordered list of phases, and track of the current phase.
type PhaseCycle struct {
	CurrentIndex int     `json:"currentIndex"`
	MaxIndex     int     `json:"maxIndex"`
	Phases       []Phase `json:"phases"`
}

// NewPhaseCycle createas a new PhaseCycle poiting at the begining phase.
func NewPhaseCycle() PhaseCycle {
	return PhaseCycle{
		CurrentIndex: 0,
		MaxIndex:     4,
		Phases:       []Phase{SpringPhase, RetreatPhase, FallPhase, RetreatPhase, BuildPhase},
	}
}

// Current gets the current cycle phase, returns an error cycle index out of bounds.
func (pc *PhaseCycle) Current() (Phase, error) {
	if pc.CurrentIndex > pc.MaxIndex || pc.CurrentIndex < 0 {
		return "", ErrPhaseOutOfSync
	}
	return pc.Phases[pc.CurrentIndex], nil
}

// Next returns the next phase in the cycle and advances its state.
func (pc *PhaseCycle) Next() (Phase, error) {
	if _, err := pc.Current(); err != nil {
		return "", err
	}
	if pc.CurrentIndex < pc.MaxIndex {
		pc.CurrentIndex++
	} else {
		pc.CurrentIndex = 0
	}
	return pc.Phases[pc.CurrentIndex], nil
}
