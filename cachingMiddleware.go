package main

//I'm relearning go and used this as an excuse to write custom middleware.  it's not used.

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type MyResponseWriter struct {
	http.ResponseWriter
	buf        *bytes.Buffer
	statusCode int
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	mrw.buf.Write(p)
	return mrw.ResponseWriter.Write(p)
}

func (mrw *MyResponseWriter) WriteHeader(statusCode int) {
	fmt.Println(statusCode)
	mrw.statusCode = statusCode
	mrw.ResponseWriter.WriteHeader(statusCode)
}

func withCaching(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware hit")
		log.Println(r.URL.Path)

		mrw := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}

		next.ServeHTTP(mrw, r)

		// response, err := ioutil.ReadAll(mrw.buf)
		// if err != nil {
		// 	log.Printf("Error reading body: %v", err)
		// 	return
		// }

		log.Println("could cache this:")
		// log.Println(string(response))
		// log.Println(mrw.statusCode)

		// if _, err := io.Copy(w, mrw.buf); err != nil {
		// 	log.Printf("Failed to send out response: %v", err)
		// }

	})
}
