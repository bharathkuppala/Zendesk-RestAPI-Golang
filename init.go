package main

import (
	"flag"

	"log"
)

const (
	secretKey   = ""
	redirectURI = "https://843f46f4.ngrok.io/zendesk/oauth"
	clientID    = "zendesk_oauth"
	subDomain   = "https://celebal.zendesk.com"
)

func init() {
	log.Println("ZenDesk Implementation Starts ....")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.String("port", "", "port to listen on")

	flag.Parse()

	if flag.Lookup("port").Value.String() == "" {
		log.Fatal("-port is required")
	}

	// request, err := http.NewRequest("GET", "https://"+subDomain+"/oauth/authorizations/new?response_type=code&redirect_uri="+redirectURI+"&client_id="+clientID+"&scope=read%20write", nil)
	zendesk()
	// readcsv()
}
