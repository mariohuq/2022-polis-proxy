package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func director(req *http.Request) {
	log.Println("Request:")
	req.Write(log.Writer())
}

func modifyResponse(res *http.Response) error {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err
	}
	log.Println("Response:")
	log.Writer().Write(dump)
	return nil
}

func newServer(addr string) http.Server {
	return http.Server{
		Handler: &httputil.ReverseProxy{
			Director:       director,
			ModifyResponse: modifyResponse,
		},
		Addr: addr,
	}
}

func main() {
	log.Println("Launching server...")
	proxy := newServer(":8080")
	log.Println(proxy.ListenAndServe())
}
