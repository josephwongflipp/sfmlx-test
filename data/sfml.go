/*
IGNORE
*/

package data

import "encoding/xml"

type Storefront struct {
	// root
	XMLName     xml.Name `xml:"storefront"`
	Title       string   `xml:"title,attr"`
	Subtitle    string   `xml:"subtitle,attr"`
	SfmlVersion string   `xml:"sfml-version,attr"`
	Head        Head     `xml:"head"`
	Body        Body     `xml:"body"`
}

type Head struct {
	XMLName xml.Name `xml:"head"`
}

type Body struct {
	XMLName          xml.Name     `xml:"body"`
	RootLinearLayout LinearLayout `xml:"linear-layout"`
}

type LinearLayout struct {
	XMLName       xml.Name        `xml:"linear-layout"`
	Height        string          `xml:"height,attr"`
	Width         string          `xml:"width,attr"`
	LinearLayout  []LinearLayout  `xml:"linear-layout"`
	Advertisement []Advertisement `xml:"advertisement"`
}

type Advertisement struct {
	XMLName   xml.Name  `xml:"advertisement"`
	Height    string    `xml:"height,attr"`
	Width     string    `xml:"width,attr"`
	URLSource URLSource `xml:"url-source"`
}

type URLSource struct {
	XMLName xml.Name `xml:"url-source"`
	Height  string   `xml:"height,attr"`
	Width   string   `xml:"width,attr"`
	URL     string   `xml:"url,attr"`
}

func FromXML(d []byte) *Storefront {
	sfml := &Storefront{}
	xml.Unmarshal([]byte(d), &sfml)
	return sfml
}
