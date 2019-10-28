package test;

import (
	"testing"
	"context"
	"log"
	"time"
	"sync"
	"fmt"

	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	"github.com/vctqs1/golang-manabie/pkg/api"

	_service "github.com/vctqs1/golang-manabie/pkg/services"
	"github.com/vctqs1/golang-manabie/database"
)
type ProductResponse struct {
	Res *protov1.BuyProductsResponse
	Err error
}

func BuyAProductOfConcurent(arg1, arg2 []*protov1.BuyProduct) []ProductResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.Connect();
	if err != nil {
		log.Printf("did not connect: <%+v>\n\n", err)
	}

	client := _service.NewProductsService(db);


	var wg sync.WaitGroup

	wg.Add(2)

	res := make([]ProductResponse, 0)
	

	go func() {
		fmt.Printf("call 1\n")
		res1, err1 := client.BuyProducts(ctx, &protov1.BuyProductsRequest{
			Products: arg1,
		});
		res = append(res, ProductResponse{res1, err1})
		wg.Done()

	}()

	go func() {
		fmt.Printf("call 2\n")
		res2, err2 := client.BuyProducts(ctx, &protov1.BuyProductsRequest{
			Products: arg2,
		});
		res = append(res, ProductResponse{res2, err2})
		wg.Done()

	}()
	wg.Wait()
	return res

}


func BuyConcurentProduct(arg1, arg2 []*protov1.BuyProduct) (bool, error) {

	responses := BuyAProductOfConcurent(arg1, arg2)

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

//
func TestBuyConcurentProducts(t *testing.T) {
	res, err := BuyConcurentProduct([]*protov1.BuyProduct{{
			ProductId: 5,
			Quantities: 1,
		},
	}, []*protov1.BuyProduct{
		{
			ProductId: 5,
			Quantities: 1,
		},
	})

	if res != true {
		t.Error(err)
	} else {
		log.Printf("Response: %+v\nError: %+v\n", res, err)
	}
}


func TestBuyInvalidConcurentProducts(t *testing.T) {
	res, err := BuyConcurentProduct([]*protov1.BuyProduct{{
			ProductId: 6,
			Quantities: 1,
		},
	}, []*protov1.BuyProduct{{
			ProductId: 6,
			Quantities: 111,
		},
	})

	if res != true {
		t.Error(err)
	} else {
		log.Printf("Response: %+v\nError: %+v\n", res, err)
	}
}