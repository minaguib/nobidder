all: openrtb.pb.go

openrtb.pb.go: openrtb.proto
	protoc --go_out=. $<
