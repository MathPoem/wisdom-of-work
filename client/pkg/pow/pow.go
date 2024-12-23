package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
	"wisdom-of-work-client/pkg/logger"
)

func SolvePoW(challenge []byte, difficulty int, powType string) []byte {
	rand.NewSource(time.Now().UnixNano())

	prefix := strings.Repeat("0", difficulty)

	tryCount := 0
	for {
		tryCount++

		nonce := make([]byte, 8)
		for i := range nonce {
			nonce[i] = byte(rand.Intn(256))
		}

		data := append(challenge, nonce...)
		sum := sha256.Sum256(data)
		sumHex := hex.EncodeToString(sum[:])

		if strings.HasPrefix(sumHex, prefix) {
			logger.Log.Infof("PoW solved in %d tries. Hash=%s\n", tryCount, sumHex)
			return nonce
		}
	}
}
