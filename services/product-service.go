package services


import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type productsServiceServer struct {
	db *sql.DB
}

func NewProductsService(db *sql.DB) pb.ProductsServiceServer {
	return &productsServiceServer{db: db}

}