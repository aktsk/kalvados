# Kalvados [![Build Status](https://travis-ci.org/aktsk/kalvados.svg?branch=master)](https://travis-ci.org/aktsk/kalvados)

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

### Deploy kalvados server to Google App Engine

You can run kalvados server on Google App Engine.

```
cd appengine/app
make deploy
```

Before deploy, you should put your private key file as `key.pem` and certificate file as `cert.pem` under appengine/app directory. Or you can set your private key and certificate in app.yaml like this.

In app.yaml:

```yaml
env_variables:
  # It seems GAE/Go could not handle environment variables
  # that has return code.So use ">-" to replace return code
  # with white space
  CERTIFICATE: >-
    -----BEGIN CERTIFICATE-----
    MIIB2zCCAUSgAwIBAgIEOZu39TANBgkqhkiG9w0BAQsFADApMRAwDgYDVQQKEwdB
    Y21lIENvMRUwEwYDVQQDEwxFZGRhcmQgU3RhcmswIBcNMTgwMzE5MTY1MzA1WhgP
    MjM4MzAzMTkxNjUzMDVaMCUxEDAOBgNVBAoTB0FjbWUgQ28xETAPBgNVBAMTCEpv
    biBTbm93MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJxF2I+OtGLk/7yY6+
    IcXv9XI9cBg30QOPjdZt6hP9MEZJkIk4LIKnxFbm4GBQ0Zf3MCovoL7lp7h6DSAN
    mj7QRy6XZqAkW3D+qF6bGRNiw/3PwUw0HpuvGkbGY4d8VmMG0Jia9iF/B4f1fRIy
    39k3ILBXDZ66TE9dryFxkgLIxwIDAQABoxIwEDAOBgNVHQ8BAf8EBAMCBaAwDQYJ
    KoZIhvcNAQELBQADgYEASHF7Wl2kRj294uM6WahMjklLj0kHRX9ZQ2xbezKf4P/Z
    o7d2zZ6xiB44wfoK/uEGfjL59Qe17mkOVamXMFMAWmVgtZzOzGkUzn45H7vmfQX+
    HA/9anzcllC0cswK7g60a7cULdcVxgsaI2q3mGL4UitneeO+BtSSSg5fLcPHZW8=
    -----END CERTIFICATE-----

includes:
  - secret.yaml
```

In secret.yaml **(Do not commit and push this file to SCM)**:

```yaml
env_variables:
  # It seems GAE/Go could not handle environment variables
  # that has return code.So use ">-" to replace return code
  # with white space
  PRIVATE_KEY: >-
    -----BEGIN PRIVATE KEY-----
    MIICXAIBAAKB.....
    .....
    .....
    .....
    .....
    .....
    -----END PRIVATE KEY-----
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
