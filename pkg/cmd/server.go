package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/gorilla/handlers"

	_grpc "github.com/vctqs1/golang-manabie/pkg/protocol/grpc"
	"github.com/vctqs1/golang-manabie/pkg/services"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// for {
	// 	ctx, cancel := context.WithCancel(context.Background())
	// 	go cancel()

	// 	row := db.QueryRowContext(ctx, `SELECT 1`)
	// 	var a int
	// 	if err := row.Scan(&a); err != nil && err != context.Canceled {
	// 		log.Fatal("Connection SQL: ", err)
	// 	}
	// }
	defer db.Close()

	v1API := protov1.NewProductsService(db)

	grpcGateway := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true, EnumsAsInts: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = protov1.RegisterProductsServiceHandlerFromEndpoint(ctx, grpcGateway, cfg.GRPCPort, opts)
	if err != nil {
		fmt.Printf("fail ", err)
	}


	grpcGatewayRouter := mux.NewRouter()
	grpcGatewayRouter.NewRoute().Handler(grpcGateway)
	if err := http.ListenAndServe(cfg.GRPCPort, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization"}),
	)(grpcGatewayRouter)); err != nil {
		fmt.Printf("failed to serve", err)
	}

	return _grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
