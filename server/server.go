package server

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aktsk/kalvados/receipt"
)

// Response is for respond base64 encoded receipt data
type Response struct {
	ReceiptData string `json:"receipt-data"`
}

// Serve is for serving rceipt generator
func Serve(port int, key *rsa.PrivateKey, cert *x509.Certificate) {
	http.HandleFunc("/", Encode(key, cert))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// Encode encodes JSON receipt data
func Encode(key *rsa.PrivateKey, cert *x509.Certificate) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := receipt.Encode(body, key, cert)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := Response{ReceiptData: res}

		responseBody, err := json.Marshal(response)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}
