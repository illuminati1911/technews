package grpc

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
	"github.com/illuminati1911/technews/service-auth/auth/delivery/grpc/proto"
	"github.com/illuminati1911/technews/utils"
	"google.golang.org/grpc/status"
)

// AuthGRPCHandler encapsulates GRPC endpoints
type AuthGRPCHandler struct {
	as auth.Service
}

var generalError *models.TNError = models.ErrGeneralServerError

// NewAuthGRPCHandler creates GRPC endpoints for given auth service
func NewAuthGRPCHandler(as auth.Service, grpc *utils.GRPC) *AuthGRPCHandler {
	handler := &AuthGRPCHandler{as}
	proto.RegisterAuthServer(grpc.Server, handler)
	return handler
}

// Login is user login gRPC endpoint which returns JWT if success
func (a *AuthGRPCHandler) Login(ctx context.Context, user *proto.UserDetails) (*proto.JWT, error) {
	if len(user.GetUsername()) == 0 || len(user.GetPassword()) == 0 {
		err := models.ErrMissingUsernamePasswordError
		return nil, status.Error(err.GRPCCode, err.Message)
	}
	token, err := a.as.Login(user.GetUsername(), user.GetPassword())
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			return nil, status.Error(te.GRPCCode, te.Message)
		}
		return nil, status.Error(generalError.GRPCCode, generalError.Message)
	}
	return &proto.JWT{Token: token}, nil
}

// CreateUser is user creation gRPC endpoint which returns error if not successful
func (a *AuthGRPCHandler) CreateUser(ctx context.Context, user *proto.UserDetails) (*empty.Empty, error) {
	if len(user.GetUsername()) == 0 || len(user.GetPassword()) == 0 {
		err := models.ErrMissingUsernamePasswordError
		return nil, status.Error(err.GRPCCode, err.Message)
	}
	_, err := a.as.CreateUser(user.GetUsername(), user.GetPassword())
	if err != nil {
		if te, ok := err.(*models.TNError); ok {
			return nil, status.Error(te.GRPCCode, te.Message)
		}
		return nil, status.Error(generalError.GRPCCode, generalError.Message)
	}
	return &empty.Empty{}, nil
}
