package infrastructure

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang-backend-service/src/logging"
	"net/http"
)

type httpRouter struct {
	Router    mux.Router
	certFile  string
	keyFile   string
	addrHttp  string
	addrHttps string
}

func NewHttpRouter(addrHttp, addrHttps, certFile, keyFile string) *httpRouter {

	r := mux.NewRouter().StrictSlash(true)

	return &httpRouter{
		Router:    *r,
		addrHttp:  addrHttp,
		addrHttps: addrHttps,
		certFile:  certFile,
		keyFile:   keyFile,
	}
}

func (hr httpRouter) HandleRequests() chan error {

	errs := make(chan error)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: logging.GetActLevel() == logging.DEBUG,
	})

	// Starting HTTP server
	go func() {
		logging.Info.Printf("Staring HTTP service on %s ...", hr.addrHttp)

		if err := http.ListenAndServe(hr.addrHttp, c.Handler(&hr.Router)); err != nil {
			errs <- err
		}
	}()

	// Starting HTTPS server
	go func() {
		logging.Info.Printf("Staring HTTPS service on %s ...", hr.addrHttps)
		if err := http.ListenAndServeTLS(hr.addrHttps, hr.certFile, hr.keyFile, c.Handler(&hr.Router)); err != nil {
			errs <- err
		}
	}()

	return errs
}
