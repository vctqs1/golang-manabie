package main

import (
	"fmt"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/vctqs1/golang-manabie/pkg/api"
)

func Ex1(address string, id []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()

	client := protov1.NewProductsServiceClient(conn)

	// get products
	req := protov1.GetProductsRequest{
		ProductIds: id,
	}
	res, err := client.GetProducts(ctx, &req)
	if err != nil {
		log.Printf("get products failed: <%+v>\n\n", err)
	}
	log.Printf("get products result: <%+v>\n\n", res)
	return err
}


func Ex2(address string, arg []*protov1.BuyProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()

	client := protov1.NewProductsServiceClient(conn)

	// get products
	req := protov1.BuyProductsRequest{
		Products: arg,
	}
	res, err := client.BuyProducts(ctx, &req)
	if err != nil {
		log.Printf("buy products failed: <%+v>\n\n", err)
	}
	log.Printf("buy products result: <%+v>\n\n", res)
	return err
}




func Ex3(address string, arg1, arg2 []*protov1.BuyProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()

	client := protov1.NewProductsServiceClient(conn)

	// get products
	
	req := []*protov1.BuyProductsRequest{{
		Products: arg1,
	}, {
		Products: arg2,
	}}
	var message error
	for _, value := range req {
		res, err := client.BuyProducts(ctx, value)
		if err != nil {
			log.Printf("call 2: failed: <%+v>\n\n", err)
			message = err
		}
		log.Printf("call 2: buy products:  <%+v>\n\n", res)
	}

	return message;
}


func main() {

	address := flag.String("server", "localhost:9090", "gRPC server in format host:port")
	id1 := flag.Int64("id1", 1, "valid id 1")
	id2 := flag.Int64("id2", 2, "valid id 2")
	quantities1 := flag.Int64("quantities1", 1, "quantities of valid id 1")
	quantities2 := flag.Int64("quantities2", 2, "quantities of valid id 2")
	invalidid := flag.Int64("invalidid", 7, "invalidid")
	invalidquantities := flag.Int64("invalidquantities", 200, "invalidquantities")

	flag.Parse();

	fmt.Printf("ID1: %d, Quantities1: %d\n", *id1, *quantities1)
	fmt.Printf("ID2: %d, Quantities2: %d\n", *id2, *quantities2)
	fmt.Printf("invalidid: %d, invalidquantities: %d\n", *invalidid, *invalidquantities)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()


	// get products
	Ex1(*address, []int64{*id1, *id2})


	

	// example 1: buy products: valid quantities
	Ex2(*address, []*protov1.BuyProduct{
		{
			ProductId: *id1,
			Quantities: *quantities1,
		}, 
	})

	
	// example 2: buy products: invalid quantities
	Ex2(*address, []*protov1.BuyProduct{
		{
			ProductId: *invalidid,
			Quantities: *invalidquantities,
		}, 
	})
	
	//example 3
	Ex3(*address, []*protov1.BuyProduct{
		{
			ProductId: *id1,
			Quantities: *quantities1,
		}, 
	}, []*protov1.BuyProduct{
		{
			ProductId: *id2,
			Quantities: *quantities2,
		}, 
	});

	//example 4
	Ex3(*address, []*protov1.BuyProduct{
		{
			ProductId: *id1,
			Quantities: *quantities1,
		}, 
	}, []*protov1.BuyProduct{
		{
			ProductId: *invalidid,
			Quantities: *invalidquantities,
		}, 
	});

	
}
