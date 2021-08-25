package addressline

type Country struct {
	Alpha2             string `json:"alpha_2"`
	Alpha3             string `json:"alpha_3"`
	Name               string `json:"name"`
	Numeric            string `json:"numeric"`
	SubdivisionCodeMap map[string]*Subdivision
	SubdivisionNameMap map[string][]*Subdivision
}

type Subdivision struct {
	Country *Country
	Code    string `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Parent  string `json:"parent"`
}
