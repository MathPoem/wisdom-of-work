package client

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"wisdom-of-work-client/internal/config"
	"wisdom-of-work-client/pkg/pow"

	"github.com/sirupsen/logrus"
)

type Client struct {
	log  *logrus.Logger
	cfg  *config.Config
	quit chan struct{}
}

func NewClient(cfg *config.Config, log *logrus.Logger) *Client {
	return &Client{
		cfg:  cfg,
		log:  log,
		quit: make(chan struct{}),
	}
}

func (c *Client) Start() {
	ticker := time.NewTicker(time.Duration(c.cfg.IntervalSec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.quit:
			c.log.Info("Client loop stopping.")
			return
		case <-ticker.C:
			conn := connectToServer(c.cfg)
			defer conn.Close()

			challengeHex, difficulty, powType, err := readServerChallenge(conn)
			if err != nil {
				c.log.WithError(err).Error("Error reading server challenge")
				continue
			}

			c.log.Infof("Got challenge: %s (difficulty: %d) of type %s", challengeHex, difficulty, powType)

			challenge, err := decodeChallenge(challengeHex)
			if err != nil {
				c.log.WithError(err).Error("Error decoding challenge")
				continue
			}

			nonce := pow.SolvePoW(challenge, difficulty, powType)
			c.log.Infof("Found nonce: %x", nonce)

			err = sendNonce(conn, nonce)
			if err != nil {
				c.log.WithError(err).Error("error sending nonce")
				continue
			}

			quote, err := readQuote(conn)
			if err != nil {
				c.log.WithError(err).Error("error reading quote")
				continue
			}

			c.log.Infof("Got quote: %s", quote)
		}
	}
}

func (c *Client) Stop() {
	close(c.quit)
}

func connectToServer(cfg *config.Config) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", cfg.ServerAddress, cfg.ServerPort))
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}
	return conn
}

func readServerChallenge(conn net.Conn) (string, int, string, error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return "", 0, "", fmt.Errorf("error reading from server: %v", err)
	}

	lines := strings.Split(string(buf), "\n")
	log.Printf("Lines: %v", lines)

	if len(lines) < 2 {
		return "", 0, "", fmt.Errorf("expected at least 2 lines from server, got: %v", lines)
	}

	firstLine := lines[0]
	parts := strings.Split(firstLine, " ")
	if len(parts) != 3 {
		return "", 0, "", fmt.Errorf("expected format <challengeHex> <difficulty> <type>, got: %s", firstLine)
	}

	challengeHex := parts[0]
	difficultyStr := parts[1]
	powType := parts[2]
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		return "", 0, "", fmt.Errorf("invalid difficulty: %v", err)
	}
	return challengeHex, difficulty, powType, nil
}

func decodeChallenge(challengeHex string) ([]byte, error) {
	challenge, err := hex.DecodeString(challengeHex)
	if err != nil {
		return nil, fmt.Errorf("invalid challenge hex: %v", err)
	}
	return challenge, nil
}

func sendNonce(conn net.Conn, nonce []byte) error {
	_, err := conn.Write([]byte(hex.EncodeToString(nonce)))
	if err != nil {
		return fmt.Errorf("error sending nonce to server: %v", err)
	}
	return nil
}

func readQuote(conn net.Conn) (string, error) {
	quoteBuf := make([]byte, 1024)
	n, err := conn.Read(quoteBuf)
	if err != nil {
		return "", fmt.Errorf("error reading quote from server: %v", err)
	}
	if n > 0 {
		return string(quoteBuf[:n]), nil
	}
	return "", nil
}
