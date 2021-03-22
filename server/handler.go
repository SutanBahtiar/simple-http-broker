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
	"simple-http-broker/util"
)

func handler(w http.ResponseWriter, req *http.Request) {
	logId := util.LogId()
	log.Printf("%s, Host:%s", logId, req.Host)
	log.Printf("%s, URL:%s", logId, req.URL)
	log.Printf("%s, Query:%s", logId, req.URL.Query())
	log.Printf("%s, Method:%s", logId, req.Method)

	for reqHKey, reqHValue := range req.Header {
		log.Printf("%s, Header Key:%s", logId, reqHKey)
		log.Printf("%s, Header Value:%s", logId, reqHValue)
	}

	// http client request
	resp, body := client.Request(&logId, req)

	// headers
	for hKey, hValue := range resp.Header {
		for _, v := range hValue {
			w.Header().Set(hKey, v)
		}
	}

	// body
	w.Write(body)
}
