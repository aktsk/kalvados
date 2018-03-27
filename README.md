# Kalvados

Kalvados is an Apple receipt generator for testing.

You can use kalvados to get base64 encoded bynary receipt data from JSON.

JSON format that kalvados can parse is like this.

```json
{
  "receipt_type": "ProductionSandbox",
  "adam_id": 0,
  "app_item_id": 0,
  "bundle_id": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "application_version": "1",
  "download_id": 0,
  "version_external_identifier": 0,
  "receipt_creation_date": "2017-04-07 03:53:44 Etc/GMT",
  "receipt_creation_date_ms": "1491537224000",
  "receipt_creation_date_pst": "2017-04-06 20:53:44 America/Los_Angeles",
  "request_date": "2018-02-21 00:13:33 Etc/GMT",
  "request_date_ms": "1519172013493",
  "request_date_pst": "2018-02-20 16:13:33 America/Los_Angeles",
  "original_purchase_date": "2013-08-01 07:00:00 Etc/GMT",
  "original_purchase_date_ms": "1375340400000",
  "original_purchase_date_pst": "2013-08-01 00:00:00 America/Los_Angeles",
  "original_application_version": "1.0",
  "in_app": [
    {
      "quantity": "1",
      "product_id": "XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
      "transaction_id": "1000000288468336",
      "original_transaction_id": "1000000288468336",
      "purchase_date": "2017-04-07 03:47:41 Etc/GMT",
      "purchase_date_ms": "1491536861000",
      "purchase_date_pst": "2017-04-06 20:47:41 America/Los_Angeles",
      "original_purchase_date": "2017-04-07 03:47:41 Etc/GMT",
      "original_purchase_date_ms": "1491536861000",
      "original_purchase_date_pst": "2017-04-06 20:47:41 America/Los_Angeles",
      "is_trial_period": "false"
    }
  ]
}
```

----

## Usage

### Compile kalvados

Run make command.

```
make
```

### As a receipt generator command line tool

Run kalvados command.

```
cat receipt.json | bin/kalvados -keyFile key.pem -certFile cert.pem
```

### As a receipt generator server

Run kalvados-server command.

```
bin/kalvados-server -keyFile key.pem -certFile cert.pem
```


### As a receipt generator library

```go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aktsk/kalvados/receipt"
)

func main() {
	jsonFile, _ := os.Open("receipt.json")
	receiptJSON, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	keyFile, _ := os.Open("key.pem")
	keyPEM, _ := ioutil.ReadAll(keyFile)
	keyFile.Close()
	keyDER, _ := pem.Decode(keyPEM)
	key, _ := x509.ParsePKCS1PrivateKey(keyDER.Bytes)

	certFile, _ := os.Open("cert.pem")
	certPEM, _ := ioutil.ReadAll(certFile)
	certFile.Close()
	certDER, _ := pem.Decode(certPEM)
	cert, _ := x509.ParseCertificate(certDER.Bytes)

	rcpt, _ := receipt.Encode(receiptJSON, key, cert)

	fmt.Println(rcpt)
}
```

----

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

----

## License

See [LICENSE](LICENSE).

----

## See Also

* [Receipt Validation Programming Guide](https://developer.apple.com/library/content/releasenotes/General/ValidateAppStoreReceipt/Introduction.html)
* [aktsk/nolmandy: Apple receipt processing server/library](https://github.com/aktsk/nolmandy)
