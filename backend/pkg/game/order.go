package game

import (
	"fmt"

	"github.com/CzarSimon/diplo/backend/pkg/id"
	"github.com/CzarSimon/diplo/backend/pkg/province"
)

const (
	Build   OrderType = "Build"
	Convoy            = "Convoy"
	Disband           = "Disband"
	Hold              = "Hold"
	Move              = "Move"
	Retreat           = "Retreat"
	Support           = "Support"

	Invalid      OrderResult = "Invalid"
	Successful               = "Successful"
	Unsuccessful             = "Unsuccessful"
)

// OrderType name of an order type.
type OrderType string

// OrderResult describes an orders outcome.
type OrderResult string

// Order describes a move by a player in the game.
type Order struct {
	ID         string
	Unit       Unit
	To         province.ShortName
	From       province.ShortName
	Type       OrderType
	Supporting *Order
}

// NewOrder creates a new order.
func NewOrder(unit Unit, from, to province.ShortName, orderType OrderType, supporting *Order) Order {
	return Order{
		ID:         id.New(),
		Unit:       unit,
		To:         to,
		From:       from,
		Type:       orderType,
		Supporting: supporting,
	}
}

// NewHoldOrder creates a new hold order.
func NewHoldOrder(unit Unit, from province.ShortName) Order {
	return NewOrder(unit, from, "", Hold, nil)
}

// NewMoveOrder creates a new move order.
func NewMoveOrder(unit Unit, from, to province.ShortName) Order {
	return NewOrder(unit, from, to, Move, nil)
}

// NewSupportingOrder creates a new order of either type Support or Convoy.
func NewSupportingOrder(unit Unit, from province.ShortName, supportType OrderType, target *Order) Order {
	return NewOrder(unit, from, "", supportType, target)
}

// NewBuildOrder creats a new build order.
func NewBuildOrder(unit Unit, to province.ShortName) Order {
	return NewOrder(unit, "", to, Build, nil)
}

// NewDisbandOrder creates a new disband order.
func NewDisbandOrder(unit Unit, from province.ShortName) Order {
	return NewOrder(unit, from, "", Disband, nil)
}

// NewRetreatOrder creates a new retreat order.
func NewRetreatOrder(unit Unit, from, to province.ShortName) Order {
	return NewOrder(unit, from, to, Retreat, nil)
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
		return fmt.Sprintf("%s %s C %s", o.Unit.Short(), o.From, o.Supporting.String())
	case Build:
		return fmt.Sprintf("%s %s %s", o.Type, o.Unit.Short(), o.To)
	case Disband:
		return fmt.Sprintf("%s %s %s", o.Type, o.Unit.Short(), o.From)
	case Retreat:
		return fmt.Sprintf("%s %s %s - %s", o.Type, o.Unit.Short(), o.From, o.To)
	default:
		return fmt.Sprintf("Invalid order type: %s", o.Type)
	}
}
