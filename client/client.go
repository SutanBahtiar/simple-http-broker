/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package client

import (
	"io/ioutil"
	"net/http"
)

// BaseURL is base url request
var BaseURL string

// Request is http client request
// return *http.Response and response body
func Request(req *http.Request) (resp *http.Response, body []byte) {
	resp, err := http.Get(BaseURL + req.URL.Path)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
