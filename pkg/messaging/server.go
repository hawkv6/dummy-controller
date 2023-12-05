package messaging

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/hawkv6/dummy-controller/pkg/api"
	"google.golang.org/grpc"
)

type MessagingServer struct {
	api.UnimplementedIntentControllerServer
	streamManager     *StreamManager
	MessagingChannels *MessagingChannels
}

type StreamManager struct {
	streams map[api.IntentController_GetIntentPathServer]context.CancelFunc
	mu      sync.Mutex
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: make(map[api.IntentController_GetIntentPathServer]context.CancelFunc),
	}
}

func NewMessagingServer(messagingChannels *MessagingChannels) *MessagingServer {
	return &MessagingServer{
		streamManager:     NewStreamManager(),
		MessagingChannels: messagingChannels,
	}
}

func (m *StreamManager) Add(stream api.IntentController_GetIntentPathServer) context.Context {
	m.mu.Lock()
	defer m.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	m.streams[stream] = cancel

	return ctx
}

func (m *StreamManager) Remove(stream api.IntentController_GetIntentPathServer) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if cancel, ok := m.streams[stream]; ok {
		cancel()
		delete(m.streams, stream)
	}
}

func (s *MessagingServer) Start() {
	listenAddress := config.Params.Address + ":" + strconv.Itoa(config.Params.Port)
	fmt.Println("Listening on " + listenAddress)
	list, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterIntentControllerServer(grpcServer, s)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *MessagingServer) GetIntentPath(stream api.IntentController_GetIntentPathServer) error {
	ctx := s.streamManager.Add(stream)
	defer s.streamManager.Remove(stream)

	for {
		in, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}
		log.Printf("Received request for domain: " + in.Ipv6DestinationAddress)
		s.MessagingChannels.ChMessageIntentRequest <- in
		go s.GetIntentPathResponse(stream, ctx)
	}
}

func (s *MessagingServer) GetIntentPathResponse(stream api.IntentController_GetIntentPathServer, ctx context.Context) {
	for {
		select {
		case response := <-s.MessagingChannels.ChMessageIntentResponse:
			if err := stream.Send(response); err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
