package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bsm/openrtb"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"github.com/minaguib/nobidder/pb"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var protoBufNoBid = []byte{0x20, 0x01}

func handleOpenRTB(w http.ResponseWriter, r *http.Request) {

	/*
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	*/

	var req *openrtb.BidRequest

	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	} else if len(in) == 0 {
		http.Error(w, "empty body", http.StatusUnsupportedMediaType)
		return
	}

	if err := json.Unmarshal(in, &req); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func handleProtoBuf(w http.ResponseWriter, r *http.Request) {

	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	} else if len(in) == 0 {
		http.Error(w, "empty body", http.StatusUnsupportedMediaType)
		return
	}

	br := &com_google_openrtb.BidRequest{}
	err = proto.Unmarshal(in, br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	w.Header().Set("content-type", "application/octet-stream")
	w.Write(protoBufNoBid)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.Header.Get("content-type") == "application/json" {
		handleOpenRTB(w, r)
	} else if r.Method == "POST" && r.Header.Get("content-type") == "application/octet-stream" {
		handleProtoBuf(w, r)
	} else {
		http.Error(w, "Request not understood", http.StatusBadRequest)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 4 * 1024,
	}

	log.Fatal(server.ListenAndServe())
}
