package tcp_server

import (
	"fmt"
	"net"
	"strings"

	"wisdom-of-work-server/pkg/pow"
	"wisdom-of-work-server/pkg/quotes"

	"github.com/sirupsen/logrus"
)

func handleConnection(conn net.Conn, log *logrus.Logger, pow pow.POW) {
	
    // Send challenge
    challenge := pow.GetChallenge()
	msg := fmt.Sprintf("%s %d %s\n", challenge.ChallengeString, challenge.Difficulty, challenge.Type)
	_, err := conn.Write([]byte(msg))
	if err != nil {
		log.WithError(err).Warn("Error sending challenge to client")
		return
	}

	log.Infof("Sent challenge %s of type %s with difficulty %d to client %s", challenge.ChallengeString, challenge.Type, challenge.Difficulty, conn.RemoteAddr())

	// Read client's solution
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.WithError(err).Warn("Error reading from client")
		return
	}

	solution := strings.TrimSpace(string(buf[:n]))

	// Verify PoW
	if !pow.VerifyProof(challenge, solution) {
		log.Warnf("Invalid proof of work from client %s", conn.RemoteAddr())
		return
	}

	// Send random quote
	quote := quotes.GetRandomQuote()
	_, err = conn.Write([]byte(quote + "\n"))
	if err != nil {
		log.WithError(err).Warn("Error sending quote to client")
		return
	}

	log.Infof("Sent quote to client %s: %q", conn.RemoteAddr(), quote)
}
