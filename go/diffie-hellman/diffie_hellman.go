// Package diffiehellman provides Diffie Hellman solution.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey returns a number between 2 and p to be used as private key.
func PrivateKey(p *big.Int) *big.Int {
	min := big.NewInt(2)
	max := new(big.Int).Sub(p, min)
	n, _ := rand.Int(rand.Reader, max)
	return new(big.Int).Add(n, min)
}

// PublicKey calculates a public key from a private key and prime numbers p and g.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair creates a new pair of private and public keys from prime numbers p and g.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	return priv, PublicKey(priv, p, g)
}

// SecretKey calculates a secret key from two public and private key pairs.
func SecretKey(private, publicPeer, p *big.Int) *big.Int {
	return new(big.Int).Exp(publicPeer, private, p)
}
