.EXPORT_ALL_VARIABLES:

install: install-go install-db
install-go:
	GO111MODULE=on go get \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
install-db: 
	docker-compose --file ./docker-compose.yml up -d