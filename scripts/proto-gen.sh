



protoc --proto_path=api --proto_path=scripts --go_out=plugins=grpc:pkg/api api/*.proto


protoc -I/usr/local/include -I./api/ \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
 --grpc-gateway_out=logtostderr=true:./pkg/api \
  ./api/*.proto