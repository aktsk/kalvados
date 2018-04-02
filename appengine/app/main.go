package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aktsk/kalvados/server"
)

func init() {
	// Get a private key
	var key *rsa.PrivateKey

	keyFile, err := os.Open("key.pem")
	defer keyFile.Close()

	if err == nil {
		keyPEM, err := ioutil.ReadAll(keyFile)
		if err != nil {
			log.Fatal(err)
		}

		keyDER, _ := pem.Decode(keyPEM)
		key, err = x509.ParsePKCS1PrivateKey(keyDER.Bytes)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		keyPEM := os.Getenv("PRIVATE_KEY")
		if keyPEM != "" {
			keyPEM = revertPEM(keyPEM)
			keyDER, _ := pem.Decode([]byte(keyPEM))
			key, err = x509.ParsePKCS1PrivateKey(keyDER.Bytes)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Get a certificate
	var cert *x509.Certificate

	certFile, err := os.Open("cert.pem")
	defer certFile.Close()

	if err == nil {
		certPEM, err := ioutil.ReadAll(certFile)
		if err != nil {
			log.Fatal(err)
		}

		certDER, _ := pem.Decode(certPEM)
		cert, err = x509.ParseCertificate(certDER.Bytes)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		certPEM := os.Getenv("CERTIFICATE")
		if certPEM != "" {
			certPEM = revertPEM(certPEM)
			certDER, _ := pem.Decode([]byte(certPEM))
			cert, err = x509.ParseCertificate(certDER.Bytes)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	http.HandleFunc("/", server.Encode(key, cert))
}

// It seems GAE/Go can not handle environment variables that has
// return code. So in YAML file, I  use ">-" to replace return
// code with white space. This function is for reverting back PEM data
// to original.
func revertPEM(c string) string {
	c = strings.Replace(c, " ", "\n", -1)
	c = strings.Replace(c, "\nCERTIFICATE", " CERTIFICATE", -1)
	c = strings.Replace(c, "\nPRIVATE\nKEY", " PRIVATE KEY", -1)
	return c
}
