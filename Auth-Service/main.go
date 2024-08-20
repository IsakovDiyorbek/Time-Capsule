package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/Time-Capsule/Auth-Service/api"
	"github.com/Time-Capsule/Auth-Service/api/handler"
	pr "github.com/Time-Capsule/Auth-Service/api/kafka"
	"github.com/Time-Capsule/Auth-Service/genproto/auth"
	"github.com/Time-Capsule/Auth-Service/genproto/user"
	consumer "github.com/Time-Capsule/Auth-Service/kafka"
	"github.com/Time-Capsule/Auth-Service/service"
	"github.com/Time-Capsule/Auth-Service/storage/postgres"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dbPostgres, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connect db:", err.Error())
	}

	liss, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, service.NewAuthService(dbPostgres))
	user.RegisterUserServiceServer(s, service.NewUserService(dbPostgres))

	AuthService := service.NewAuthService(dbPostgres)
	UserService := service.NewUserService(dbPostgres)

	brokers := []string{"localhost:9092"}

	kcm := consumer.NewKafkaConsumerManager()
	if err := kcm.RegisterConsumer(brokers, "auth", "register", consumer.AuhtRegister(AuthService)); err != nil {
		if err == consumer.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'auth' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "user", "update", consumer.Change(UserService)); err != nil {
		if err == consumer.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'user' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	log.Printf("server listening at %v", liss.Addr())

	go func() {
		if err := s.Serve(liss); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	userConn, err := grpc.NewClient(fmt.Sprintf("auth_service%s", ":9999"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting: ", err.Error())
	}

	defer userConn.Close()

	auth := auth.NewAuthServiceClient(userConn)
	userService := user.NewUserServiceClient(userConn)
	kafkaProducer, err := pr.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		slog.Info("error:", err)
	}
	redisDb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	h := handler.Handler{Auth: auth, User: userService, Redis: redisDb, Kafka: kafkaProducer}

	r := api.NewGin(h)

	fmt.Println("Server started on port: 8085")
	r.Run(":8085")

}
