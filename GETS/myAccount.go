package GETS

import (
	"fmt"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusGatewayTimeout)
	fmt.Fprintln(rw, "Home")
}
