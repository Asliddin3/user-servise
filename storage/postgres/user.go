package postgres

import (
	"fmt"

	pb "github.com/Asliddin3/user-servis/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CheckField(req *pb.CheckFieldRequest) (*pb.CheckFieldResponse,error){
	exists:=pb.CheckFieldResponse{
		Exists: false,
	}
	var num int
	err:=r.db.QueryRow(`
	select 1 from users where $1=$2`,req.Key,req.Value).Scan(&num)
	fmt.Println("error while selecting key value",err)
	if err != nil {
		return &pb.CheckFieldResponse{}, err
	}
	if num!=0{
		return &exists,nil
	}
	exists.Exists=false
	return &exists,nil
}



func (r *userRepo) CreateUser(req *pb.User) (*pb.User,error){
	userResp:=pb.User{}
	err:=r.db.QueryRow(`
	insert into users(name,username,email,password)
	values($1,$2,$3,$4) returning name,username,email,password`,req.Uuid,req.Name,req.Email,req.Password).Scan(
		&userResp.Uuid,
		&userResp.Username,
		&userResp.Name,
		&userResp.Email,
		&userResp.Password,
	)
	if err != nil {
		return &pb.User{},err
	}
	return &userResp,nil
}

