package main

import (
	"fmt"
	"context"
	"flag"
	"log"
	"time"
	"sync"

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
	} else {
		log.Printf("get products result: <%+v>\n\n", res)
	}
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
	} else {
		log.Printf("buy products result: <%+v>\n\n", res)
	}
	return err
}



type ProductResponse struct {
	Res *protov1.BuyProductsResponse
	Err error
}
func BuyAProductOfConcurent(address string, arg1, arg2 []*protov1.BuyProduct) []ProductResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}
	defer conn.Close()

	client := protov1.NewProductsServiceClient(conn)


	var wg sync.WaitGroup

	wg.Add(2)

	res := make([]ProductResponse, 0)
	

	go func() {
		res1, err1 := client.BuyProducts(ctx, &protov1.BuyProductsRequest{
			Products: arg1,
		});
		res = append(res, ProductResponse{res1, err1})
		wg.Done()

	}()

	go func() {
		res2, err2 := client.BuyProducts(ctx, &protov1.BuyProductsRequest{
			Products: arg2,
		});
		res = append(res, ProductResponse{res2, err2})
		wg.Done()

	}()
	wg.Wait()
	return res

}


func BuyConcurentProductRoutine(address string, arg1, arg2 []*protov1.BuyProduct) (bool, error) {


	responses := BuyAProductOfConcurent(address, arg1, arg2)

	success := 0;
	res := make([] ProductResponse, 0, 2)
	message := make([] error, 0, 2)

	for _, value := range responses {

		if value.Res != nil && value.Res.Successful == true {
			success = success + 1
		} else if value.Err != nil {
			message = append(message, value.Err)

		}

		res = append(res, value)
	}

	if success > 0 {
		if success == 2 {
			return true, nil;

		} else {
			return true, fmt.Errorf("%+v", message);

		}
	} else {
		return false, fmt.Errorf("%+v", message);
	}
}

func Ex3(address string, arg1, arg2 []*protov1.BuyProduct) error {

	_, err := BuyConcurentProductRoutine(address, arg1, arg2)
	return err;
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
	err = Ex3(*address, []*protov1.BuyProduct{
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
	fmt.Printf("Buy concurent valid mess: %v\n", err)
	//example 4
	err = Ex3(*address, []*protov1.BuyProduct{
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
	fmt.Printf("Buy concurent invalid mess: %v\n", err)

	
}
