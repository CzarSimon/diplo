package province

import "testing"

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

func TestProvinceConnections(t *testing.T) {
	provinces := GetProvinces()
	for _, province := range provinces {
		for _, toName := range province.Neighbours {
			toProvince, ok := provinces[toName]
			if !ok {
				t.Fatalf("Province: %s not available\n", toName)
			}
			if !checkProvincePointBack(province, toProvince, t) {
				t.Errorf("Province %s does not point back to %s", toProvince.Name, province.Name)
			}
		}
	}
}

func checkProvincePointBack(from, to Province, t *testing.T) bool {
	for _, name := range to.Neighbours {
		if name == from.ShortName {
			return true
		}
	}
	return false
}
