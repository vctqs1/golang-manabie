sh ./scripts/proto-gen.sh 
./server -grpc-port=9090 -db-host=localhost:3306 -db-user=root -db-password= -db-schema=golang_manabie

./test -server=localhost:9090