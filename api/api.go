package api

import (
	"fmt"
	"net/http"

	"github.com/albantanie/mahfudzot-generator/service"
)

type API struct {
	mux *http.ServeMux
	service.MahfudzotService
}

func NewAPI(mahfudzotService service.MahfudzotService) *API {
	mux := http.NewServeMux()
	api := &API{
		mux: mux,
	}
	api.MahfudzotService = mahfudzotService

	// API version prefix
	v1 := "/api/v1"

	mux.Handle(v1+"/mahfudzot", http.HandlerFunc(api.MahfudzotHandler))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8081")
	http.ListenAndServe(":8081", api.mux)
}
