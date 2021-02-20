/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package server

import (
	"log"
	"net/http"
	"simple-http-broker/client"
)

func handler(w http.ResponseWriter, req *http.Request) {
	log.Println("Host		: ", req.Host)
	log.Println("URL			: ", req.URL)
	log.Println("Method		: ", req.Method)

	for reqHKey, reqHValue := range req.Header {
		log.Println("Header Key		: ", reqHKey)
		log.Println("Header Value	: ", reqHValue)
	}

	log.Println("=================================================")

	// http client request
	resp, body := client.Request(req)

	// headers
	for hKey, hValue := range resp.Header {
		for _, v := range hValue {
			w.Header().Set(hKey, v)
		}
	}

	// body
	w.Write(body)
}
