
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
		log.Fatalf("did not connect: %v", err)
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
		log.Fatalf("get products failed: %v", err)
	}
	log.Printf("get products result: <%+v>\n\n", res1)



	// buy products
	req2 := protov1.BuyProductsRequest{
		Products: []*protov1.BuyProduct{{
				ProductId: 1,
				Quantities: 2,	
			},
		},
	}
	res2, err := c.BuyProducts(ctx, &req2)
	if err != nil {
		log.Fatalf("buy products failed: %v", err)
	}
	log.Printf("buy products result: <%+v>\n\n", res2)

}