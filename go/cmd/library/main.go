package main

import (
	"log"
	"net"

	"github.com/bvk/bazel-repo-template/go/library"
	"google.golang.org/grpc"

	librarypb "github.com/bvk/bazel-repo-template/proto/library"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	library, err := library.New()
	if err != nil {
		log.Fatalf("couldn't create library: %v", err)
	}

	grpcServer := grpc.NewServer()
	librarypb.RegisterLibraryServiceServer(grpcServer, library)
	grpcServer.Serve(listener)
}
