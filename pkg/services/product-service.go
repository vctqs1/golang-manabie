package protov1


import (
	"strings"
	"context"
	"database/sql"
	"fmt"
	"time"

	// "github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
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

// connect returns SQL database connection from the pool
func (s *productsServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}


func arrayToString(a []int64, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func (rcv *productsServiceServer) GetProducts(ctx context.Context, req *protov1.GetProductsRequest) (*protov1.GetProductsResponse, error) {

	// get SQL connection from pool
	c, err := rcv.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	//get products by list ids
	query := ""
	if len(req.ProductIds) > 0 {
		query =  "WHERE id IN (" + arrayToString(req.ProductIds, ", ") + ")";
	}
		
	rows, err := c.QueryContext(ctx, "SELECT id, title, quantities FROM products " + query)
	
	if err != nil {
		return nil, errors.Wrap(err, "db.QueryEx")
	}

	defer rows.Close()

	
	products := make([]*protov1.Product, 0);
    for rows.Next() {
		e := &protov1.Product{};
		err = rows.Scan(&e.Id, &e.Title, &e.Quantities);
		if err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}
	
		products = append(products, e)

	}

	if rows.Err() != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	return &protov1.GetProductsResponse{
		Products: products,
	}, nil;
}
func (rcv *productsServiceServer) BuyProducts(ctx context.Context, req *protov1.BuyProductsRequest) (*protov1.BuyProductsResponse, error) {
	// get SQL connection from pool
	c, err := rcv.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	
	
	products := make([]*protov1.Product, 0);

    for _, value := range req.Products {

		//get products by list ids
		query := fmt.Sprintf("SELECT id, title, quantities FROM products WHERE id = %d AND quantities >= %d;" , value.ProductId, value.Quantities);
		rows, err := c.QueryContext(ctx, query)
		
		if err != nil {
			return nil, errors.Wrap(err, "db.QueryEx")
		}

		
		e := &protov1.Product{};
		
		defer rows.Close()

		if !rows.Next() {
			if err := rows.Err(); err != nil {
				return nil, status.Error(codes.Unknown, "failed to retrieve data: "+err.Error())
			}
			return nil, status.Error(codes.NotFound, fmt.Sprintf("product with id=%d is not found", value.ProductId))
		}

		err = rows.Scan(&e.Id, &e.Title, &e.Quantities);
		if err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}
		
		e.Quantities = e.Quantities - value.Quantities;
		fmt.Printf("\nQuantities: %d\n", e.Quantities)

		products = append(products, e)

	}

	if len(products) != len(req.Products) {
		return nil, status.Error(codes.InvalidArgument, "buy error");
	} 

	now := time.Now().UTC().Format("2006-01-02 03:04:05");
	fmt.Printf(now)
    // for _, value := range products {
	// 	query := fmt.Sprintf("UPDATE products SET `quantities`=%d, `updated_at`='%s' WHERE `id`=%d;", value.Quantities, now, value.Id);
		
	// 	res, err := c.ExecContext(ctx, query)

	// 	if err != nil {
	// 		return nil, status.Error(codes.Unknown, "failed to update after buy products -> "+err.Error())
	// 	}

	// 	fmt.Printf("\nQuery: %s\n", query)
	// 	rows2, err := res.RowsAffected()
	// 	if err != nil {
	// 		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	// 	}

	// 	if rows2 == 0 {
	// 		return nil, status.Error(codes.NotFound, fmt.Sprintf("product with id=%d is not found", value.Id))
	// 	}

	// }


	fmt.Printf("buy products result: <%+v>\n\n", products)

	return &protov1.BuyProductsResponse{}, nil;
}