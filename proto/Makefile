generate:
	rm -rf ./gen/*
	protoc -I=. \
		--go_out=./gen/ --go_opt paths=source_relative \
		--go-grpc_out=./gen/ --go-grpc_opt paths=source_relative \
		--grpc-gateway_out=./gen/ --grpc-gateway_opt paths=source_relative \
		**/**/*.proto 
	protoc-go-inject-tag -input="./gen/**/**/*.pb.go"
	

lint:
	protoc --lint_out=. *.proto

install:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/favadi/protoc-go-inject-tag
	sudo apt install protobuf-compiler
