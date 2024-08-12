package servers

import (
	"context"
	pb "iroha/proto/cryptographic"
	"iroha/services"
)

type CryptographicServer struct {
	pb.UnimplementedCryptographicServer
}

func (s *CryptographicServer) Encrypt(ctx context.Context, req *pb.CryptoRequest) (*pb.CryptoResponse, error) {
	// TODO Parameter validation, such as IP, client ID, etc.
	cryptographic := services.CryptographicService{}
	cryptographic.ValidateParams(req)
	// TODO Encryption processing
	cryptographic.EncryptProcess(req)
	return &pb.CryptoResponse{}, nil
}

func (s *CryptographicServer) Decrypt(ctx context.Context, req *pb.CryptoRequest) (*pb.CryptoResponse, error) {
	// TODO Parameter validation, such as IP, client ID, etc.
	cryptographic := services.CryptographicService{}
	cryptographic.ValidateParams(req)
	// TODO Encryption processing
	cryptographic.EncryptProcess(req)
	return &pb.CryptoResponse{}, nil
}
