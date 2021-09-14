package grpc

import (
	"context"
	chat "github.com/sla10132000/a"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) SayHello(ctx context.Context, input *chat.Message) (*chat.Message, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &chat.Message{
		Body: "Hello again ",
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	chat.RegisterChatServiceServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice")
	}
	return nil
}
