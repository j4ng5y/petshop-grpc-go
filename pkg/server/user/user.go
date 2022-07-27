package user

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/peer"
	"petshop-grpc-go/internal/petstore/v1"
)

type (
	UserService struct {
		v1.UnimplementedUserServiceServer
	}
)

var db = map[int64]*v1.User{}

func (svc *UserService) logPeer(ctx context.Context, method string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		log.Info().Str("peer-address", p.Addr.String()).Str("peer-auth-type", p.AuthInfo.AuthType()).Str("rpc_method", method).Msg("incoming request")
	}
}

func (svc *UserService) logError(err error, method string) error {
	log.Error().Err(err).Str("rpc_method", method).Send()
	return err
}

func (svc *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.Response, error) {
	svc.logPeer(ctx, "CreateUser")

	if u := req.GetUser(); u != nil {
		db[u.GetId()] = u

		return &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
			Type:    "user",
			Message: "successfully created user",
		}, nil
	}
	if ua := req.GetWithArray(); ua != nil {
		for _, u := range ua.Users {
			db[u.GetId()] = u
		}

		return &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
			Type:    "user",
			Message: "successfully created all users",
		}, nil
	}
	if ul := req.GetWithList(); ul != nil {
		for _, u := range ul.Users {
			db[u.GetId()] = u
		}

		return &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
			Type:    "user",
			Message: "successfully created all users",
		}, nil
	}

	err := fmt.Errorf("invalid oneof type, cannot create user")

	return &v1.Response{
		Code:    v1.ReponseCode_ResponseCode_INVALID_INPUT,
		Type:    "user",
		Message: err.Error(),
	}, svc.logError(err, "CreateUser")
}

func (svc *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	svc.logPeer(ctx, "GetUser")

	if un := req.GetByUsername(); un != nil {
		for k, v := range db {
			if v.GetUsername() == un.GetUsername() {
				return &v1.GetUserResponse{
					GetUserResponse_OneOf: &v1.GetUserResponse_User{User: db[k]},
				}, nil
			} else {
				continue
			}
		}
		err := fmt.Errorf("no user exists with username %s", req.GetByUsername().GetUsername())

		return &v1.GetUserResponse{GetUserResponse_OneOf: &v1.GetUserResponse_Error{Error: &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_NOT_FOUND,
			Type:    "user",
			Message: err.Error(),
		}}}, svc.logError(err, "GetUser")
	}

	err := fmt.Errorf("GetUserResponse_OneOf was nil")
	return &v1.GetUserResponse{GetUserResponse_OneOf: &v1.GetUserResponse_Error{Error: &v1.Response{
		Code:    v1.ReponseCode_ResponseCode_INVALID_INPUT,
		Type:    "user",
		Message: err.Error(),
	}}}, svc.logError(err, "GetUser")
}

func (svc *UserService) UpdateUser(ctx context.Context, req *v1.User) (*v1.Response, error) {
	svc.logPeer(ctx, "UpdateUser")

	_, ok := db[req.GetId()]
	if ok {
		db[req.GetId()] = &v1.User{
			Username:    req.GetUsername(),
			FirstName:   req.GetFirstName(),
			LastName:    req.GetLastName(),
			Email:       req.GetEmail(),
			Password:    req.GetPassword(),
			PhoneNumber: req.GetPhoneNumber(),
			UserStatus:  req.GetUserStatus(),
		}

		return &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
			Type:    "user",
			Message: fmt.Sprintf("successfully updated user %d", req.Id),
		}, nil
	}

	err := fmt.Errorf("cannot find user with id %d", req.GetId())

	return &v1.Response{
		Code:    v1.ReponseCode_ResponseCode_INVALID_INPUT,
		Type:    "user",
		Message: err.Error(),
	}, svc.logError(err, "UpdateUser")
}

func (svc *UserService) DeleteUser(ctx context.Context, req *v1.User) (*v1.Response, error) {
	svc.logPeer(ctx, "DeleteUser")

	_, ok := db[req.GetId()]
	if ok {
		delete(db, req.GetId())

		return &v1.Response{
			Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
			Type:    "user",
			Message: fmt.Sprintf("successfully deleted user %d", req.Id),
		}, nil
	}

	err := fmt.Errorf("cannot find user with id %d", req.GetId())

	return &v1.Response{
		Code:    v1.ReponseCode_ResponseCode_INVALID_INPUT,
		Type:    "user",
		Message: err.Error(),
	}, svc.logError(err, "DeleteUser")
}

func (svc *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	svc.logPeer(ctx, "Login")

	return &v1.LoginResponse{
		LoginResponse_OneOf: &v1.LoginResponse_Successful{Successful: &v1.LoginResponse_LoginSuccessful{
			TokenExpires: time.Now().Add(1 * time.Hour).String(),
			RateLimit:    1000,
		}},
	}, nil
}

func (svc *UserService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.Response, error) {
	return &v1.Response{
		Code:    v1.ReponseCode_ResponseCode_SUCCESSFUL,
		Type:    "user",
		Message: fmt.Sprintf("successfully logged out"),
	}, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
