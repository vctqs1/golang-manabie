package database


import (
	"flag"
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "github.com/go-sql-driver/mysql"
)

// Config is configuration for Server
type ConfigStruct struct {
	GRPCPort string

}
var config ConfigStruct;

func init() {

	flag.StringVar(&config.GRPCPort, "grpc-port", "9090", "gRPC port to bind")

	flag.Parse()
}
func GetConfig() (cfg ConfigStruct) {
	cfg = config;
	return cfg;
}
func Connect() (*sql.DB, error) {
	
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	db, err := sql.Open("mysql", "root:golang@tcp(localhost:3333)/golang_manabie?parseTime=true&charset=utf8mb4,utf8")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return db, nil;
	
}
func Conn() (*sql.Conn, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	db, err := Connect();
	if err != nil {
		return nil, err;
	}
	c, err := db.Conn(ctx)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
	
}