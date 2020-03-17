package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kaboc/piano_server/protos"
)

const defaultPort = 50051

var port int

func init() {
	flag.IntVar(&port, "p", defaultPort, "Port number for RPCs")
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on port %d\n", port)

	s := grpc.NewServer()
	pb.RegisterPianoServer(s, &pianoService{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
