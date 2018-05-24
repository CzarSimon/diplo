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
