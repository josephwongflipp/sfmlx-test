/*
This package is responsible business logic of inserting ads into an existing sfml file
*/

package sfmlparser

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/beevik/etree"
)

// Our custom advertisement element wrapper
type Advertisement struct {
	XMLName   xml.Name  `xml:"advertisement"`
	Height    string    `xml:"height,attr"`
	Width     string    `xml:"width,attr"`
	URLSource URLSource `xml:"url-source"`
}

// Our customer url-source element that lives inside an advetisement
type URLSource struct {
	XMLName xml.Name `xml:"url-source"`
	Height  string   `xml:"height,attr"`
	Width   string   `xml:"width,attr"`
	URL     string   `xml:"url,attr"`
}

// Generic error handler
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Slurps a file and returns its contents
func ReadSFMLFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	handleErr(err)
	return dat
}

// iterates into body's linear-layout and injects ads between all children and returns the editted xml
// TODO: Make this configurable
func InjectAdsIntoSFML(d []byte) []byte {
	doc := etree.NewDocument()
	doc.ReadFromBytes(d)

	// Grabbing reference to our body's linear-layout
	rootLL := doc.SelectElement("storefront").SelectElement("body").SelectElement("linear-layout")

	// result slice
	var s []*etree.Element

	// Iterate each item in the linear-layout
	for i, element := range rootLL.ChildElements() {
		s = append(s, element)

		// Do not inject an ad before the first index or after the last index.
		if i < len(rootLL.ChildElements())-1 {
			advertisement := generateAdvertisementElement()
			_, err := xml.Marshal(advertisement)
			handleErr(err)

			// create new xml element
			adXMLString, err := xml.Marshal(generateAdvertisementElement())
			handleErr(err)
			ae := etree.NewDocument()
			ae.ReadFromBytes(adXMLString)

			// fmt.Printf("Injected an ad after index %v\n", i)
			s = append(s, ae.FindElement("advertisement"))
		}
	}

	for i := 0; i < len(rootLL.ChildElements()); i++ {
		doc.SelectElement("storefront").SelectElement("body").SelectElement("linear-layout").RemoveChildAt(i)
	}

	for i := 0; i < len(s); i++ {
		doc.SelectElement("storefront").SelectElement("body").SelectElement("linear-layout").AddChild(s[i])
	}

	b, err := doc.WriteToBytes()
	handleErr(err)
	// fmt.Printf("%v", string(b))
	return b
}

// A static hardcoded advertisement element
func generateAdvertisementElement() Advertisement {
	urlSource := &URLSource{
		Height: "match-parent",
		Width:  "match-parent",
		URL:    "https://fh-uploads-addressreport.netdna-ssl.com/d79274ca-877f-437f-9ca0-e7230ea1ff46",
	}

	advertisement := &Advertisement{
		Height:    "wrap-content",
		Width:     "preserve-aspect",
		URLSource: *urlSource,
	}

	return *advertisement
}
