package main

import (
	"encoding/xml"

	"github.com/wishabi/sfmlx-test/data"
	"github.com/wishabi/sfmlx-test/sfmlparser"
)

func main() {
	sfml := sfmlparser.ReadSFMLFile("./samples/loblaws.xml")

	mySFML := &data.Storefront{}

	err := xml.Unmarshal(sfml, mySFML)
	if err != nil {
		panic(err)
	}

	// fmt.Println(mySFML)
	// fmt.Printf("%+v\n", mySFML)

	// myXml, err := xml.Marshal(mySFML)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(myXml)
	// fmt.Printf(string(myXml))

	sfmlparser.PraseEtree(sfml)
}
