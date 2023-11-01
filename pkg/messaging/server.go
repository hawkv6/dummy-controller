package messaging

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/hawkv6/dummy-controller/pkg/api"
	"google.golang.org/grpc"
)

type MessagingServer struct {
	api.UnimplementedIntentServiceServer
}

func NewMessagingServer() *MessagingServer {
	return &MessagingServer{}
}

func (s *MessagingServer) Start() {
	listenAddress := config.Params.Address + ":" + strconv.Itoa(config.Params.Port)
	fmt.Println("Listening on " + listenAddress)
	list, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterIntentServiceServer(grpcServer, s)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *MessagingServer) GetIntentDetails(ctx context.Context, intent *api.Intent) (*api.Response, error) {
	response := &api.Response{
		Ipv6Addresses: []string{
			"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			"2001:0db8:85a3:0000:0000:8a2e:0370:7335",
		},
	}

	return response, nil
}
