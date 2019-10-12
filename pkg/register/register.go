package register

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"time"
)

const (
	CertificateInfo = "fanli.wpp.org"
	RsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDmH7ux8WHKeUCb8APXTLCHhet
/e4EG/j0DBAgx58MxviWJN0YnacsYB2R7DCQ4xUWIjyV2aYd+1bD/ji02FwFi1v2
B1OBGXViT9CU0GIkfBca/x9vKDrVh6ypC3u7zWGtV7CM3uK+hRi6vi815vTIPfO/
7PoG9M3ClWJb5oHbtwIDAQAB
-----END PUBLIC KEY-----`
	RsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDDmH7ux8WHKeUCb8APXTLCHhet/e4EG/j0DBAgx58MxviWJN0Y
nacsYB2R7DCQ4xUWIjyV2aYd+1bD/ji02FwFi1v2B1OBGXViT9CU0GIkfBca/x9v
KDrVh6ypC3u7zWGtV7CM3uK+hRi6vi815vTIPfO/7PoG9M3ClWJb5oHbtwIDAQAB
AoGANL7F1AxpNvbUO+D40OvYCULmLdRhQBhu/RjXrI9IU8DAPnT4bm/tKelNcBFa
U2f5Qru+zMYhpsolbrr6fcIuphNJTcDTyMact3feeMHekxwi5TbfL2qOp+kugHRL
EHftpkBYW4HkWSoPZQvnz7heCiMIPdgAA98uLYYu+BXNjGECQQDql5JjDQX7ljKx
pLqRY8GyuMBPcean6mbnokXoCEfoQ6ANVNQVqUDWxj/hBHCDbUDFD7AXDfESN3L+
AGBFV0ORAkEA1XHosB3XOTLWfs+ztcTZrBfzQ5jBy2t4M8zFkc66fy6e3li2BAXV
1UAypmZiP5DSTXkESYvniycYs3Z2G0/2xwJACASmPDx1t+OqV+gJeG6wcCtgZ1a9
S3/3hHNHcGbYDlhBYDNGDHd8f9rG1CoSrmtNi2691gvj8XtzsrrQj44sAQJBALfI
4yByMVVw7rw2P3ktzHegD7iOmZ98I/4GPb/0jyTfka/GFsOT+rEqG/KnicVN/6bx
or1pF6/7tAsi30NZMRUCQBuuWAdQ7kpe5fFaOOUU2IK+fr7BV++fktKvEU/HkydE
TpI/w3pUjWs3frTwP6/8IPMFVGJ8AB+ZQlCK4oV9Nzk=
-----END RSA PRIVATE KEY-----`
)

func Regist(code string) {

}

func GenerateActivationCode(mainVersion string) string {
	// 三十天授权
	ts := time.Now().Add(30 * 24 * time.Hour).Unix()
	ac := &ActivationCode{
		ExpireTimestamp: ts,
		CertificateInfo: CertificateInfo,
		MainVersion: mainVersion,
	}
	acStr, _ := json.Marshal(ac)
	encodeData, _ := RsaEncrypt(acStr)
	base64Str := base64.StdEncoding.EncodeToString(encodeData)
	return base64Str
}

func ValidateActivationCode(code string) (*ActivationCode, bool) {
	decodeBase64, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, false
	}
	decodeData, err := RsaDecrypt(decodeBase64)
	if err != nil {
		return nil, false
	}
	ac := &ActivationCode{}
	if err := json.Unmarshal(decodeData, ac); err != nil {
		return nil, false
	}
	return ac, true
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(RsaPublicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(RsaPrivateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
