package main

import (
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"context"
	"github.com/sessfeld/rules-go-ptypes-repro/protos"
	"google.golang.org/grpc"
)

type Server struct {
	protos.UnimplementedReturnsLroServer
}

func (s *Server) ExportUsers(ctx context.Context, req *protos.ExportRequest) (*longrunningpb.Operation, error) {
	return &longrunningpb.Operation{}, nil
}

func main() {
	srv := grpc.NewServer()
	protos.RegisterReturnsLroServer(srv, &Server{})
}
