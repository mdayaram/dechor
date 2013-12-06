package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Pony struct{}

func (p *Pony) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	status := strings.Split(strings.TrimPrefix(req.URL.Path, "/"), "/")[0]
	queries := req.URL.Query()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// don't care, just set body to empty string.
		body = []byte("")
	}

	code := 200
	if len(status) != 0 {
		code, err = strconv.Atoi(status)
		if err != nil {
			code = 200
		}
	}

	// And now override with what we got from the request.
	for k, vs := range queries {
		rw.Header()[k] = vs
	}
	rw.WriteHeader(code)
	rw.Write(body)

}

func main() {
	addr := os.Getenv("PORT")
	if len(addr) == 0 {
		addr = "8080"
	}
	http.ListenAndServe(":"+addr, &Pony{})
}
