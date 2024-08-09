package servers

import (
	"context"
	pbe "iroha/proto/encrypt"
	"iroha/services"
)

type Encrypt struct {
	pbe.UnimplementedEncryptServer
}

func (e *Encrypt) Encrypt(ctx context.Context, req *pbe.EncryptRequest) (res *pbe.EncryptResponse, err error) {
	// TODO Parameter validation, such as IP, client ID, etc.
	cryptographic := services.CryptographicService{}
	cryptographic.ValidateParams(req)
	// TODO Encryption processing
	cryptographic.EncryptProcess(req)
	return nil, nil
}
