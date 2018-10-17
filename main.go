/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Additions - (C) 2018 Hans Donner
 * - response includes some ENV output
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n\n", host)

	fmt.Fprintf(w, "For %s %s\n", r.Host, r.URL.Path)

	for _, cookie := range r.Cookies() {
        fmt.Fprint(w, "%s: %s", cookie.Name, cookie.Value)
    }
	fmt.Fprintf(w, "\n")

	// add info from HELLO_ environment variables
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "HELLO_") {
			pair := strings.Split(e, "="); 
			fmt.Fprintf(w, "%s: %s\n", pair[0], pair[1])
		}
    }
}