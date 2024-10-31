package services

import (
	"Inventory_Microservice/internal/stock_level_service/model"
	"Inventory_Microservice/internal/stock_level_service/repository"
	"Inventory_Microservice/pkg/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type StockServer struct {
	store repository.StockRepository
	pb.UnimplementedStockLevelsServer
}

func NewStockServer(store repository.StockRepository) *StockServer {
	return &StockServer{store: store}
}

func (s StockServer) AddStockLevel(ctx context.Context, req *pb.AddStockLevelRequest) (*pb.AddStockLevelResponse, error) {
	log.Println("<<<<<<<>>> AddStock Request <<<>>>>>>")

	if req.StockLevel.ProductId == "" || req.StockLevel.ProductId == " " {
		errMsg := "Product ID is required!"
		log.Println(errMsg)
		return nil, status.Errorf(codes.InvalidArgument, errMsg)
	}

	stockLevel := &model.StockLevel{
		ProductID:    req.StockLevel.ProductId,
		Quantity:     req.StockLevel.Quantity,
		MinThreshold: req.StockLevel.MinThreshold,
		MaxThreshold: req.StockLevel.MaxThreshold,
	}

	err := s.store.AddStock(stockLevel)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating stock levels: %v", err)
	}

	response := &pb.AddStockLevelResponse{
		Message: "stock update success!",
		StockLevel: &pb.StockLevel{
			ProductId:    stockLevel.ProductID,
			Quantity:     stockLevel.Quantity,
			MinThreshold: stockLevel.MinThreshold,
			MaxThreshold: stockLevel.MaxThreshold,
		},
	}

	log.Printf("<<<<<<>>>>>> Stock update successfully: %+v\n", response)
	return response, nil
}
