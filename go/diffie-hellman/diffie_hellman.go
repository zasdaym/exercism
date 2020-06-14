// Package diffiehellman is solution for problem Diffie Hellman.
package diffiehellman

import (
	"math/big"
	"crypto/rand"
)

// PrivateKey returns new private key.
func PrivateKey(p *big.Int) *big.Int {
	l := big.NewInt(1).Sub(p, big.NewInt(2))
	n, err := rand.Int(rand.Reader, l)
	if err != nil {
		return big.NewInt(2)
	}

	n.Add(n, big.NewInt(2))
	return n
}

// PublicKey generates public key from private key.
func PublicKey(a, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), a, p)
}

// SecretKey generates secret key.
func SecretKey(a, b, p *big.Int) *big.Int {
	return big.NewInt(1).Exp(b, a, p)
}

// NewPair generates new pair of private and public key.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	return priv, PublicKey(priv, p, g)
}
