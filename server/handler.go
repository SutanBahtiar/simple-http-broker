/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package server

import (
	"fmt"
	"net/http"
	"simple-http-broker/client"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("====================== start ========================================")
	fmt.Println("Host			: ", req.Host)
	fmt.Println("URL			: ", req.URL)
	fmt.Println("Method			: ", req.Method)

	for reqHKey, reqHValue := range req.Header {
		fmt.Println("Header Key		: ", reqHKey)
		fmt.Println("Header Value		: ", reqHValue)
	}

	fmt.Println("======================  end  ========================================")

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
