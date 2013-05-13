// webServer
package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	buf, err := ioutil.ReadFile("website/website.html")
	if err != nil {
		return
	}

	io.WriteString(
		res,
		string(buf),
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe("localhost:4000", nil)
}
