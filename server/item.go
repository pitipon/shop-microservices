package server

import (
	"log"

	"github.com/pitipon/shop-microservices/modules/item/itemHandler"
	"github.com/pitipon/shop-microservices/modules/item/itemRepository"
	"github.com/pitipon/shop-microservices/modules/item/itemUsecase"

	itemPb "github.com/pitipon/shop-microservices/modules/item/itemPb"
	grpccon "github.com/pitipon/shop-microservices/pkg/gprccon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// start gRPC server
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gPRC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item/v1")

	// Health check
	item.GET("/", s.healthCheckService)
}
