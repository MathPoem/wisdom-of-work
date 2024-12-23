package pow

import (
	"encoding/hex"
	"testing"
)

// TestPOWHashcashGetChallenge ensures that we get a valid challenge back
func TestPOWHashcashGetChallenge(t *testing.T) {
	difficulty := 2
	hashcash := &POWHashcash{difficulty: difficulty}

	challenge := hashcash.GetChallenge()
	if challenge.Type != "hashcash" {
		t.Errorf("Expected challenge type to be 'hashcash', got '%s'", challenge.Type)
	}
	if challenge.Difficulty != difficulty {
		t.Errorf("Expected difficulty %d, got %d", difficulty, challenge.Difficulty)
	}
	if len(challenge.ChallengeString) == 0 {
		t.Error("Challenge string should not be empty")
	}
}

// TestPOWHashcashVerifyProof tries to generate a valid solution for difficulty=1
func TestPOWHashcashVerifyProof(t *testing.T) {
	difficulty := 1
	hashcash := &POWHashcash{difficulty: difficulty}
	challenge := hashcash.GetChallenge()

	// Attempt to brute force a valid nonce at difficulty = 1
	// This is feasible for small difficulty in testing
	validNonce := ""
	for i := 0; i < 1000000; i++ {
		nonceBytes := []byte{byte(i % 256), byte((i / 256) % 256), byte((i / 65536) % 256), byte((i / 16777216) % 256)}
		nonceHex := hex.EncodeToString(nonceBytes)
		if hashcash.VerifyProof(challenge, nonceHex) {
			validNonce = nonceHex
			break
		}
	}

	if validNonce == "" {
		t.Errorf("Could not find a valid nonce for difficulty=%d within 1,000,000 tries", difficulty)
		return
	}

	// Verify that the validNonce indeed passes
	if !hashcash.VerifyProof(challenge, validNonce) {
		t.Error("Expected VerifyProof to return true for a known valid solution, got false")
	}

	// Verify that a random (invalid) nonce fails
	if hashcash.VerifyProof(challenge, "deadbeef") {
		t.Error("Expected VerifyProof to return false for an invalid solution, got true")
	}
}
