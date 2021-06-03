/*
This package is intended to house our http handlers
*/

package handlers

import (
	"net/http"

	"github.com/wishabi/sfmlx-test/sfmlparser"
)

// GET /coasters
// the (h *coasterHandlers) is adding functionality to the struct?
func GetSFMLX(w http.ResponseWriter, r *http.Request) {

	sfml := sfmlparser.ReadSFMLFile("./samples/loblaws.sfml")
	x := sfmlparser.InjectAdsIntoSFML(sfml)

	w.Header().Add("content-type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(x)
}
