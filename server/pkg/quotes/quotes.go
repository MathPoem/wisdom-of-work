package quotes

import (
    "crypto/rand"
    "math/big"
)

var wisdomList = []string{
    "The only true wisdom is in knowing you know nothing. – Socrates",
    "In the middle of difficulty lies opportunity. – Albert Einstein",
    "There is no substitute for hard work. – Thomas Edison",
    "Well done is better than well said. – Benjamin Franklin",
    "Quality is not an act, it is a habit. – Aristotle",
}

func GetRandomQuote() string {
    bigIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(wisdomList))))
    if err != nil {
        return wisdomList[0]
    }
    return wisdomList[bigIndex.Int64()]
}