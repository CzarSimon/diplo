package game

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if board.Year != StartingYear {
		t.Errorf("Wrong starting year. Expected=%d Got=%d", StartingYear, board.Year)
	}

	initalPhase, err := board.Phase.Current()
	if err != nil {
		t.Fatalf("phases.Current unexpected error: %s", err)
	}
	if initalPhase != SpringPhase {
		t.Errorf("Wrong starting phase. Expected=%s Got=%s", SpringPhase, initalPhase)
	}
}

func TestPhaseCycle(t *testing.T) {
	expectedCycle := []Phase{SpringPhase, RetreatPhase, FallPhase, RetreatPhase, BuildPhase}
	expectedCycle = append(expectedCycle, expectedCycle...)
	expectedCycle = append(expectedCycle, expectedCycle...)
	phases := NewPhaseCycle()

	for i := 0; i < len(expectedCycle)-1; i++ {
		phase, err := phases.Current()
		if err != nil {
			t.Fatalf("phases.Current unexpected error: %s", err)
		}
		if phase != expectedCycle[i] {
			t.Errorf("Wrong current phase. Expected=%s Got=%s", expectedCycle[i], phase)
		}

		nextPhase, err := phases.Next()
		if err != nil {
			t.Fatalf("phases.Next unexpected error: %s", err)
		}
		if nextPhase != expectedCycle[i+1] {
			t.Errorf("Wrong next phase. Expected=%s Got=%s", expectedCycle[i+1], nextPhase)
		}
	}
}

func TestPhaseCycleError(t *testing.T) {
	phases := NewPhaseCycle()
	phases.CurrentIndex = phases.MaxIndex + 1

	_, err := phases.Current()
	if err != ErrPhaseOutOfSync {
		t.Error("phases.Current expected error not nil")
	}

	_, err = phases.Next()
	if err != ErrPhaseOutOfSync {
		t.Error("phases.Next expected error not nil")
	}
}
