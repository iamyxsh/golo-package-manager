package main

import (
	"context"
	"fmt"
	"github.com/iamyxsh/golo-package-magager/server/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("PG_PASSWORD")
	username := os.Getenv("PG_USERNAME")
	dsn := fmt.Sprintf("host=surus.db.elephantsql.com user=%s password=%s dbname=%s port=5432 sslmode=disable", username, password, username)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	DB = db

}

func main() {

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLinkServiceServer(grpcServer, &Server{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}

type Server struct {
	pb.LinkServiceServer
}

func (s *Server) GetLink(ctx context.Context, req *pb.LinkMessageRequest) (*pb.LinkMessageResponse, error) {
	pkg := Package{}

	err := DB.First(&pkg, "name = ?", req.Package)

	fmt.Println(pkg)

	if err.Error != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Link not found for %s", req.Package)
	}

	return &pb.LinkMessageResponse{Link: pkg.Link}, nil
}

type Package struct {
	Name string `gorm:"primaryKey"`
	Link string
}
