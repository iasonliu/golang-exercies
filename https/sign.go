package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	// Generate a private key
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(priv.PublicKey)
	// Hash a message
	msg := []byte("The bourgeois human is a virus on the hard drive of the working robot!")
	hash := sha512.Sum512(msg)

	//Sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		log.Fatalln(err)
	}

	// pubilc part verifies signature of hash
	pub := &priv.PublicKey
	fmt.Println(ecdsa.Verify(pub, hash[:], r, s))
}
