package jackanalyzer

import "encoding/xml"

type class struct {
	Name     xml.Name      `xml:"class"`
	Elements []interface{} `xml:",any"`
}

type classVarDec struct {
	Name     xml.Name      `xml:"classVarDec"`
	Elements []interface{} `xml:",any"`
}

type subroutineDec struct {
	Name     xml.Name      `xml:"subroutineDec"`
	Elements []interface{} `xml:",any"`
}

type subroutineBody struct {
	Name     xml.Name      `xml:"subroutineBody"`
	Elements []interface{} `xml:",any"`
}

type varDec struct {
	Name     xml.Name      `xml:"varDec"`
	Elements []interface{} `xml:",any"`
}

type parameterList struct {
	Name     xml.Name      `xml:"parameterList"`
	Elements []interface{} `xml:",any"`
}

type Element struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
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
