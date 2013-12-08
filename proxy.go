package main

import (
        "flag"
        "net/http"
        "net/http/httputil"
        "net/url"
        "time"
        "log"
)


// Command line usage:
// 
//     reversal <src> <dst>
//
// For exampke, here the command to proxy localhost:4000 
// to localhost:8000:
//
//     reversal :4000 http://localhost:8000
//
func main() {
        flag.Parse()
        args := flag.Args()

        if len(args) != 2 {
		log.Fatal("Need both destination and src")
	}

	source := args[0]
        url, err := url.Parse(args[1])

        if err != nil {
                log.Fatal("Bad destination.")
        }

        h := httputil.NewSingleHostReverseProxy(url)
        s := &http.Server{
                Addr:           source,
                Handler:        h,
                ReadTimeout:    10 * time.Second,
                WriteTimeout:   10 * time.Second,
                MaxHeaderBytes: 1 << 20,
        }

	log.Printf("Listening on %s", source)
        log.Fatal(s.ListenAndServe())
}

