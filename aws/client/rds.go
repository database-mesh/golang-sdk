package rds

import (
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/rds"
)

type RDS interface{
	CreateRDS()
	DeleteRDS()
	OperateRDS()
	DescribeRDS()
}

type service struct {
    core *rds.RDS
}

func NewService(sess *session.Session) *service  {
    return &service{core: rds.New(sess)}
}

func (s *service)CreateRDS() {

}

func (s *service)DeleteRDS() {

}

func (s *service)OperateRDS() {

}

func (s *service)DescribeRDS() {

}


