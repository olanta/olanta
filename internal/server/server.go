package server

import (
	"context"
	"log"

	"github.com/olanta/olanta/proto"
)

type Server struct {
	proto.UnimplementedScannerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Scan(ctx context.Context, req *proto.ScanRequest) (*proto.ScanResponse, error) {
	log.Printf("Received scan request for language: %s, path: %s", req.Language, req.Path)
	for _, issue := range req.Issues {
		log.Printf("Issue: Description=%s, Severity=%s, File=%s, Line=%d, Column=%d",
			issue.Description, issue.Severity, issue.File, issue.Line, issue.Column)
	}
	response := &proto.ScanResponse{
		Success: true,
		Message: "Scan completed successfully",
	}
	return response, nil
}
