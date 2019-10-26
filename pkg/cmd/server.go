package cmd

import (
	"flag"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	_grpc "github.com/vctqs1/golang-manabie/pkg/protocol/grpc"
	_services "github.com/vctqs1/golang-manabie/pkg/services"
	"github.com/vctqs1/golang-manabie/pkg/api"
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




func RunServer() error {
	
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


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
	param := "parseTime=true&charset=utf8mb4,utf8"

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
	// go RunGRPC(cfg);

	v1API := _services.NewProductsService(db)

	return _grpc.RunServer(ctx, v1API, cfg.GRPCPort)

}
// RunGRPC runs gRPC server and HTTP gateway
func RunGRPC(cfg Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel();



	grpcServerEndpoint := ":"+cfg.GRPCPort;
	fmt.Printf("GRPC Port: %s\n",grpcServerEndpoint)

	grpcGateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := protov1.RegisterProductsServiceHandlerFromEndpoint(ctx, grpcGateway, grpcServerEndpoint, opts)
	if err != nil {
		fmt.Printf("fail %s", err)
	}


	grpcGatewayRouter := mux.NewRouter()
	grpcGatewayRouter.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`hello world!`))
	})).Methods("GET");

	grpcGatewayRouter.NewRoute().Handler(grpcGateway)

	if err := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization"}),
	)(grpcGatewayRouter)); err != nil {
		fmt.Printf("failed to serve %s", err)
	}

}
