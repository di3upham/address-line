package addressline

import "testing"

func TestSubdivisionName(t *testing.T) {
	subdivisionName := SubdivisionName("VN-67")
	if subdivisionName != "Nam Định" {
		t.Error("SubdivisionName not working")
	}
}

func TestCountryName(t *testing.T) {
	countryName := CountryName("VN")
	if countryName != "Viet Nam" {
		t.Error("CountryName not working")
	}
	countryName = CountryName("VNM")
	if countryName != "Viet Nam" {
		t.Error("CountryName not working")
	}
}

func TestSubdivisionCode(t *testing.T) {
	subdivisionCode := SubdivisionCode("VN", "Nam Định")
	if subdivisionCode != "VN-67" {
		t.Error("SubdivisionCode not working", subdivisionCode)
	}
}
