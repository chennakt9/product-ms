package services

import (
	"context"
	"net/http"

	"github.com/chennakt9/product-ms/pkg/db"
	"github.com/chennakt9/product-ms/pkg/models"
	pb "github.com/chennakt9/product-ms/pkg/pb"
)

type Server struct{
	H db.Handler
	pb.ProductServiceServer
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.NoParam) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Message: "Hello, Product service is up",
	}, nil
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Name = req.Name
	product.Stock = req.Stock
	product.Price = req.Price

	if res := s.H.DB.Create(&product); res.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error: res.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id: product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	if res := s.H.DB.First(&product, req.Id); res.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error: res.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id: product.Id,
		Name: product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data: data,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if res := s.H.DB.First(&product, req.Id); res.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error: res.Error.Error(),
		}, nil
	}

	if product.Stock <= req.Quantity {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error: "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if res := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); res.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error: "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - req.Quantity

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Save(&log)

	return &pb.DecreaseStockResponse {
		Status: http.StatusOK,
	}, nil
}