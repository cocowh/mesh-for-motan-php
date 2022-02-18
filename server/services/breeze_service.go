package services

import (
	"mesh-for-motan-php/breeze/test_msg"
)

type BreezeService struct {}

func (s *BreezeService) TransObject(msg *test_msg.TestMsg) *test_msg.TestSubMsg {
	return msg.SubMsg
}
