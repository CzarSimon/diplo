package game

import (
	"testing"

	"github.com/CzarSimon/diplo/backend/pkg/province"
)

func TestHoldOrderString(t *testing.T) {
	order := NewHoldOrder(Army, province.Bud)
	orderString := order.String()
	if orderString != "A Bud Hold" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestMoveOrderString(t *testing.T) {
	order := NewMoveOrder(Army, province.Bud, province.Ser)
	orderString := order.String()
	if orderString != "A Bud - Ser" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestSupportOrderString(t *testing.T) {
	moveOrder := NewMoveOrder(Fleet, province.Alb, province.Gre)
	order := NewSupportingOrder(Army, province.Ser, Support, &moveOrder)
	orderString := order.String()
	if orderString != "A Ser S F Alb - Gre" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
	holdOrder := NewHoldOrder(Army, province.Rum)
	order = NewSupportingOrder(Fleet, province.Bla, Support, &holdOrder)
	orderString = order.String()
	if orderString != "F Bla S A Rum Hold" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestConvoyOrderString(t *testing.T) {
	moveOrder := NewMoveOrder(Army, province.Nap, province.Tun)
	order := NewSupportingOrder(Fleet, province.Ion, Convoy, &moveOrder)
	orderString := order.String()
	if orderString != "F Ion C A Nap - Tun" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestBuildOrderString(t *testing.T) {
	order := NewBuildOrder(Fleet, province.Lon)
	orderString := order.String()
	if orderString != "Build F Lon" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestDisbandOrderString(t *testing.T) {
	order := NewDisbandOrder(Fleet, province.Nth)
	orderString := order.String()
	if orderString != "Disband F Nth" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestReteatOrderString(t *testing.T) {
	order := NewRetreatOrder(Army, province.StP, province.Mos)
	orderString := order.String()
	if orderString != "Retreat A StP - Mos" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestInvalidTypeOrderString(t *testing.T) {
	order := Order{
		Type: "Invalid",
	}
	orderString := order.String()
	if orderString != "Invalid order type: Invalid" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}
