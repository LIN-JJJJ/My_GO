package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// pkcs1
func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse key error")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}

const pri_key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAx6M+3R5bDOUk/fVIKUEMRmXPbjF/R3B/H2PG7gX6zF+LMas5
xRxU4FgJocRYhtIFwTl7aEeOB2nSr2hc/KBF+KK7ZLPjxoTI4g3dtKiH6H7ZHVj4
9nZmBkXE8PAoU0U8tWofSstKiOQuU9H981rXNPN16rIlO/CTH64Q6T2bcFYaASoM
//oXpY0o8iJGxhpG5pkJKBlNKMGqFRIne44aIRYMzDL6Xp7lKUIg3qrm7GuRo8Fj
GqbWgbR+1dIa2PjXJWH2iVtlNOOc0guNkwljGiTdl6FmUdJuwTNqpyfOZH/gvoRF
wX2B5ZsTy34wUtlTKfm73uvdF9KLRZCMfFvFVwIDAQABAoIBAQC5DwOEF9KRNoy+
+bOFwm2tiEzruehhgc/leil6lYJyFxNB3JZ1uJiZSiLLmOXzPBbnkfBqrwHir18F
CLpB2BRksf0CEZtQd+B7ZB0jjaJGi+eZQ9OtK/3RTLWOHyFSZ3kqNhCB1cbm9JZZ
TuyYvJmH1TN2duL7GDxZfpLgkyJW02ZCT7/++02VjGqa1zcuCwh0bmWtxr2ujgc2
hj41xCGv2faEaF7zqFIhmFSjgFfkcIFcx80bclPtlY1+PtbUr9dnEWtB6CcaVx4I
pwT1MYWcYuZ38kz/E1DSEve42cS6yfl2lafcLy++2SydCDsMNo94dR7Hhw1U0Blf
X77dXs9BAoGBAOyYAmUH3mEanuz0ubA12POHc6iFRmGf4jJoSqQsamzOjNwQ23Ag
0zfmdey3KmT1SlTzinbvd707OGshfkxhc/9TcvcpSY18nLw48wztdapvwPgMwRWX
6mgjfb9+1b2PwD3zVA7S6+MPEzaGCmJHJ5QizQhvKX361ZylGw0H9k0hAoGBANgD
O5QIw767ISNLkscgu5pph5yqtenoJqMsQ/OAJVXd4iP3XZO4+/k0YCy0baLVuXHc
1LNgaZP1t/cEUaCp3TEXpkOOHzagv+ZM6Xfm6XEYwlXh0IwNa7cvqfi6aCGxXH5w
FkjRBDKCtBiPcOntZdrTwyQjr5/YB5JOYkgIXYt3AoGBAMh+NKFLHr8pIP5qkKTr
rNMKuQ9ZCYCXTccrq/0eCn30N1gSDPMjTfq0GMClo450cy40R+VsHwtEERZwBqhR
eUwoemdLHKKCtAupMwaEgE8TbvKFVGapGyJu2RQbNqPyGpYlCtmZEf6TetOcmVYZ
OEgHibqZAQ9aLgUVwSu00JshAoGAJmfMu/Ei/FJA/gl9uzGyqS2CEvS3CzNfSzuf
iTeLa8zbXBGq7YzCH/iT8N7Tb1QYeIoOtyW5H7lcT+rQqIQK86OEyBYIrqm99LBQ
AiVn89e3FZXkgkIQmK08xyA9S2BBVamQDLo8yM48PvGh3bV7moBz/YTqwLyj+/s7
thrmptkCgYEAy1bcuL7wV8wHUZ4x6Yy9hBKGqcb9am8GEXyQw3EAfW84M/0qAqm2
znIk+8TBHvhSNe2rp1Pc+H0PMU3VwWKTWrJn/pt90HVFwvciK/MkMgc3jC6PVVCs
dqUbw+MT5dTVApaquhGfxKytAtrQjt/w38LfI1cRYx23E3DcUCXJEEg=
-----END RSA PRIVATE KEY-----
`

const pub_key = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAx6M+3R5bDOUk/fVIKUEMRmXPbjF/R3B/H2PG7gX6zF+LMas5xRxU
4FgJocRYhtIFwTl7aEeOB2nSr2hc/KBF+KK7ZLPjxoTI4g3dtKiH6H7ZHVj49nZm
BkXE8PAoU0U8tWofSstKiOQuU9H981rXNPN16rIlO/CTH64Q6T2bcFYaASoM//oX
pY0o8iJGxhpG5pkJKBlNKMGqFRIne44aIRYMzDL6Xp7lKUIg3qrm7GuRo8FjGqbW
gbR+1dIa2PjXJWH2iVtlNOOc0guNkwljGiTdl6FmUdJuwTNqpyfOZH/gvoRFwX2B
5ZsTy34wUtlTKfm73uvdF9KLRZCMfFvFVwIDAQAB
-----END RSA PUBLIC KEY-----
`

func generateTokenUsingRS256() (string, error) {
	claim := MyCustomClaims{
		UserID:     000001,
		Username:   "Tom",
		GrantScope: "read_user_info",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   // 签发者
			Subject:   "Tom",                                           // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // jwt ID, 类似于盐值
		},
	}
	rsa_pri_key, err := parsePriKeyBytes([]byte(pri_key))                                    // 读取私钥
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claim).SignedString(rsa_pri_key) // 私钥签名
	return token, err
}

func parsePubKeyBytes(pub_key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub_key)
	if block == nil {
		return nil, errors.New("block nil")
	}
	pub_ret, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("x509.ParsePKCS1PublicKey error")
	}

	return pub_ret, nil
}

func parseTokenRs256(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) { // 验签
		pub, err := parsePubKeyBytes([]byte(pub_key)) // 读取公钥
		if err != nil {
			fmt.Println("err = ", err)
			return nil, err
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
