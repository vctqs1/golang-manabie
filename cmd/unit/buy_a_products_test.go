package test;

import (
	"testing"
	"context"
	"log"
	"time"

	_service "github.com/vctqs1/golang-manabie/pkg/services"
	"github.com/vctqs1/golang-manabie/pkg/api"
	"github.com/vctqs1/golang-manabie/database"
)



func BuyAProduct(arg []*protov1.BuyProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	db, err := database.Connect();
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	client := _service.NewProductsService(db);

	// get products
	req := protov1.BuyProductsRequest{
		Products: arg,
	}
	res, err := client.BuyProducts(ctx, &req)
	if err != nil {
		log.Printf("buy products failed: <%+v>\n\n", err)
	} else {
		log.Printf("buy products result: <%+v>\n\n", res)
	}
	return err
}




func TestBuyAProducts(t *testing.T) {
	err := BuyAProduct([]*protov1.BuyProduct{
		{
			ProductId: 6,
			Quantities: 1,
		}, 
	})
	if err != nil {
		t.Error(err)
	}
}


func TestBuyInvalidAProducts(t *testing.T) {
	err := BuyAProduct([]*protov1.BuyProduct{{
		ProductId: 6,
		Quantities: 111,
		}, 
	})
	if err != nil {
		t.Error(err)
	}
}