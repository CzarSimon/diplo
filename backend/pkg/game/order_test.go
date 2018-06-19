package game

import (
	"testing"

	"github.com/CzarSimon/diplo/backend/pkg/province"
)

func TestHoldOrderString(t *testing.T) {
	order := Order{
		From: province.Bud,
		To:   province.Ser,
		Type: Hold,
		Unit: Army,
	}
	orderString := order.String()
	if orderString != "A Bud Hold" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestMoveOrderString(t *testing.T) {
	order := Order{
		From: province.Bud,
		To:   province.Ser,
		Type: Move,
		Unit: Army,
	}
	orderString := order.String()
	if orderString != "A Bud - Ser" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestSupportOrderString(t *testing.T) {
	order := Order{
		Unit: Army,
		From: province.Ser,
		Type: Support,
		Supporting: &Order{
			Unit: Fleet,
			From: province.Alb,
			To:   province.Gre,
			Type: Move,
		},
	}
	orderString := order.String()
	if orderString != "A Ser S F Alb - Gre" {
		t.Errorf("Order.String wrong result: %s", orderString)
	}
}

func TestConvoyOrderString(t *testing.T) {
	order := Order{
		Unit: Fleet,
		From: province.Ion,
		Type: Convoy,
		Convoying: &Order{
			Unit: Army,
			From: province.Nap,
			To:   province.Tun,
			Type: Move,
		},
	}
	orderString := order.String()
	if orderString != "F Ion C A Nap - Tun" {
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
