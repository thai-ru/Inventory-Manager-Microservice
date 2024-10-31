package main

import (
	"Inventory_Microservice/config"
	"Inventory_Microservice/db"
	"Inventory_Microservice/internal/product_service/repository"
	"Inventory_Microservice/internal/product_service/services"
	repository2 "Inventory_Microservice/internal/stock_level_service/repository"
	services2 "Inventory_Microservice/internal/stock_level_service/services"
	"Inventory_Microservice/pkg/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net" // Import the net package
)

func main() {

	// Load the config based on environment
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database
	database, err := db.NewDB(&cfg.ProductDB)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize stock database
	stockDatabase, err := db.NewStockDB(&cfg.StockDB)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	productStore := repository.NewProductRepository(database)
	stockStore := repository2.NewStockRepository(stockDatabase)

	productService := services.NewProductServer(productStore)
	stockService := services2.NewStockServer(stockStore)

	grpcServer := grpc.NewServer()

	pb.RegisterInventoryManagerServer(grpcServer, productService)
	pb.RegisterStockLevelsServer(grpcServer, stockService)

	// Enable reflection
	reflection.Register(grpcServer)

	// Create a TCP listener
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start gRPC server
	log.Printf("Starting gRPC server on %s...", addr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
