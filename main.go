/*
 * Created on Tue Feb 16 2021
 *
 * Copyright (c) 2021 http://github.com/SutanBahtiar/simple-http-broker
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-http-broker/client"
	"simple-http-broker/server"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Input port to use in http server")
	fmt.Println("Example `8080`")
	fmt.Print("Port Server :")
	port, _ := r.ReadString('\n')

	fmt.Println("Input base url to use in http client")
	fmt.Println("Example `http://github.com`")
	fmt.Print("URL Client :")
	url, _ := r.ReadString('\n')

	server.Port = ":" + strings.Trim(port, "\n")
	client.BaseURL = strings.Trim(url, "\n")

	fmt.Println("=====================================================================")
	fmt.Println("Http Server started on port", server.Port)
	fmt.Println("Http Client base url", client.BaseURL)
	fmt.Println("=====================================================================")

	server.Start()
}
