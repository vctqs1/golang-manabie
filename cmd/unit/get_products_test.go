package test;

import (
	"testing"
	"context"
	"log"
	"time"

	"github.com/vctqs1/golang-manabie/pkg/api"
	_service "github.com/vctqs1/golang-manabie/pkg/services"
	"github.com/vctqs1/golang-manabie/database"
)



func GetProducts(id []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.Connect();
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	client := _service.NewProductsService(db);
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


func TestGetProducts(t *testing.T) {
	err := GetProducts([]int64 {1, 2})
	if err != nil {
		t.Error(err)
	}
}