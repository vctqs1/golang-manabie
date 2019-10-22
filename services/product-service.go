package services


import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/vctqs1/golang-manabie/pkg/api"
)

type productsServiceServer struct {
	db *sql.DB
}

func NewProductsService(db *sql.DB) protov1.ProductsServiceServer {
	return &productsServiceServer{
		db: db,
	}
}