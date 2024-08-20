package main

import (
	"log"
	"net"

	"github.com/Time-Capsule/Memory-Service/config"
	"github.com/Time-Capsule/Memory-Service/genproto"
	"github.com/Time-Capsule/Memory-Service/service"
	"github.com/Time-Capsule/Memory-Service/storage/mongo"
	"github.com/Time-Capsule/Memory-Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	db, err := postgres.DbConnection()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	mongo, err := mongo.SetupMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal("error while listening: %v", err)
	}

	s := grpc.NewServer()
	genproto.RegisterMemoryServiceServer(s, service.NewMemoryService(mongo))
	genproto.RegisterMediaServiceServer(s, service.NewMediaService(db))
	genproto.RegisterCommentServiceServer(s, service.NewCommentService(db))

	log.Printf("Server started on port: %v", cfg.HTTPPort)
	if err := s.Serve(lis); err != nil {
		log.Fatal("error while serving: %v", err)
	}
}
