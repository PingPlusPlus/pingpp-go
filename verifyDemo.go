package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	signed := `PcU0SMJhbPObiIVinNnalZOjI02koWozxLrxa3WQW3rK/n7I+EuVGuXvhsq2MIfUaNiHZDgRFYybGtKr1uuFzEXjA4PwmnDHfWgwRPdjgseoU0eke6ZqGpklBRVTbF6PUy6/vAqur4xb7h1wpdrteUpCPafzDmVPsQLicdojJ/TF9ACjQW8gTNiS6tE9gL5hxy0RJ3/okRJo6dz2pvJBWkjCrgp/r98z/LQijA1o//atZrH63+DcL/GwEOgaymqbodzusXF+g6WMJ/GTJgjdPRHvpO9UAAUKkOQqvwthJvsXIH/L1xqvy+tFpo2J0Ptwg85bowKoyy1qC5ak3sqWqw==`
	sig := `{"id":"evt_04qN8cXQvIhssduhS4hpqd9p","created":1427555016,"livemode":false,"type":"account.summary.available","data":{"object":{"acct_id":"acct_0eHSiDyzv9G09ejT","object":"account_daily_summary","acct_display_name":"xx公司","created":1425139260,"summary_from":1425052800,"summary_to":1425139199,"charges_amount":1000,"charges_count":100}},"object":"event","pending_webhooks":2,"request":null,"scope":"acct_1234567890123456","acct_id":"acct_1234567890123456"}`
	parser, err := loadPublicKey("my-server.pub")
	if err != nil {
		fmt.Errorf("could not sign request: %v", err)
	}
	aaa, _ := base64.StdEncoding.DecodeString(signed)
	errs := parser.Unsign([]byte(sig), []byte(aaa))
	if errs != nil {
		fmt.Errorf("could not sign request: %v", errs)
	} else {
		fmt.Println("success")
	}

}

func loadPublicKey(path string) (Unsigner, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parsePublicKey(data)
}

func parsePublicKey(pemBytes []byte) (Unsigner, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}
	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return newUnsignerFromKey(rawkey)
}

type Unsigner interface {
	Unsign(data []byte, sig []byte) error
}

func newUnsignerFromKey(k interface{}) (Unsigner, error) {
	var sshKey Unsigner

	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

func (r *rsaPublicKey) Unsign(data []byte, sig []byte) error {
	hash := crypto.SHA256
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, hashed, sig)
	return err
}
