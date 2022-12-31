all:
	@echo No magic target - see README

#
# Server runners (do not exit)
# 
#
nobidder-go-vanilla:
	cd "go-vanilla" && go run .

nobidder-go-fast:
	cd "go-fast" && go run .

nobidder-rust-actix:
	cd "rust-actix" && cargo run --release

nobidder-rust-axum:
	cd "rust-axum" && cargo run --release

nobidder-python-aiohttp:
	cd "python-aiohttp" && pipenv sync && pipenv run python main.py

#
# Client runners
#
curl-ping:
	curl -s -g -i http://127.0.0.1:8080/ping

curl-bid:
	cd data/ && curl -s -g -i -H "content-type: application/json" -d @bid_request.json http://127.0.0.1:8080/bidder

wrk-ping:
	wrk --latency http://127.0.0.1:8080/ping

wrk-bid:
	cd data/ && wrk --latency -s bid_request.wrk.lua http://127.0.0.1:8080/bidder

