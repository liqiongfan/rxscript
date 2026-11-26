//
// The RXScript author
// built-in functions

package ast

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

type GenerateRSAKey struct {
	FuncDef
}

func (p *GenerateRSAKey) String() string {
	return "'generateRSAKey'"
}

func NewGenerateRSAKey() *GenerateRSAKey {
	return &GenerateRSAKey{}
}

// Execute generateRSAKey(2048, "public.pem", "private.pem")
func (p *GenerateRSAKey) Execute(rxe RxeHandler) any {
	if len(p.Args) != 3 {
		return ThrowSystemError("RXScript: generateRSAKey need three one string argument.")
	}

	bits := p.Args[0]
	publicPath := p.Args[1]
	privatePath := p.Args[2]

	switch bl := bits.(type) {
	case *IntDef:
		switch pubPath := publicPath.(type) {
		case *StringDef:
			switch privPath := privatePath.(type) {
			case *StringDef:
				privateKey, err := rsa.GenerateKey(rand.Reader, int(bl.Value))
				if err != nil {
					return ThrowSystemError(err)
				}
				publicKey := &privateKey.PublicKey

				pubWriter, err := os.OpenFile(pubPath.Value, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
				if err != nil {
					return ThrowSystemError(err)
				}
				defer pubWriter.Close()

				privWriter, err := os.OpenFile(privPath.Value, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
				if err != nil {
					return ThrowSystemError(err)
				}
				defer privWriter.Close()

				err = pem.Encode(privWriter, &pem.Block{
					Type:  "RSA PRIVATE KEY",
					Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
				})
				if err != nil {
					return ThrowSystemError(err)
				}

				pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
				if err != nil {
					return ThrowSystemError(err)
				}
				err = pem.Encode(pubWriter, &pem.Block{
					Type:  "RSA PUBLIC KEY",
					Bytes: pubBytes,
				})
				if err != nil {
					return ThrowSystemError(err)
				}

				return &BoolDef{true}
			}
		}
	}

	return &BoolDef{Value: false}
}
