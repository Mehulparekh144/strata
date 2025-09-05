package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strata/api/strata/api"
	"strata/pkg/engine"
	"strata/pkg/strata"

	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// 1. Initialize the storage engine
	db, err := engine.NewDBEngine("./data")
	if err != nil {
		log.Fatalf("Failed to open db: %v 🙁", err)
		return
	}
	defer db.Close()

	// 2. Initialize the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v 😓", err)
		return
	}

	server := strata.NewServer(db)
	grpcServer := grpc.NewServer()

	api.RegisterStrataServer(grpcServer, server)

	// Enable gRPC reflection for grpcurl
	reflection.Register(grpcServer)

	log.Println(getASCIIWelcome())
	log.Println("Server is running on port 50051 🚀")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v 🙁", err)
		return
	}
}

func getASCIIWelcome() string {
	file, err := os.Open("ascii.txt")
	if err != nil {
		log.Fatalf("Failed to open ascii.txt: %v 🙁", err)
		return ""
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}
