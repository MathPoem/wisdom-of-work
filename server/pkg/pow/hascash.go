package pow

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"wisdom-of-work-server/pkg/logger"
)


type POWHashcash struct {
	difficulty int
}

func (p *POWHashcash) GetChallenge() Challenge {
	challenge := make([]byte, 16)
	_, err := rand.Read(challenge)
	if err != nil {
		return Challenge{}
	}
    challengeHex := hex.EncodeToString(challenge)

	return Challenge{ChallengeString: challengeHex, Difficulty: p.difficulty, Type: "hashcash"}
}

func (p *POWHashcash) VerifyProof(challenge Challenge, solution string) bool {
	
    nonce, err := hex.DecodeString(solution)
	if err != nil {
		return false
	}

    challengeBytes, err := hex.DecodeString(challenge.ChallengeString)
	if err != nil {
		return false
	}

    data := append(challengeBytes, nonce...)
    hash := sha256.Sum256(data)
    hashHex := hex.EncodeToString(hash[:])

    logger.Log.Infof("Hash: %s", hashHex)
    prefix := ""
    for i := 0; i < p.difficulty; i++ {
		prefix += "0"
	}

	return hashHex[:p.difficulty] == prefix
}
