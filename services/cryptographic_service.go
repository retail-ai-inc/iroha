package services

import (
	pb "iroha/proto/encrypt"
)

type CryptographicService struct{}

func (c *CryptographicService) ValidateParams(req *pb.EncryptRequest) error {
	return nil
}

func (c *CryptographicService) EncryptProcess(req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	return nil, nil
}

func (c *CryptographicService) DecryptProcess(req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	return nil, nil
}
