package receipt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/aktsk/nolmandy/receipt"
)

func TestReceipt(t *testing.T) {
	privKey, cert := generateKeyAndCert()

	rcpt, err := Encode([]byte(receiptJSON), privKey, cert)
	if err != nil {
		t.Fatal(err)
	}

	parsed, err := receipt.Parse(cert, rcpt)
	if err != nil {
		t.Fatal(err)
	}

	if parsed.ReceiptType != "ProductionSandbox" {
		t.Fatalf("Wrong receipt_type: %s", parsed.ReceiptType)
	}

	if parsed.BundleID != "jp.aktsk.kalvados.test" {
		t.Fatalf("Wrong bundle_id: %s", parsed.BundleID)
	}

	creationDate := time.Unix(1518284220, 0)
	date := time.Time(parsed.CreationDate.Date)
	if date.UTC() != creationDate.UTC() {
		t.Fatalf("Wrong creation_date: %v", date)
	}

	inApp := parsed.InApp[1]

	if inApp.Quantity != 1 {
		t.Fatalf("Wrong qutantity: %d", inApp.Quantity)
	}

	if inApp.ProductID != "jp.aktsk.kalvados.test.iap1" {
		t.Fatalf("Wrong product_id: %s", inApp.ProductID)
	}

	if inApp.TransactionID != "220000359893979" {
		t.Fatalf("Wrong transaction_id: %s", inApp.TransactionID)
	}

	purchaseDate := time.Unix(1503544635, 0)
	date = time.Time(inApp.PurchaseDate.Date)
	if date.UTC() != purchaseDate.UTC() {
		t.Fatalf("Wrong purchase_date: %v", date)
	}

	if inApp.OriginalTransactionID != "220000348788557" {
		t.Fatalf("Wrong transaction_id: %s", inApp.OriginalTransactionID)
	}

	originalPurchaseDate := time.Unix(1500261436, 0)
	date = time.Time(inApp.OriginalPurchaseDate.Date)
	if date.UTC() != originalPurchaseDate.UTC() {
		t.Fatalf("Wrong original_purchase_date: %v", date)
	}

	if inApp.WebOrderLineItemID != 220000072586770 {
		t.Fatalf("Wrong web_order_line_item_id: %d", inApp.WebOrderLineItemID)
	}

	if parsed.OriginalApplicationVersion != "49" {
		t.Fatalf("Wrong original_application_version: %s", parsed.OriginalApplicationVersion)
	}

	originalPurchaseDate = time.Unix(1499441767, 0)
	date = time.Time(parsed.OriginalPurchaseDate.Date)
	if date.UTC() != originalPurchaseDate.UTC() {
		t.Fatalf("Wrong original_purchase_date: %v", date)
	}

}

func generateKeyAndCert() (*rsa.PrivateKey, *x509.Certificate) {
	privKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 32)
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)

	certTemplate := x509.Certificate{
		SerialNumber:       serialNumber,
		SignatureAlgorithm: x509.SHA256WithRSA,
		Subject: pkix.Name{
			CommonName:   "Test Issuer",
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(365, 0, 0),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		IsCA:      true,
	}

	issuerCert := &certTemplate
	issuerKey := privKey

	cert, _ := x509.CreateCertificate(rand.Reader, &certTemplate, issuerCert, privKey.Public(), issuerKey)

	leaf, _ := x509.ParseCertificate(cert)

	return privKey, leaf
}

var receiptJSON = `
{
  "receipt_type": "ProductionSandbox",
  "adam_id": 0,
  "app_item_id": 0,
  "bundle_id": "jp.aktsk.kalvados.test",
  "application_version": "51",
  "download_id": 0,
  "version_external_identifier": 0,
  "original_application_version": "49",
  "in_app": [
    {
      "quantity": "0",
      "product_id": "jp.aktsk.kalvados.test.iap0",
      "transaction_id": "220000350729970",
      "original_transaction_id": "220000348788557",
      "web_order_line_item_id": 220000071891787,
      "is_trial_period": "false",
      "purchase_date": "2017-07-24 03:17:15 Etc/GMT",
      "purchase_date_ms": "1500866235000",
      "purchase_date_pst": "2017-07-23 20:17:15 America/Los_Angeles",
      "original_purchase_date": "2017-07-17 03:17:16 Etc/GMT",
      "original_purchase_date_ms": "1500261436000",
      "original_purchase_date_pst": "2017-07-16 20:17:16 America/Los_Angeles"
    },
    {
      "quantity": "1",
      "product_id": "jp.aktsk.kalvados.test.iap1",
      "transaction_id": "220000359893979",
      "original_transaction_id": "220000348788557",
      "web_order_line_item_id": 220000072586770,
      "is_trial_period": "false",
      "purchase_date": "2017-08-24 03:17:15 Etc/GMT",
      "purchase_date_ms": "1503544635000",
      "purchase_date_pst": "2017-08-23 20:17:15 America/Los_Angeles",
      "original_purchase_date": "2017-07-17 03:17:16 Etc/GMT",
      "original_purchase_date_ms": "1500261436000",
      "original_purchase_date_pst": "2017-07-16 20:17:16 America/Los_Angeles"
    },
    {
      "quantity": "2",
      "product_id": "jp.aktsk.kalvados.test.iap2",
      "transaction_id": "220000368932558",
      "original_transaction_id": "220000348788557",
      "web_order_line_item_id": 220000075821143,
      "is_trial_period": "false",
      "purchase_date": "2017-09-24 03:17:15 Etc/GMT",
      "purchase_date_ms": "1506223035000",
      "purchase_date_pst": "2017-09-23 20:17:15 America/Los_Angeles",
      "original_purchase_date": "2017-07-17 03:17:16 Etc/GMT",
      "original_purchase_date_ms": "1500261436000",
      "original_purchase_date_pst": "2017-07-16 20:17:16 America/Los_Angeles"
    }
  ],
  "receipt_creation_date": "2018-02-10 17:37:00 Etc/GMT",
  "receipt_creation_date_ms": "1518284220000",
  "receipt_creation_date_pst": "2018-02-10 09:37:00 America/Los_Angeles",
  "request_date": "2018-03-26 12:00:27 Etc/GMT",
  "request_date_ms": "1522065627000",
  "request_date_pst": "2018-03-26 05:00:27 America/Los_Angeles",
  "original_purchase_date": "2017-07-07 15:36:07 Etc/GMT",
  "original_purchase_date_ms": "1499441767000",
  "original_purchase_date_pst": "2017-07-07 08:36:07 America/Los_Angeles"
}`
