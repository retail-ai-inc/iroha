package servers

import (
	"context"
	pbd "iroha/proto/decrypt"
)

type Decrypt struct {
	pbd.UnimplementedDecryptServer
}

func (e *Encrypt) Decrypt(ctx context.Context, req *pbd.DecryptRequest) (res *pbd.DecryptResponse, err error) {
	// TODO Parameter validation, such as IP, client ID, etc.
	// cryptographic := services.CryptographicService{}
	// cryptographic.ValidateParams(req)
	// TODO Encryption processing
	// cryptographic.EncryptProcess(req)
	return nil, nil
}
