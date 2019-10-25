package test_unit;

import (
	"test"
	"net/http/httptest"
	
	"github.com/vctqs1/golang-manabie/pkg/api"
)



func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/product/get", protov1.GetProductsRequest{
		ProductIds: []int64{1,2},
	})
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(protov1.GetProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if len(rr.Body) == 0 {
		t.Errorf("handler returned unexpected body: got %+v want %v",
			rr.Body, expected)
	}
}