package pow


type POW interface {
	GetChallenge() Challenge
	VerifyProof(challenge Challenge, solution string) bool
}

type Challenge struct {
	ChallengeString string
	Difficulty int
	Type string
}

func NewPOWhashcash(difficulty int) POW {
	return &POWHashcash{difficulty: difficulty}
}

func NewPOWQuadraticResidue(difficulty int) POW {
	return &POWQuadraticResidue{difficulty: difficulty}
}

