package handler

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>" + quote.Go())
}