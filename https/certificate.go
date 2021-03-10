package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"
)

func main() {
	// Generate a private key
	caPriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	// Generate a self-signed certificate
	caTmpl := &x509.Certificate{
		Subject:               pkix.Name{CommonName: "My-ca"},
		SerialNumber:          newSerialNum(), // Choose a random, big number
		BasicConstraintsValid: true,
		IsCA:                  true,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage: x509.KeyUsageDataEncipherment |
			x509.KeyUsageCertSign |
			x509.KeyUsageCertSign,
	}
	caCertDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, caPriv.Public(), caPriv)
	if err != nil {
		log.Fatal(err)
	}
	caPrivDER, err := x509.MarshalECPrivateKey(caPriv)
	if err != nil {
		log.Fatal(err)
	}
	// PEM encode the certificate and private key
	caCertPEM := pem.EncodeToMemory(&pem.Block{Bytes: caCertDER, Type: "CERTIFICATE"})
	caPrivPEM := pem.EncodeToMemory(&pem.Block{Bytes: caPrivDER, Type: "EC PRIVATE KEY"})

	fmt.Println(string(caCertPEM))
	fmt.Println(string(caPrivPEM))

	caCert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		log.Fatal(err)
	}
	// Generate a key pair and certificate template
	servPriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	servTmpl := &x509.Certificate{
		Subject:      pkix.Name{CommonName: "my-server"},
		SerialNumber: newSerialNum(),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour),
		DNSNames:     []string{"localhost"},
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	// sign the serving cert with the CA private key
	servCertDER, err := x509.CreateCertificate(rand.Reader, servTmpl, caCert, servPriv.Public(), caPriv)
	if err != nil {
		log.Fatal(err)
	}
	servPrivDER, err := x509.MarshalECPrivateKey(servPriv)
	if err != nil {
		log.Fatal(err)
	}
	servCertPEM := pem.EncodeToMemory(&pem.Block{Bytes: servCertDER, Type: "CERTIFICATE"})
	servPrivPEM := pem.EncodeToMemory(&pem.Block{Bytes: servPrivDER, Type: "CERTIFICATE"})

	fmt.Println(string(servCertPEM))
	fmt.Println(string(servPrivPEM))
	// Load the certificate and private key as a TLS certificate
	servTLSCert, err := tls.X509KeyPair(servCertPEM, servPrivPEM)
	if err != nil {
		// ...
	}
	serv := http.Server{
		Addr: "localhost:8443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "You're using HTTPS")
		}),
		// Configure TLS options
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{servTLSCert}},
	}
	// Begin serving TLS
	err = serv.ListenAndServeTLS("", "")

	// ...

	// Configure a client to trust the server
	// certPool := x509.NewCertPool()
	// certPool.AppendCertsFromPEM(caCertPEM)
	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{RootCAs: certPool},
	// 	},
	// }
	// resp, err := client.Get("https://localhost:8443/")

}

func newSerialNum() *big.Int {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
