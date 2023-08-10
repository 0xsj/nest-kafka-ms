package token

import (
	pb "authentication/generated"
	"context"
)

type Repository interface {
	GenerateToken(ctx context.Context, id string) (*Pair, error)
}

type RepositoryImpl struct {
	client pb.TokenServiceClient
}

func NewRepository(client pb.TokenServiceClient) Repository {
	return &RepositoryImpl{client: client}
}

func (r *RepositoryImpl) GenerateToken(ctx context.Context, id string) (*Pair, error) {
	response, err := r.client.GenerateToken(ctx, &pb.GenerateTokenRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &Pair{AccessToken: response.GetAccessToken(), RefreshToken: response.GetRefreshToken()}, err
}

// func (r * RepositoryImpl) VerifyToken() {}

// func (r * RepositoryImpl) RevokeToken() {}