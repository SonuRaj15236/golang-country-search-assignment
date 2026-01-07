package model

type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Currency   any    `json:"currency"`
	Population uint64 `json:"population"`
}

type Req struct {
	CountryName string `json:"countryName"`
}
