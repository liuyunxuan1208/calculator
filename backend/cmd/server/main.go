package main

import (
	"calculator/backend/internal/calculator/repository"
	"calculator/backend/internal/calculator/service"
	"calculator/backend/internal/transport/grpc"
	
	"calculator/backend/pkg/protobuf/calculator/v1/v1connect"
	
	
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {

	// 初始化存储
	repo := repository.NewMemoryRepository()
	// 初始化服务（注入存储）
	calcService := service.NewCalculatorService(repo)

	grpcHandler := grpc.NewHandler(calcService)

	mux := http.NewServeMux()
	path, handler := v1connect.NewCalculatorServiceHandler(grpcHandler)
	mux.Handle(path, handler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}).Handler(h2c.NewHandler(mux, &http2.Server{}))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
