/*
This package is responsible for all CRUD actions against an SFML file.

Actions include:
- Counting the number of valid advertisement slots in an SFML document
- Injecting advertisements in an SFML document
- Returning valid SFML
*/

package sfmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// Generic error handler
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Slurps a file and returns its contents
func ReadSFMLFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	checkError(err)
	return dat
}

// Generic XML node
type Node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []Node `xml:",any"`
}

// Root element of an SFML file
type Storefront struct {
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

// Only used to represent the Root linear-layout in the Body
type LinearLayout struct {
	XMLName xml.Name `xml:"linear-layout"`
	Height  string   `xml:"height,attr"`
	Width   string   `xml:"width,attr"`
	Content []byte   `xml:",innerxml"`
	Nodes   []Node   `xml:",any"`
}

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

// Reads in an SFML file and injects up to <maxAds> advertisements every <offset> slots
func injectAds(xmlData []byte, maxAds int, offset int) []byte {
	buf := bytes.NewBuffer(xmlData)
	dec := xml.NewDecoder(buf)

	var sf Storefront
	err := dec.Decode(&sf)
	checkError(err)

	nodes := sf.Body.RootLinearLayout.Nodes

	// for i, node := range nodes {

	// 	// We do not insert ads into the first and last index
	// 	if i == 0 {
	// 		continue
	// 	}

	// 	us := &URLSource{
	// 		Height: "wrap-content",
	// 		Width:  "match-parent",
	// 		URL:    "https://fh-uploads-addressreport.netdna-ssl.com/d79274ca-877f-437f-9ca0-e7230ea1ff46",
	// 	}

	// 	advertisement := &Advertisement{
	// 		Height:    "wrap-content",
	// 		Width:     "wrap-content",
	// 		URLSource: *us,
	// 	}

	// 	encoder := xml.NewEncoder(os.Stdout)

	// 	nodes[i] = advertisement
	// }

	return sf
}

func CountAdSlots(xmlData []byte) int {

	buf := bytes.NewBuffer(xmlData)
	dec := xml.NewDecoder(buf)

	var sf Storefront
	err := dec.Decode(&sf)
	checkError(err)

	firstLinearLayout := sf.Body.RootLinearLayout

	return countLinearLayoutElements(firstLinearLayout.Nodes)
}

func countLinearLayoutElements(nodes []Node) int {

	for i, node := range nodes {
		fmt.Printf("index: %v, element: %s\n", i, node.XMLName)
	}

	return len(nodes) - 1
}
