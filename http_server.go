package main

import (
	"net"
	"net/http"
)


func StartHTTPServ() string {

	// Request an available port assigned by system
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	ChkErr(err)

	go func() {

		// Close listener before quit
		defer listener.Close()
		
		panic(http.Serve(listener, http.FileServer(http.Dir("webview_ui"))))
	}()

	url := "http://" + listener.Addr().String()

	return url
}
