package game

import (
	"fmt"

	"github.com/CzarSimon/diplo/backend/pkg/province"
)

const (
	Hold    OrderType = "Hold"
	Move              = "Move"
	Support           = "Support"
	Convoy            = "Convoy"
)

// OrderType name of an order type.
type OrderType string

// Order describes a move by a player in the game.
type Order struct {
	Unit       Unit
	To         province.ShortName
	From       province.ShortName
	Type       OrderType
	Supporting *Order
	Convoying  *Order
}

// String returns a string representation of an order.
func (o *Order) String() string {
	switch o.Type {
	case Hold:
		return fmt.Sprintf("%s %s %s", o.Unit.Short(), o.From, o.Type)
	case Move:
		return fmt.Sprintf("%s %s - %s", o.Unit.Short(), o.From, o.To)
	case Support:
		return fmt.Sprintf("%s %s S %s", o.Unit.Short(), o.From, o.Supporting.String())
	case Convoy:
		return fmt.Sprintf("%s %s C %s", o.Unit.Short(), o.From, o.Convoying.String())
	default:
		return fmt.Sprintf("Invalid order type: %s", o.Type)
	}
}
