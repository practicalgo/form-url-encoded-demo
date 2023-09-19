package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handleForm(w http.ResponseWriter, req *http.Request) {
	log.Println("handling request for /submit-form")
	log.Println("Content-Type", req.Header.Get("Content-Type"))

	// this will look for application/x-www-form-urlencoded
	// content type
	err := req.ParseForm()
	if err != nil {
		// if there was an error here, we return the error
		// as response along with a 400 http response code
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// once the form has been parsed correctly
	// we can obtain the form fields  using
	// req.FormValue passing the field name as the key
	// see ./post-method.html for the field names

	// it's worth noting here that FormValue() will return
	// the first value if there are multiple values specified
	// in the request
	// if you want to access the multiple values specified,
	// use req.Form directly (see https://pkg.go.dev/net/http#Request)
	say := req.FormValue("say")
	to := req.FormValue("to")

	fmt.Fprintf(w, fmt.Sprintf("%s %s", strings.ToUpper(say), strings.ToUpper(to)))
}
