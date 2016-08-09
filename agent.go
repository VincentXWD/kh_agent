package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"io"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	res, err := http.DefaultClient.Do(r)
	defer res.Body.Close()
	if err != nil {	log.Panicln(err.Error()) }
	for k, v := range res.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	for _, c := range res.Cookies() {
		w.Header().Add("Set-Cookie", c.Raw)
	}
	w.WriteHeader(res.StatusCode)
	result, err := ioutil.ReadAll(res.Body)
	if err != nil && err != io.EOF { log.Panicln(err.Error()) }
	w.Write(result)
}

func main() {
	http.HandleFunc("/", Handler)
	log.Infoln("Starting serving on port ", os.Args[1])
	http.ListenAndServe(":"+os.Args[1], nil)
}