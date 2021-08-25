package addressline

import (
	"encoding/json"
	"io/ioutil"
)

func SubdivisionCode(countryCode, name string) string {
	if countryCode == "" || name == "" {
		return ""
	}
	country := CountryCodeMap[countryCode]
	if country == nil || len(country.SubdivisionNameMap[name]) == 0 {
		return ""
	}
	subdivisions := country.SubdivisionNameMap[name]
	subdivisionsLen := len(subdivisions)
	if subdivisionsLen == 1 {
		return subdivisions[0].Code
	}
	code := subdivisions[0].Code
	for i := 1; i < subdivisionsLen; i++ {
		if subdivisions[i].Parent == "" {
			code = subdivisions[i].Code
			break
		}
	}
	return code
}

func CountryCode(name string) (string, string) {
	if name == "" {
		return "", ""
	}
	country := CountryNameMap[name]
	if country == nil {
		return "", ""
	}
	return country.Alpha2, country.Alpha3
}

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
		if country.SubdivisionNameMap == nil {
			country.SubdivisionNameMap = make(map[string][]*Subdivision)
		}
		if country.SubdivisionCodeMap[subdivision.Code] != nil {
			panic("duplicate subdivision")
		}
		country.SubdivisionCodeMap[subdivision.Code] = subdivision
		country.SubdivisionNameMap[subdivision.Name] = append(country.SubdivisionNameMap[subdivision.Name], subdivision)

		subdivision.Country = country
	}
}
