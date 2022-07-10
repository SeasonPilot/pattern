package proxy

import "testing"

func TestStationProxy_sell(t *testing.T) {
	var seller Seller

	stationProxy := &StationProxy{
		station: &Station{stock: 10},
	}
	seller = stationProxy
	seller.sell("season")
	seller.sell("season11")

	stationProxy.station.sell("42535")

}

func TestStation_sell(t *testing.T) {

}
