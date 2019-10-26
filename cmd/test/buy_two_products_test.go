package test;

import (
	"testing"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"github.com/vctqs1/golang-manabie/pkg/api"
)



func BuyTwoProduct(address string, arg1, arg2 []*protov1.BuyProduct) error {
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
		} else {
			log.Printf("call 2: buy products:  <%+v>\n\n", res)
		}
	}

	return message;
}


func TestBuyTwoProducts(t *testing.T) {
	err := BuyTwoProduct(":9090", []*protov1.BuyProduct{
		{
			ProductId: 6,
			Quantities: 1,
		}, 
	}, []*protov1.BuyProduct{
		{
			ProductId: 6,
			Quantities: 1,
		}, 
	})
	if err != nil {
		t.Error(err)
	}
}


func TestBuyInvalidATwoProducts(t *testing.T) {
	err := BuyTwoProduct(":9090", []*protov1.BuyProduct{{
		ProductId: 6,
		Quantities: 1,
		}, 
	}, []*protov1.BuyProduct{{
		ProductId: 6,
		Quantities: 111,
		}, 
	})
	if err != nil {
		t.Error(err)
	}
}