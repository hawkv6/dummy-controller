package messaging

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/hawkv6/dummy-controller/pkg/api"
	"github.com/hawkv6/dummy-controller/pkg/intent"
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

func (s *MessagingServer) GetIntentDetails(ctx context.Context, i *api.Intent) (*api.Response, error) {
	sidList, err := intent.GetIntentDetails(i.DomainName, intent.IntentTypeToString(*i.Intent.Enum()))
	if err != nil {
		response := &api.Response{
			Ipv6Addresses: nil,
		}
		return response, err
	}
	response := &api.Response{
		DomainName:    i.DomainName,
		Intent:        i.Intent,
		Ipv6Addresses: sidList,
	}

	return response, nil
}
