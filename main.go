package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("zendesk Oauth"))
		// http.Redirect(w, r, subDomain+"/oauth/authorizations/new", http.StatusSeeOther)
	})

	// router.HandleFunc("/zendesk/oauth", zendesk).Methods("POST")

	flag.VisitAll(func(flag *flag.Flag) {
		log.Println(flag.Name, "->", flag.Value)
	})

	server := &http.Server{
		Addr:    ":" + flag.Lookup("port").Value.String(),
		Handler: cors.Default().Handler(gziphandler.GzipHandler(noCacheMW(router))),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
func noCacheMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")

		h.ServeHTTP(w, r)
	})
}
