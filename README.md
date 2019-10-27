#database
copy script /database to create database

#generate proto
sh ./scripts/proto-gen.sh 


#how to run server port 9090
cd cmd/server
go build .
#start - params is optional
./server -db-host=localhost:3306 -db-user=root -db-password= -db-schema=golang_manabie


#how to run client
cd cmd/client
go build .
#start - params is optional
./client -id1=1 -id2=2 -quantities1=1 -quantities2=2 -invaliid=7 -invalidquantities=111


#test
cmd/test
go test .