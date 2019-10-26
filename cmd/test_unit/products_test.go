package main;

import (
	// "fmt"
	// "flag"
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
)


func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/product/get", bytes.NewBuffer([]byte(`{"product_ids": [1, 2]}`)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
}