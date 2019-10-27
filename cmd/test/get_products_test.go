package test;

import (
	"testing"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"github.com/vctqs1/golang-manabie/pkg/api"
	"github.com/vctqs1/golang-manabie/database"
)



func GetProducts(id []int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := database.GetConfig()

	conn, err := grpc.Dial(":" + cfg.GRPCPort, grpc.WithInsecure())
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


func TestGetProducts(t *testing.T) {
	err := GetProducts([]int64 {1, 2})
	if err != nil {
		t.Error(err)
	}
}