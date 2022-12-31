package main

import (
	"bytes"
	"fmt"
	"github.com/bsm/openrtb/v3"
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func handleBidder(ctx *fasthttp.RequestCtx) {
	m := ctx.Method()
	ct := ctx.Request.Header.Peek("content-type")
	if !bytes.Equal(m, []byte("POST")) || !bytes.Equal(ct, []byte("application/json")) {
		ctx.Error("Request not understood", fasthttp.StatusBadRequest)
		return
	}
	/*
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	*/

	req := &openrtb.BidRequest{}

	if err := json.Unmarshal(ctx.PostBody(), req); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusUnsupportedMediaType)
		return
	}

	//fmt.Printf("%+v\n", &req)
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func handlePing(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("pong"))
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/ping":
		handlePing(ctx)
	case "/bidder":
		handleBidder(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}

}

func main() {
	// fmt.Printf("Running with gomaxprocs=%v\n", runtime.GOMAXPROCS(0))
	fmt.Println("NOBIDDER GO-FAST Running on http://127.0.0.1:8080/")
	if err := fasthttp.ListenAndServe("127.0.0.1:8080", fastHTTPHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
