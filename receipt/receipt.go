package receipt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/aktsk/nolmandy/receipt"
	"github.com/fullsailor/pkcs7"
)

// Encode encodes JSON receipt data
func Encode(receiptJSON []byte, key *rsa.PrivateKey, cert *x509.Certificate) (string, error) {
	rcpt := receipt.Receipt{}
	json.Unmarshal(receiptJSON, &rcpt)

	payload, _ := encodeReceipt(rcpt)

	signed, _ := signReceipt(payload, key, cert)

	encodedReceipt := base64.StdEncoding.EncodeToString(signed)

	return encodedReceipt, nil
}

type attribute struct {
	Type    int
	Version int
	Value   []byte
}

func encodeReceipt(r receipt.Receipt) ([]byte, error) {
	payload := []attribute{}

	var ra attribute

	// 0: receipt_type
	receiptType, err := asn1.Marshal(r.ReceiptType)
	if err != nil {
		return nil, err
	}
	ra.Type = 0
	ra.Value = receiptType
	payload = append(payload, ra)

	// 2: bundle_id
	bundleID, err := asn1.Marshal(r.BundleID)
	if err != nil {
		return nil, err
	}
	ra.Type = 2
	ra.Value = bundleID
	payload = append(payload, ra)

	// 12: receipt_creation_date
	t := time.Time(r.CreationDate.Date)
	creationDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 12
	ra.Value = creationDate
	payload = append(payload, ra)

	// 17: in_app
	for _, inApp := range r.InApp {
		encodedInApp, err := encodeInApp(inApp)
		if err != nil {
			return nil, err
		}
		ra.Type = 17
		ra.Value = encodedInApp
		payload = append(payload, ra)
	}

	// 18: original_purchase_date
	t = time.Time(r.OriginalPurchaseDate.Date)
	originalPurchaseDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 18
	ra.Value = originalPurchaseDate
	payload = append(payload, ra)

	// 19: original_application_version
	originalApplicationVersion, err := asn1.Marshal(r.OriginalApplicationVersion)
	if err != nil {
		return nil, err
	}
	ra.Type = 19
	ra.Value = originalApplicationVersion
	payload = append(payload, ra)

	// 21: expiration_date
	t = time.Time(r.ExpirationDate)
	expirationDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 21
	ra.Value = expirationDate
	payload = append(payload, ra)

	data, err := asn1.Marshal(payload)
	return data, err
}

func encodeInApp(inApp *receipt.InApp) ([]byte, error) {
	payload := []attribute{}

	var ra attribute

	// 1701: quantity
	quantity, err := asn1.Marshal(inApp.Quantity)
	if err != nil {
		return nil, err
	}
	ra.Type = 1701
	ra.Value = quantity
	payload = append(payload, ra)

	// 1702: product_id
	productID, err := asn1.Marshal(inApp.ProductID)
	if err != nil {
		return nil, err
	}
	ra.Type = 1702
	ra.Value = productID
	payload = append(payload, ra)

	// 1703: transaction_id
	transactionID, err := asn1.Marshal(inApp.TransactionID)
	if err != nil {
		return nil, err
	}
	ra.Type = 1703
	ra.Value = transactionID
	payload = append(payload, ra)

	// 1704: purchase_date
	t := time.Time(inApp.PurchaseDate.Date)
	purchaseDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 1704
	ra.Value = purchaseDate
	payload = append(payload, ra)

	// 1705: original_transaction_id
	originalTransactionID, err := asn1.Marshal(inApp.OriginalTransactionID)
	if err != nil {
		return nil, err
	}
	ra.Type = 1705
	ra.Value = originalTransactionID
	payload = append(payload, ra)

	// 1706: original_purachase_date
	t = time.Time(inApp.OriginalPurchaseDate.Date)
	originalPurchaseDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 1706
	ra.Value = originalPurchaseDate
	payload = append(payload, ra)

	// 1708: expires_date
	t = time.Time(inApp.ExpiresDate.Date)
	expiresDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 1708
	ra.Value = expiresDate
	payload = append(payload, ra)

	// 1711: web_order_line_item_id
	webOrderLineItemID, err := asn1.Marshal(inApp.WebOrderLineItemID)
	if err != nil {
		return nil, err
	}
	ra.Type = 1711
	ra.Value = webOrderLineItemID
	payload = append(payload, ra)

	// 1712: cancellation_date
	t = time.Time(inApp.CancellationDate.Date)
	cancellationDate, err := asn1.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	ra.Type = 1712
	ra.Value = cancellationDate
	payload = append(payload, ra)

	// 1719: is_intro_price
	var isIntroPrice []byte
	if inApp.IsInIntroPrice == true {
		isIntroPrice, _ = asn1.Marshal(1)
	} else {
		isIntroPrice, _ = asn1.Marshal(0)
	}
	ra.Type = 1719
	ra.Value = isIntroPrice
	payload = append(payload, ra)

	// All of InApp
	data, err := asn1.Marshal(payload)
	return data, err
}

func signReceipt(data []byte, key *rsa.PrivateKey, cert *x509.Certificate) ([]byte, error) {
	toBeSigned, err := pkcs7.NewSignedData(data)
	if err != nil {
		return nil, err
	}

	conf := pkcs7.SignerInfoConfig{}
	toBeSigned.AddSigner(cert, key, conf)
	signed, err := toBeSigned.Finish()
	if err != nil {
		return nil, err
	}

	return signed, nil
}
