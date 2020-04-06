package grpc

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
	"github.com/illuminati1911/technews/service-auth/auth/delivery/grpc/proto"
	"github.com/illuminati1911/technews/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGRPCHandler struct {
	as auth.Service
}

// NewAuthGRPCHandler creates GRPC endpoints for given auth service
func NewAuthGRPCHandler(as auth.Service, grpc *utils.GRPC) *AuthGRPCHandler {
	handler := &AuthGRPCHandler{as}
	proto.RegisterAuthServer(grpc.Server, handler)
	return handler
}

func (a *AuthGRPCHandler) Login(ctx context.Context, user *proto.UserDetails) (*proto.JWT, error) {
	if len(user.GetUsername()) == 0 || len(user.GetPassword()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Missing username and/or password")
	}
	token, err := a.as.Login(user.GetUsername(), user.GetPassword())
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			return nil, status.Error(codes.Unknown, te.Message)
		}
		return nil, status.Error(codes.Unknown, models.ErrGeneralServerError.Message)
	}
	return &proto.JWT{Token: token}, nil
}

func (a *AuthGRPCHandler) CreateUser(ctx context.Context, user *proto.UserDetails) (*empty.Empty, error) {
	if len(user.GetUsername()) == 0 || len(user.GetPassword()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Missing username and/or password")
	}
	_, err := a.as.CreateUser(user.GetUsername(), user.GetPassword())
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			return nil, status.Error(codes.Unknown, te.Message)
		}
		return nil, status.Error(codes.Unknown, models.ErrGeneralServerError.Message)
	}
	return &empty.Empty{}, nil
}

func (a *AuthGRPCHandler) Test(ctx context.Context, user *proto.UserDetails) (*proto.UserDetails, error) {
	return &proto.UserDetails{Username: user.GetUsername(), Password: user.GetPassword()}, nil
}