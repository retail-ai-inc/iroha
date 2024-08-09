package services

import (
	pb "iroha/proto/cryptographic"
)

type CryptographicService struct{}

func (c *CryptographicService) ValidateParams(req *pb.CryptoRequest) error {
	return nil
}

func (c *CryptographicService) EncryptProcess(req *pb.CryptoRequest) (*pb.CryptoResponse, error) {
	return nil, nil
}

func (c *CryptographicService) DecryptProcess(req *pb.CryptoRequest) (*pb.CryptoResponse, error) {
	return nil, nil
}
