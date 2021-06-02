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

// Generic node
type Node struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []Node `xml:",any"`
}

// STRUCTS begin

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
	XMLName xml.Name `xml:"linear-layout"`
	Height  string   `xml:"height,attr"`
	Width   string   `xml:"width,attr"`
	Content []byte   `xml:",innerxml"`
	Nodes   []Node   `xml:",any"`
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

// STRUCTS end

func CountAdSlots(xmlData []byte) int {

	buf := bytes.NewBuffer(xmlData)
	dec := xml.NewDecoder(buf)

	var sf Storefront
	err := dec.Decode(&sf)
	checkError(err)

	firstLinearLayout := sf.Body.RootLinearLayout

	return countLinearLayoutElements(firstLinearLayout.Nodes)

	// walk(firstLinearLayout.Nodes, func(n Node) bool {
	// 	fmt.Println("Checking element: ", n.XMLName.Local)
	// 	if n.XMLName.Local == "text" {
	// 		fmt.Println(string(n.Content))
	// 	}
	// })
}

// func walk(ll []Node, f func(Node) bool) {
// 	for _, n := range ll {
// 		if f(n) {
// 			walk(n.Nodes, f)
// 		}
// 	}
// }

func countLinearLayoutElements(nodes []Node) int {

	for i, node := range nodes {
		fmt.Printf("index: %v, element: %s\n", i, node.XMLName)
	}

	return len(nodes) - 1
}
