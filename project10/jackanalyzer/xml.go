package jackanalyzer

import "encoding/xml"

type class struct {
	Name       xml.Name `xml:"class"`
	Keyword    string   `xml:"keyword"`
	Identifier string   `xml:"identifier"`
	Symbol     string   `xml:"symbol"`
}

type classVarDec struct {
	Name     xml.Name      `xml:"classVarDec"`
	Elements []interface{} `xml:",any"`
}

type Element struct {
	XMLName xml.Name `xml:"keyword"`
	Value   string   `xml:",chardata"`
}

type Keyword struct {
	XMLName xml.Name `xml:"keyword"`
	Value   string   `xml:",chardata"`
}

type Symbol struct {
	XMLName xml.Name `xml:"symbol"`
	Value   string   `xml:",chardata"`
}

type Identifier struct {
	XMLName xml.Name `xml:"identifier"`
	Value   string   `xml:",chardata"`
}
