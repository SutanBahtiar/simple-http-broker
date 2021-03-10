/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is base url request
var BaseURL string

// Request is http client request
// return *http.Response and response body
func Request(req *http.Request) (response *http.Response, body []byte) {
	var client = &http.Client{}

	url := BaseURL + req.URL.Path
	request, err := http.NewRequest(req.Method, url, req.Body)

	// header
	for hKey, hValue := range req.Header {
		for _, v := range hValue {
			request.Header.Set(hKey, v)
		}
	}

	// query
	query := request.URL.Query()
	for qKey, qValue := range req.URL.Query() {
		for _, v := range qValue {
			query.Add(qKey, v)
		}
	}
	request.URL.RawQuery = query.Encode()

	if err != nil {
		panic(err)
	}

	log.Println("Request:", request.URL.String())
	log.Println("Query:", request.URL.Query())

	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, error = ioutil.ReadAll(response.Body)
	return
}
