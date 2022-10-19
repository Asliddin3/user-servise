package repo

import (
	pb "github.com/Asliddin3/user-servis/genproto"
)

type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User,error)
	CheckField(*pb.CheckFieldRequest) (*pb.CheckFieldResponse,error)
}
