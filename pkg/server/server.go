package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	v1 "petshop-grpc-go/internal/petstore/v1"
	"petshop-grpc-go/pkg/server/user"
)

type (
	Server struct {
		addr string
		port uint
	}
)

func NewServer(addr string, port uint) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}

func (s Server) listenAddr() string {
	return fmt.Sprintf("%s:%d", s.addr, s.port)
}

func (s Server) Run() error {
	runChan := make(chan os.Signal, 1)
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	lis, err := net.Listen("tcp", s.listenAddr())
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	v1.RegisterUserServiceServer(grpcServer, user.NewUserService())
	reflection.Register(grpcServer)

	log.Info().Msgf("running gRPC server on %s", s.listenAddr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error().Err(err).Send()
		}
	}()

	sig := <-runChan
	log.Debug().Msgf("got %s message from os", sig.String())

	log.Info().Msgf("stopping gRPC server due to %s", sig.String())
	grpcServer.GracefulStop()
	log.Debug().Msg("successfully stopped gRPC server")

	return nil
}
