package province

import (
	"testing"
)

func TestProvinces(t *testing.T) {
	for shortName, province := range GetProvinces() {
		if shortName != province.ShortName {
			t.Errorf("Incorect province shortname: Expected=%s Got=%s", province.ShortName, shortName)
		}
	}
}

func TestNumberOfSupplyCenters(t *testing.T) {
	noCenters := 0
	for _, province := range GetProvinces() {
		if province.HasSupplyCenter {
			noCenters++
		}
	}
	if noCenters != 34 {
		t.Errorf("Wrong number of Supply Centers. Expected=34 Got=%d", noCenters)
	}
}

func TestProvinceTypes(t *testing.T) {
	testNumberOfType(19, Water, t)
	testNumberOfType(14, Land, t)
	testNumberOfType(42, Costal, t)
}

func TestProvinceConnections(t *testing.T) {
	provinces := GetProvinces()
	for _, province := range provinces {
		fromProvince := getParentProvince(province, provinces, t)
		for _, toName := range province.Neighbours {
			toProvince, ok := provinces[toName]
			if !ok {
				t.Fatalf("Province: %s not available\n", toName)
			}
			if !checkProvincePointBack(fromProvince, toProvince, t) {
				t.Errorf("Province %s does not point back to %s", toProvince.Name, fromProvince.Name)
			}
		}
	}
}

func getParentProvince(province Province, provinces map[ShortName]Province, t *testing.T) Province {
	if !province.HasParent() {
		return province
	}
	parentProvince, ok := provinces[province.Parent]
	if !ok {
		t.Fatalf("Province: %s not available\n", province.Parent)
	}
	return parentProvince
}

func checkProvincePointBack(from, to Province, t *testing.T) bool {
	for _, name := range to.Neighbours {
		if name == from.ShortName {
			return true
		}
	}
	return false
}

func testNumberOfType(expectedNumber int, provinceType Type, t *testing.T) {
	numberOfProvinces := 0
	for _, province := range GetProvinces() {
		if province.Type == provinceType && !province.HasParent() {
			numberOfProvinces++
		}
	}
	if numberOfProvinces != expectedNumber {
		t.Errorf("Wrong number of provinces of type: %s. Expected=%d Got=%d",
			provinceType, expectedNumber, numberOfProvinces)
	}
}
