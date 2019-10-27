package test;

import (
	"testing"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/vctqs1/golang-manabie/pkg/api"
	"fmt"

	"github.com/vctqs1/golang-manabie/database"
)



func BuyTwoProduct(arg1, arg2 []*protov1.BuyProduct) (*protov1.BuyProductsResponse, error) {
	cfg := database.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.Dial(":" + cfg.GRPCPort, grpc.WithInsecure())
	if err != nil {
		return &protov1.BuyProductsResponse{
			Successful: false,
		}, err;
	}
	defer conn.Close()

	client := protov1.NewProductsServiceClient(conn)

	// buy products
	
	req := []*protov1.BuyProductsRequest{{
		Products: arg1,
	}, {
		Products: arg2,
	}}
	var message []error;
	var result []*protov1.BuyProductsResponse;
	for _, value := range req {
		res, err := client.BuyProducts(ctx, value)

		if err != nil {
			message = append(message, err)
			log.Printf("call 2: failed: <%+v>\n\n", err)
		} else {
			result = append(result, res)
			log.Printf("call 2: buy products:  <%+v>\n\n", res)
		}
	}
	if len(result) != 0 {
		return &protov1.BuyProductsResponse{
			Successful: false,
		}, status.Error(codes.Unknown, fmt.Sprintf("buy fail %+v", message));
	}
	return &protov1.BuyProductsResponse{
		Successful: true,
	}, nil;
}


func TestBuyTwoProducts(t *testing.T) {
	_, err := BuyTwoProduct([]*protov1.BuyProduct{
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
	_, err := BuyTwoProduct([]*protov1.BuyProduct{{
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