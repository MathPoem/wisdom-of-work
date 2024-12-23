package tcp_server

import (
	"net"

	"wisdom-of-work-server/internal/config"
	"wisdom-of-work-server/pkg/pow"

	"github.com/sirupsen/logrus"
)

type TcpServer struct {
	log      *logrus.Logger
	cfg      *config.Config
	listener net.Listener
	pow      pow.POW
}

func NewServer(cfg *config.Config, log *logrus.Logger) *TcpServer {
	tcpServer := &TcpServer{cfg: cfg, log: log}
	
	if cfg.POWType == "hashcash" {
		tcpServer.pow = pow.NewPOWhashcash(cfg.Difficulty)
	} else if cfg.POWType == "quadratic_residue" {
		tcpServer.pow = pow.NewPOWQuadraticResidue(cfg.Difficulty)
	}

	return tcpServer
}

func (s *TcpServer) Start() {

	var err error
	s.listener, err = net.Listen("tcp", s.cfg.Port)
	if err != nil {
		s.log.WithError(err).Fatalf("Error listening on %s", s.cfg.Port)
	}
	defer s.listener.Close()

	s.log.Infof("Server listening on %s", s.cfg.Port)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.log.WithError(err).Warn("Error accepting connection")
			continue
		}
		go handleConnection(conn, s.log, s.pow)
	}
}

func (s *TcpServer) Stop() {
	s.log.Info("Shutting down TCP server...")
	s.listener.Close()
}
