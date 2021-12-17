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

func PublicKey(a *big.Int, p *big.Int, g int64) *big.Int {
	temp := big.NewInt(g)
	return temp.Exp(temp, a, p)
}

func SecretKey(a *big.Int, B *big.Int, p *big.Int) *big.Int {
	return B.Exp(B, a, p)

}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	panic("Please implement the NewPair function")
}
