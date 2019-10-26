package cmd

import (
	"flag"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	_grpc "github.com/vctqs1/golang-manabie/pkg/protocol/grpc"
	_services "github.com/vctqs1/golang-manabie/pkg/services"
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
	flag.StringVar(&cfg.GRPCPort, "grpc-port", ":9090", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "localhost", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "root", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "golang_manabie", "Database schema")
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
}
