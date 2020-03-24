package grpc_test

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/trueone/beetest/internal/app/secret"
	"github.com/trueone/beetest/internal/app/secret/presenter/http/grpc/proto"
)

func TestGRPC(t *testing.T) {
	// GRPC server - start
	lis, err := net.Listen("tcp", ":8081")
	assert.NoError(t, err)

	server := grpc.NewServer()

	secretRegistry := secret.New()
	secretRegistry.RegisterGRPCHandlers(server)

	go func() {
		server.Serve(lis)
	}()
	// GRPC server - end

	// GRPC client - start
	grpcConn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	assert.NoError(t, err)
	defer grpcConn.Close()

	secretClient := proto.NewSecretClient(grpcConn)
	ctx := context.Background()
	// GRPC client - end

	// Create test cases - start
	createCases := []struct {
		valid bool
		name  string
		data  *proto.SecretRequest
	}{
		{
			valid: true,
			name:  "OK",
			data: &proto.SecretRequest{
				Secret:           "secret",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
		{
			valid: false,
			name:  "Failed",
			data: &proto.SecretRequest{
				Secret:           "secret",
				ExpireAfterViews: 0,
				ExpireAfter:      0,
			},
		},
		{
			valid: false,
			name:  "Failed",
			data: &proto.SecretRequest{
				Secret:           "",
				ExpireAfterViews: 1,
				ExpireAfter:      0,
			},
		},
	}

	for _, c := range createCases {
		_, err = secretClient.Create(ctx, c.data)
		if c.valid {
			assert.NoError(t, err, c.name)
		} else {
			assert.Error(t, err, c.name)
		}
	}
	// Create test cases - end

	// Get test cases - start
	hash := md5.Sum([]byte("secret"))
	hashString := hex.EncodeToString(hash[:])

	getCases := []struct {
		valid bool
		name  string
		data  *proto.HashRequest
	}{
		{
			valid: true,
			name:  "OK",
			data:  &proto.HashRequest{Hash: hashString},
		},
		{
			valid: false,
			name:  "View limit exceeded",
			data:  &proto.HashRequest{Hash: hashString},
		},
		{
			valid: false,
			name:  "Not found",
			data:  &proto.HashRequest{Hash: "secret_not_found"},
		},
	}

	for _, c := range getCases {
		_, err = secretClient.Get(ctx, c.data)
		if c.valid {
			assert.NoError(t, err, c.name)
		} else {
			assert.Error(t, err, c.name)
		}
	}
	// Get test cases - end
}
