package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HTTP API wrapper
type HTTPWrapper struct{}

// API method
func (_ HTTPWrapper) Method(path string, handler func(*Ctx)) {
	http.HandleFunc(path, func(rw http.ResponseWriter, req *http.Request) {
		handler(HTTPWrapperCtx(rw))
	})
}

// Serve on port
func (_ HTTPWrapper) Serve(port int) {
	log.Print("Listening on http://localhost:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

// fill API context
func HTTPWrapperCtx(rw http.ResponseWriter) *Ctx {
	return &Ctx{
		func(str string) {
			rw.Write([]byte(str))
		},

		func(bts []byte) {
			rw.Write(bts)
		},

		func(data interface{}) {
			bts, err := json.Marshal(data)
			if err != nil {
				log.Fatal("Cannot marshal JSON:", data)
			}
			rw.Write(bts)
		},
	}
}
