package main

import (
	"context"
	"flag"
	"log"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	"github.com/vctqs1/golang-manabie/pkg/api"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()

	c := protov1.NewProductsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get products
	req1 := protov1.GetProductsRequest{
		ProductIds: []int64{},
	}
	res1, err := c.GetProducts(ctx, &req1)
	if err != nil {
		log.Printf("get products failed: <%+v>\n\n", err)
	}
	log.Printf("get products result: <%+v>\n\n", res1)

	// example 1: buy products: valid quantities
	req2 := protov1.BuyProductsRequest{
		Products: []*protov1.BuyProduct{
			{
				ProductId:  1,
				Quantities: 2,
			}, {
				ProductId:  2,
				Quantities: 6,
			},
		},
	}
	res2, err := c.BuyProducts(ctx, &req2)
	if err != nil {
		log.Printf("buy products: valid quantities failed: <%+v>\n\n", err)
	}
	log.Printf("buy products: valid quantities result: <%+v>\n\n", res2)

	// example 2: buy products: invalid quantities
	req3 := protov1.BuyProductsRequest{
		Products: []*protov1.BuyProduct{
			{
				ProductId:  1,
				Quantities: 2,
			}, {
				ProductId:  2,
				Quantities: 6,
			},
		},
	}
	res3, err := c.BuyProducts(ctx, &req3)
	if err != nil {
		log.Printf("buy products: invalid quantities failed: <%+v>\n\n", err)
	}
	log.Printf("buy products: invalid quantities result: <%+v>\n\n", res3)

	//example 3
	req4 := []*protov1.BuyProductsRequest{{
		Products: []*protov1.BuyProduct{
			{
				ProductId:  4,
				Quantities: 1,
			}, {
				ProductId:  5,
				Quantities: 1,
			},
		},
	}, {
		Products: []*protov1.BuyProduct{
			{
				ProductId:  4,
				Quantities: 1,
			}, {
				ProductId:  5,
				Quantities: 2,
			},
		},
	}}

	for _, value := range req4 {
		res, err := c.BuyProducts(ctx, value)
		if err != nil {
			log.Printf("Example 3: buy products: invalid quantities failed: <%+v>\n\n", err)
		}
		log.Printf("Example 3: buy products:  <%+v>\n\n", res)
	}

	//example 4
	req5 := []*protov1.BuyProductsRequest{{
		Products: []*protov1.BuyProduct{
			{
				ProductId:  6,
				Quantities: 1,
			}, {
				ProductId:  7,
				Quantities: 1,
			},
		},
	}, {
		Products: []*protov1.BuyProduct{
			{
				ProductId:  6,
				Quantities: 1,
			}, {
				ProductId:  7,
				Quantities: 2,
			},
		},
	}}

	for _, value := range req5 {
		res, err := c.BuyProducts(ctx, value)
		if err != nil {
			log.Printf("Example 4: buy products: invalid quantities failed: <%+v>\n\n", err)
		}
		log.Printf("Example 4: buy products:  <%+v>\n\n", res)
	}

}
