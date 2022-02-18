package services

import (
	"context"
	"mesh-for-motan-php/pb/testmsg"
)

type GrpcPBService struct {
	testmsg.GreeterServer
}

func (s *GrpcPBService) TransObject (ctx context.Context, in *testmsg.ObjectMessage) (*testmsg.ObjectMessage, error) {
	return in,nil
}
