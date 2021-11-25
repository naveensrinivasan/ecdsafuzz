package main

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	crand "crypto/rand"
	"io"
	"math/big"
	"os"

	"github.com/pkg/errors"
)

func main() {
	data, _ := os.ReadFile("da39a3ee5e6b4b0d3255bfef95601890afd80709.output")
	x := ecdsa.PrivateKey{}
	z := new(big.Int)
	z.SetBytes(data)
	x.X = z
	x.Y = z
	x.D = z
	SignMessage(bytes.NewReader(data), &x)
}

func selectRandFromOpts() io.Reader {
	rand := crand.Reader
	return rand
}

func hashMessage(rawMessage io.Reader, hashFunc crypto.Hash) ([]byte, error) {
	if rawMessage == nil {
		return nil, errors.New("message cannot be nil")
	}
	if hashFunc == crypto.Hash(0) {
		return io.ReadAll(rawMessage)
	}
	hasher := hashFunc.New()
	// avoids reading entire message into memory
	if _, err := io.Copy(hasher, rawMessage); err != nil {
		return nil, errors.Wrap(err, "hashing message")
	}
	return hasher.Sum(nil), nil
}

func SignMessage(message io.Reader, e *ecdsa.PrivateKey) ([]byte, error) {
	rand := selectRandFromOpts()
	digest, err := hashMessage(message, crypto.SHA512)
	if err != nil {
		return nil, err
	}
	return ecdsa.SignASN1(rand, e, digest)
}
