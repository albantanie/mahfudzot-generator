package api

import (
	"fmt"
	"net/http"
)

func (api *API) MahfudzotHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Mahfudzot!")
}
