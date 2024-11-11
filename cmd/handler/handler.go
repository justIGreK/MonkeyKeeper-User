package handler

import (
	userProto "github.com/justIGreK/MoneyKeeper-User/pkg/go/user"

	"google.golang.org/grpc"
)

type Handler struct {
	server grpc.ServiceRegistrar
	user   UserService
}

func NewHandler(grpcServer grpc.ServiceRegistrar, userSRV UserService) *Handler {
	return &Handler{server: grpcServer, user: userSRV}
}
func (h *Handler) RegisterServices() {
	h.registerUserService(h.server, h.user)
}

func (h *Handler) registerUserService(server grpc.ServiceRegistrar, user UserService) {
	userProto.RegisterUserServiceServer(server, &UserServiceServer{UserSRV: user})
}
