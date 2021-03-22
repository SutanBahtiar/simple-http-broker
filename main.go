/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package main

import (
	"flag"
	"fmt"
	"simple-http-broker/client"
	"simple-http-broker/server"
)

func main() {
	port := flag.Int64("port", 8080, "port for http server(default 8080)")
	url := flag.String("url", "http://localhost", "base url for http client/e.g http://github.com")
	flag.Parse()
	server.Port = ":" + fmt.Sprint(*port)
	client.BaseURL = *url

	fmt.Println("=====================================================================")
	fmt.Println("Http Server started on port", server.Port)
	fmt.Println("Http Client base url", client.BaseURL)
	fmt.Println("=====================================================================")

	server.Start()
}
