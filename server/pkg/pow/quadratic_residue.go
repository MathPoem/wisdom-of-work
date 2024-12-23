package pow

import (
	"crypto/rand"
	"encoding/hex"
)

type POWQuadraticResidue struct {
	difficulty int
}

func (p *POWQuadraticResidue) GetChallenge() Challenge {
	challenge := make([]byte, 16)
	_, err := rand.Read(challenge)
	if err != nil {
		return Challenge{}
	}
	challengeHex := hex.EncodeToString(challenge)
	return Challenge{ChallengeString: challengeHex, Difficulty: p.difficulty, Type: "quadratic_residue"}
}

func (p *POWQuadraticResidue) VerifyProof(challenge Challenge, solution string) bool {
	return false
}
