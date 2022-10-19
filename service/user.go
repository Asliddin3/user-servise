package service

import (
	"context"

	pb "github.com/Asliddin3/user-servis/genproto"
	l "github.com/Asliddin3/user-servis/pkg/logger"
	"github.com/Asliddin3/user-servis/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}


func (s *UserService) CheckField(ctx context.Context,req *pb.CheckFieldRequest) (*pb.CheckFieldResponse,error){
	exist,err:=s.storage.User().CheckField(req)
	if err != nil {
		s.logger.Error("error while checkfiled User",l.Any("error check filed",err))
		return &pb.CheckFieldResponse{},status.Error(codes.Internal,"something went wrong")
	}
	return exist,nil
}


func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error("error while creating User", l.Any("error creating User", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong")
	}
	return user, nil
}
