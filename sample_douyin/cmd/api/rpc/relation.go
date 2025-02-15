// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package rpc

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation/relationservice"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func CreateRelation(ctx context.Context, req *relation.CreateRelationRequest) (r *relation.BaseResp, err error) {
	return relationClient.CreateRelation(ctx, req)
}

func DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest) (r *relation.BaseResp, err error) {
	return relationClient.DeleteRelation(ctx, req)
}

func GetFollower(ctx context.Context, req *relation.GetFollowerListRequest) (r *relation.GetFollowerListResponse, err error) {
	return relationClient.GetFollower(ctx, req)
}

func GetFollow(ctx context.Context, req *relation.GetFollowListRequest) (r *relation.GetFollowListResponse, err error) {
	return relationClient.GetFollow(ctx, req)
}

func ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest) (r *relation.ValidIfFollowResponse, err error) {
	return relationClient.ValidIfFollowRequest(ctx, req)
}

func GetFriend(ctx context.Context, req *relation.GetFriendRequest) (r *relation.GetFriendResponse, err error) {
	return relationClient.GetFriend(ctx, req)
}
