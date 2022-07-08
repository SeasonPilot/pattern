package proxy_pattern_practice

import "testing"

func TestStationProxy_Sell(t *testing.T) {
	var fa Seller

	// s := Station{}
	s1 := StationProxy{}

	fa = s1
	fa.Sell()

}
