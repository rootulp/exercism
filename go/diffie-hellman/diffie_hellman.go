package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	max := big.NewInt(0)
	two := big.NewInt(2)
	max.Sub(p, two)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return two.Add(two, n)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	panic("Please implement the PublicKey function")
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	panic("Please implement the NewPair function")
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	panic("Please implement the SecretKey function")
}
