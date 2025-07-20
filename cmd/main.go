package main

import (
	"context"
	"fmt"
	cluster_manager_service "github.com/Na322Pr/cluster-manager-service/internal/app/cluster-manager-service"
	"github.com/Na322Pr/cluster-manager-service/internal/config"
	"github.com/Na322Pr/cluster-manager-service/internal/service"
	desc "github.com/Na322Pr/cluster-manager-service/pkg/api"
	"go.uber.org/zap"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os/signal"
	"syscall"

	nomad "github.com/hashicorp/nomad/api"
)

func main() {
	cfg := config.MustLoad()
	grpcAddress := cfg.GetGRPCAddress()

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	//logger := zap.NewProduction()
	//defer logger.Sync()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("Service config info",
		zap.String("grpcAddress", grpcAddress),
	)

	nomadClient, _ := nomad.NewClient(&nomad.Config{Address: "http://localhost:4646"})

	cmService := service.NewClusterManagerService(grpcAddress, nomadClient, logger)
	cmApp := cluster_manager_service.NewImplementation(cmService)

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	desc.RegisterClusterManagerServer(grpcServer, cmApp)

	logger.Info("Starting cluster...")
	go func() {
		if err := cmService.RunCluster(); err != nil {
			log.Fatalf("failed to run cluster: %v", err)
		}
	}()

	logger.Info(fmt.Sprintf("Starting grpc server on %s...", grpcAddress))
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to listen and server grpc server: %v", err)
		}
	}()

	<-stop
	fmt.Println("\nShutting down servers...")
	os.Exit(0)
}
