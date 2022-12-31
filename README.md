# nobidder

Experiments in writing a no-bid openrtb http server

## Implementations

* `go-vanilla` - Go vanilla implementation with standard lib HTTP server and JSON parser
* `go-fast` - Go implementatin with faster 3rd-party HTTP and JSON libraries
* `rust-actix` - Rust implementation on the actix web framework
* `rust-axum` - Rust implementation on the axum web framework

## Behavior

Each server listens on http://127.0.0.1/ and supports these paths:

* `GET /ping` - self-explanatory
* `POST /bidder` - Accepts and parses OpenRTB JSON and returns a no-bid response (HTTP 204)

## Other

* `data` - Supporting test payloads
