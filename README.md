sh ./scripts/proto-gen.sh 

#server
cd cmd/server
go build .
./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password= -db-schema=golang_manabie


#test
cd cmd/test
go build .
./test -server=localhost:9090