// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"

	douyinfavorite "github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest, callOptions ...callopt.Option) (r *douyinfavorite.FavoriteActionResponse, err error)
	GetList(ctx context.Context, req *douyinfavorite.GetListRequest, callOptions ...callopt.Option) (r *douyinfavorite.GetListResponse, err error)
	GetIsFavorite(ctx context.Context, req *douyinfavorite.GetIsFavoriteRequest, callOptions ...callopt.Option) (r *douyinfavorite.GetIsFavoriteResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest, callOptions ...callopt.Option) (r *douyinfavorite.FavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, req)
}

func (p *kFavoriteServiceClient) GetList(ctx context.Context, req *douyinfavorite.GetListRequest, callOptions ...callopt.Option) (r *douyinfavorite.GetListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetList(ctx, req)
}

func (p *kFavoriteServiceClient) GetIsFavorite(ctx context.Context, req *douyinfavorite.GetIsFavoriteRequest, callOptions ...callopt.Option) (r *douyinfavorite.GetIsFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetIsFavorite(ctx, req)
}
