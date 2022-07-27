package cli

import (
	"flag"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"petshop-grpc-go/pkg/server"
)

var (
	bindAddr string
	bindPort uint
	logLevel int
)

func Run() {
	flag.StringVar(&bindAddr, "address", "0.0.0.0", "The address to bind the gRPC Server to.")
	flag.UintVar(&bindPort, "port", 50051, "The port to bind the gRPC Server to.")
	flag.IntVar(&logLevel, "log-level", 1, "The log level to set. (-1 .. 5)")
	flag.Parse()

	if 5 >= logLevel && logLevel >= -1 {
		zerolog.SetGlobalLevel(zerolog.Level(logLevel))
		log.Debug().Msgf("setting global log level to %s", zerolog.Level(logLevel).String())
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Debug().Msgf("setting global log level to %s", zerolog.InfoLevel.String())
	}

	s := server.NewServer(bindAddr, bindPort)
	if err := s.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
