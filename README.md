sh ./scripts/proto-gen.sh 

#server
cd cmd/server
go build .
./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password= -db-schema=golang_manabie


#test
cd cmd/test
go build .
./test -server=localhost:9090


protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  path/to/your_service.proto



protoc --proto_path=api --proto_path=scripts --go_out=plugins=grpc:pkg api product.proto
protoc --proto_path=api --proto_path=scripts--grpc-gateway_out=logtostderr=true:pkg/api product.proto


protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:pkg/api \
  api/*.proto
