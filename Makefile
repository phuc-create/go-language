run:
	go run main.go
build: 
	go build main.go
gen-grpc:
	protoc ./calculator.proto --go-grpc_out=.
gen-pro:
	protoc ./calculator.proto --go-out=.
