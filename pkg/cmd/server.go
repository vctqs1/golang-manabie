package cmd

import (
	"context"
	"fmt"
	// "flag"


	_grpc "github.com/vctqs1/golang-manabie/pkg/protocol/grpc"
	_services "github.com/vctqs1/golang-manabie/pkg/services"
	"github.com/vctqs1/golang-manabie/database"

)



func RunServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, cfg, err := database.Connect();
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	
	v1API := _services.NewProductsService(db)

	
	if len(cfg.GRPCPort) == 0 {
		fmt.Printf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}
	return _grpc.RunServer(ctx, v1API, cfg.GRPCPort)

}
