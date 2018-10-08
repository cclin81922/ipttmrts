package main

import (
	"testing"
)

func TestFindNearTaipeiMRTStation(t *testing.T) {
	testcases := []struct {
		poi       string
		latitude  float64
		longitude float64
		station   string
	}{
		{"誠正國中", 25.054361, 121.619107, "南港展覽館"},
		{"中國信託金融園區", 25.059200, 121.615447, "南港軟體園區"},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.poi, func(t *testing.T) {
			t.Parallel()
			station := findNearTaipeiMRTStation(tc.latitude, tc.longitude)
			if station.NameTW != tc.station {
				t.Fatalf("expected %s | got %s", tc.station, station.NameTW)
			}
		})
	}
}
