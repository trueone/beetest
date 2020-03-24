package grpc

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	"github.com/trueone/beetest/internal/app/entity"
	"github.com/trueone/beetest/internal/app/secret/data"
	"github.com/trueone/beetest/internal/app/secret/presenter"
	"github.com/trueone/beetest/internal/app/secret/presenter/http/grpc/proto"
)

type Handler struct {
	registry presenter.Registry
}

func New(registry presenter.Registry) *Handler {
	return &Handler{
		registry: registry,
	}
}

func Register(server *grpc.Server, h *Handler) {
	proto.RegisterSecretServer(server, h)
}

func (h *Handler) Create(c context.Context, in *proto.SecretRequest) (out *proto.SecretResponse, err error) {
	s, err := h.registry.Create().Execute(data.SecretRequest{
		Secret:           in.GetSecret(),
		ExpireAfterViews: int(in.GetExpireAfterViews()),
		ExpireAfter:      int(in.GetExpireAfter()),
	})
	if err != nil {
		return
	}

	out, err = h.deserialize(s)
	if err != nil {
		return
	}

	return
}

func (h *Handler) Get(c context.Context, in *proto.HashRequest) (out *proto.SecretResponse, err error) {
	hash := data.HashRequest(in.GetHash())

	s, err := h.registry.Get().Execute(hash)
	if err != nil {
		return
	}

	out, err = h.deserialize(s)
	if err != nil {
		return
	}

	return
}

func (h *Handler) deserialize(s entity.Secret) (out *proto.SecretResponse, err error) {
	out = new(proto.SecretResponse)

	created, err := ptypes.TimestampProto(s.Created)
	if err != nil {
		return
	}

	expired, err := ptypes.TimestampProto(s.Created.Add(time.Minute * time.Duration(s.ExpireAfter)))
	if err != nil {
		return
	}

	out.Hash = s.Hash
	out.SecretText = s.Text
	out.CreatedAt = created
	if s.ExpireAfter != 0 {
		out.ExpiresAt = expired
	}
	out.RemainingViews = int64(s.RemainingViews)

	return
}
