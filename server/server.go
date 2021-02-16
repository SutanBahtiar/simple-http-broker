/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package server

import (
	"net/http"
)

// Port http server
var Port string

// Start server
func Start() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		panic(err)
	}
}
