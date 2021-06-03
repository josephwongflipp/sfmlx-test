/*
This package is responsible for all CRUD actions against an SFML file.

Actions include:
- Counting the number of valid advertisement slots in an SFML document
- Injecting advertisements in an SFML document
- Returning valid SFML
*/

package sfmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/beevik/etree"
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
// func injectAds(xmlData []byte, maxAds int, offset int) []byte {
// 	buf := bytes.NewBuffer(xmlData)
// 	dec := xml.NewDecoder(buf)

// 	var sf data.Storefront
// 	err := dec.Decode(&sf)
// 	checkError(err)

// 	for i, node := range nodes {

// 		// We do not insert ads into the first and last index
// 		if i == 0 {
// 			continue
// 		}

// 		us := &URLSource{
// 			Height: "wrap-content",
// 			Width:  "match-parent",
// 			URL:    "https://fh-uploads-addressreport.netdna-ssl.com/d79274ca-877f-437f-9ca0-e7230ea1ff46",
// 		}

// 		advertisement := &Advertisement{
// 			Height:    "wrap-content",
// 			Width:     "wrap-content",
// 			URLSource: *us,
// 		}

// 		// panic idk
// 		adNode := &Node{
// 			// TODO: shove in advertisement info here
// 		}

// 		// TODO: insert adNode into nodes and increment i

// 		encoder := xml.NewEncoder(os.Stdout)

// 		nodes[i] = advertisement
// 	}

// 	return sf
// }

// func CountAdSlots(xmlData []byte) int {

// 	buf := bytes.NewBuffer(xmlData)
// 	dec := xml.NewDecoder(buf)

// 	var sf Storefront
// 	err := dec.Decode(&sf)
// 	checkError(err)

// 	firstLinearLayout := sf.Body.RootLinearLayout

// 	return countLinearLayoutElements(firstLinearLayout.Nodes)
// }

func countLinearLayoutElements(nodes []Node) int {

	for i, node := range nodes {
		fmt.Printf("index: %v, element: %s\n", i, node.XMLName)
	}

	return len(nodes) - 1
}

func PraseEtree(d []byte) {
	doc := etree.NewDocument()
	doc.ReadFromBytes(d)
	root := doc.SelectElement("storefront")
	fmt.Println("STOREFRONT element:", root.Tag)

	body := root.SelectElement("body")
	fmt.Println("BODY element:", body.Tag)

	rootLL := body.SelectElement("linear-layout")
	fmt.Println("ROOT LL element:", rootLL.Tag)

	for _, element := range rootLL.ChildElements() {
		fmt.Println("CHILD element:", element.Tag)
	}

	// for _, book := range root.SelectElements("book") {
	// 	fmt.Println("CHILD element:", book.Tag)
	// 	if title := book.SelectElement("title"); title != nil {
	// 		lang := title.SelectAttrValue("lang", "unknown")
	// 		fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
	// 	}
	// 	for _, attr := range book.Attr {
	// 		fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
	// 	}
	// }
}
