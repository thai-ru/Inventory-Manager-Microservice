package services

import (
	"Inventory_Microservice/internal/product_service/model"
	"Inventory_Microservice/internal/product_service/repository"
	"Inventory_Microservice/pkg/pb"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type ProductServer struct {
	store repository.ProductRepository
	pb.UnimplementedInventoryManagerServer
}

func NewProductServer(store repository.ProductRepository) *ProductServer {
	return &ProductServer{store: store}
}

func (s ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Print("<<<<<<<>>>>>> CreateProduct request <<<<<<<<>>>>>>>")

	if req.Product.ProductName == "" && req.Product.BasePrice <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Please make sure to include product name or a valid price ")
	}

	productID := uuid.New().String()

	product := &model.Product{
		ID:          productID,
		ProductName: req.Product.ProductName,
		Category:    req.Product.Category,
		Description: req.Product.Description,
		BasePrice:   float64(req.Product.BasePrice),
	}

	err := s.store.Save(product)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to save product!")
	}

	response := &pb.CreateProductResponse{
		Message: "Product Added successfully!",
		Product: &pb.Product{
			Id:          productID,
			ProductName: product.ProductName,
			Category:    product.Category,
			Description: product.Description,
			BasePrice:   float32(product.BasePrice),
		},
	}

	log.Printf("<<<<<<>>> Product created successfully: %+v", response)
	return response, nil
}
