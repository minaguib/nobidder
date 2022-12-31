package main

import (
	"encoding/json"
	"fmt"
	"github.com/bsm/openrtb/v3"
	"log"
	"net/http"
	"time"
)

func handleBidder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Request not understood", http.StatusBadRequest)
		return
	}
	/*
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	*/

	var req openrtb.BidRequest

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	// fmt.Printf("%+v\n", &req)

	w.WriteHeader(http.StatusNoContent)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {

	fmt.Println("NOBIDDER GO Running on http://127.0.0.1:8080/")
	//fmt.Printf("Running with gomaxprocs=%v\n", runtime.GOMAXPROCS(0))
	mux := http.NewServeMux()
	mux.HandleFunc("/bidder", handleBidder)
	mux.HandleFunc("/ping", handlePing)

	server := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 4 * 1024,
	}

	log.Fatal(server.ListenAndServe())
}
