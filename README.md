# How to run project


## install
```
make install
````

## database
copy script /database to create database

## generate proto
```
sh ./scripts/proto-gen.sh
```


## how to run server
go to cmd/sever
```
cd cmd/server
```

build 
```
go build .
```

start - params is optional
```
./server -grpc-port=9090
```


## how to run client
go to cmd/client
```
cd cmd/client
```

build 
```
go build .
```

start - params is optional
```
./client -server=localhost:9090 -id1=1 -id2=2 -quantities1=1 -quantities2=2 -invalidid=7 -invalidquantities=111
```


## test
```
cmd/test
go test .
```