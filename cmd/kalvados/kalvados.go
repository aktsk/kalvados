// +build ignore

package main

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aktsk/kalvados/receipt"
)

func main() {
	var keyFileName string
	var certFileName string

	flag.StringVar(&keyFileName, "keyFile", "key.pem", "Private Key file")
	flag.StringVar(&certFileName, "certFile", "cert.pem", "Cetificate file")

	flag.Parse()

	keyFile, err := os.Open(keyFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer keyFile.Close()

	keyPEM, err := ioutil.ReadAll(keyFile)
	if err != nil {
		log.Fatal(err)
	}

	keyDER, _ := pem.Decode(keyPEM)
	key, err := x509.ParsePKCS1PrivateKey(keyDER.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	certFile, err := os.Open(certFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer certFile.Close()

	certPEM, err := ioutil.ReadAll(certFile)
	if err != nil {
		log.Fatal(err)
	}

	certDER, _ := pem.Decode(certPEM)
	cert, err := x509.ParseCertificate(certDER.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	encodedReceipt, _ := receipt.Encode(stdin.Bytes(), key, cert)

	fmt.Println(encodedReceipt)
}
