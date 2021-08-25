package addressline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func CountryName(code string) string {
	if code == "" {
		return ""
	}
	country := CountryCodeMap[code]
	if country == nil {
		return ""
	}
	return country.Name
}

func SubdivisionName(code string) string {
	if len(code) < 2 {
		return ""
	}
	country := CountryCodeMap[code[:2]]
	if country == nil {
		return ""
	}
	subdivision := country.SubdivisionCodeMap[code]
	if subdivision == nil {
		return ""
	}
	return subdivision.Name
}

type Iso_3166_1 struct {
	Countries []*Country `json:"3166-1"`
}

type Iso_3166_2 struct {
	Subdivisions []*Subdivision `json:"3166-2"`
}

var CountryCodeMap map[string]*Country
var CountryNameMap map[string]*Country

func init() {
	var err error
	iso31661b, err := ioutil.ReadFile("iso_3166-1.json")
	if err != nil {
		panic(err)
	}
	iso31661 := &Iso_3166_1{}
	err = json.Unmarshal(iso31661b, iso31661)
	if err != nil {
		panic(err)
	}
	CountryCodeMap = make(map[string]*Country, 2*len(iso31661.Countries))
	CountryNameMap = make(map[string]*Country, len(iso31661.Countries))
	for _, country := range iso31661.Countries {
		if CountryCodeMap[country.Alpha2] != nil || CountryCodeMap[country.Alpha3] != nil || CountryNameMap[country.Name] != nil {
			panic("duplicate country")
		}
		CountryCodeMap[country.Alpha2] = country
		CountryCodeMap[country.Alpha3] = country
		CountryNameMap[country.Name] = country
	}

	iso31662b, err := ioutil.ReadFile("iso_3166-2.json")
	if err != nil {
		panic(err)
	}
	iso31662 := &Iso_3166_2{}
	err = json.Unmarshal(iso31662b, iso31662)
	if err != nil {
		panic(err)
	}
	var country *Country
	for _, subdivision := range iso31662.Subdivisions {
		country = CountryCodeMap[subdivision.Code[:2]]
		if country == nil {
			panic("country not found")
		}
		if country.SubdivisionCodeMap == nil {
			country.SubdivisionCodeMap = make(map[string]*Subdivision)
		}
		if country.SubdivisionCodeMap[subdivision.Code] != nil {
			fmt.Println(country.SubdivisionCodeMap[subdivision.Code])
			fmt.Println(subdivision)
			panic("duplicate subdivision")
		}
		country.SubdivisionCodeMap[subdivision.Code] = subdivision

		subdivision.Country = country
	}
}
