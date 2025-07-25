build:
	go build -o bin/main main.go

# Run the application
run:
	go run main.go

gen-proto:
	protoc --go_out=. --go-grpc_out=. proto/user/user.proto