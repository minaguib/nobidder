local f = io.open("bid_request.json", "rb")
wrk.method = "POST"
wrk.body   = f:read("*all")
wrk.headers["Content-Type"] = "application/json"
