package grpc

import (
	"context"

	version "github.com/hashicorp/go-version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/beneath-core/beneath-go/core/middleware"
	pb "github.com/beneath-core/beneath-go/proto"
)

func (s *gRPCServer) SendClientPing(ctx context.Context, req *pb.ClientPing) (*pb.ClientPong, error) {
	// set log payload
	middleware.SetTagsPayload(ctx, clientPingTags{
		ClientID:      req.ClientId,
		ClientVersion: req.ClientVersion,
	})

	spec := clientSpecs[req.ClientId]
	if spec.IsZero() {
		return nil, grpc.Errorf(codes.InvalidArgument, "unrecognized client ID")
	}

	v, err := version.NewVersion(req.ClientVersion)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "client version is not a valid semver")
	}

	status := ""
	if v.GreaterThanOrEqual(spec.RecommendedVersion) {
		status = "stable"
	} else if v.GreaterThanOrEqual(spec.WarningVersion) {
		status = "warning"
	} else {
		status = "deprecated"
	}

	secret := middleware.GetSecret(ctx)
	return &pb.ClientPong{
		Authenticated:      !secret.IsAnonymous(),
		Status:             status,
		RecommendedVersion: spec.RecommendedVersion.String(),
	}, nil
}
