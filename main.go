package main

import (
	"fmt"

	"github.com/wishabi/sfmlx-test/sfmlparser"
)

func main() {
	fmt.Println("Hello World.")

	s := sfmlparser.ReadSFMLFile("./skeleton.xml")
	// fmt.Printf("a %s", s)
	// sfml := data.FromXML(s)
	// fmt.Println(sfml)
	// buf, _ := xml.MarshalIndent(sfml, "", "\t")
	// fmt.Println(string(buf))

	sfmlparser.CountAdSlots(s)
}
