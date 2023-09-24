package currency

type Currency struct {
	CrossOrder int    `xml:"CrossOrder,attr" json:"cross_order,omitempty"`
	Code       string `xml:"CurrencyCode,attr" json:"code,omitempty"`
	Unit       int    `xml:"Unit" json:"unit,omitempty"`
	Title      string `xml:"Isim" json:"title,omitempty"`
	Buying     string `xml:"BanknoteBuying" json:"buying,omitempty"`
	Selling    string `xml:"BanknoteSelling" json:"selling,omitempty"`
}
