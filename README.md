sh ./scripts/proto-gen.sh 

#server
cd cmd/server
go build .
./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password= -db-schema=golang_manabie


#client
cd cmd/client
go build .
./client -server=localhost:9090 -id1=1 -id2=2 -quantities1=1 -quantities2=2


protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  path/to/your_service.proto



protoc --proto_path=api --proto_path=scripts --go_out=plugins=grpc:pkg/api ./api/*.proto


protoc -I/usr/local/include -I./api/ \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
 --grpc-gateway_out=logtostderr=true:./pkg/api \
  ./api/*.proto

protoc -I/usr/local/include -I./api/ \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:./pkg/api \
  ./api/*.proto

  protoc -I/usr/local/include -I./api/ \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:./pkg/api \
  ./api/*.proto