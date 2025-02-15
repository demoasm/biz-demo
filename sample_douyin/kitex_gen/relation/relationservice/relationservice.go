// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"

	relation "github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateRelation":       kitex.NewMethodInfo(createRelationHandler, newRelationServiceCreateRelationArgs, newRelationServiceCreateRelationResult, false),
		"DeleteRelation":       kitex.NewMethodInfo(deleteRelationHandler, newRelationServiceDeleteRelationArgs, newRelationServiceDeleteRelationResult, false),
		"GetFollow":            kitex.NewMethodInfo(getFollowHandler, newRelationServiceGetFollowArgs, newRelationServiceGetFollowResult, false),
		"GetFollower":          kitex.NewMethodInfo(getFollowerHandler, newRelationServiceGetFollowerArgs, newRelationServiceGetFollowerResult, false),
		"GetFriend":            kitex.NewMethodInfo(getFriendHandler, newRelationServiceGetFriendArgs, newRelationServiceGetFriendResult, false),
		"ValidIfFollowRequest": kitex.NewMethodInfo(validIfFollowRequestHandler, newRelationServiceValidIfFollowRequestArgs, newRelationServiceValidIfFollowRequestResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceCreateRelationArgs)
	realResult := result.(*relation.RelationServiceCreateRelationResult)
	success, err := handler.(relation.RelationService).CreateRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceCreateRelationArgs() interface{} {
	return relation.NewRelationServiceCreateRelationArgs()
}

func newRelationServiceCreateRelationResult() interface{} {
	return relation.NewRelationServiceCreateRelationResult()
}

func deleteRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceDeleteRelationArgs)
	realResult := result.(*relation.RelationServiceDeleteRelationResult)
	success, err := handler.(relation.RelationService).DeleteRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceDeleteRelationArgs() interface{} {
	return relation.NewRelationServiceDeleteRelationArgs()
}

func newRelationServiceDeleteRelationResult() interface{} {
	return relation.NewRelationServiceDeleteRelationResult()
}

func getFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowArgs)
	realResult := result.(*relation.RelationServiceGetFollowResult)
	success, err := handler.(relation.RelationService).GetFollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceGetFollowArgs() interface{} {
	return relation.NewRelationServiceGetFollowArgs()
}

func newRelationServiceGetFollowResult() interface{} {
	return relation.NewRelationServiceGetFollowResult()
}

func getFollowerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowerArgs)
	realResult := result.(*relation.RelationServiceGetFollowerResult)
	success, err := handler.(relation.RelationService).GetFollower(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceGetFollowerArgs() interface{} {
	return relation.NewRelationServiceGetFollowerArgs()
}

func newRelationServiceGetFollowerResult() interface{} {
	return relation.NewRelationServiceGetFollowerResult()
}

func getFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFriendArgs)
	realResult := result.(*relation.RelationServiceGetFriendResult)
	success, err := handler.(relation.RelationService).GetFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceGetFriendArgs() interface{} {
	return relation.NewRelationServiceGetFriendArgs()
}

func newRelationServiceGetFriendResult() interface{} {
	return relation.NewRelationServiceGetFriendResult()
}

func validIfFollowRequestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceValidIfFollowRequestArgs)
	realResult := result.(*relation.RelationServiceValidIfFollowRequestResult)
	success, err := handler.(relation.RelationService).ValidIfFollowRequest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newRelationServiceValidIfFollowRequestArgs() interface{} {
	return relation.NewRelationServiceValidIfFollowRequestArgs()
}

func newRelationServiceValidIfFollowRequestResult() interface{} {
	return relation.NewRelationServiceValidIfFollowRequestResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateRelation(ctx context.Context, req *relation.CreateRelationRequest) (r *relation.BaseResp, err error) {
	var _args relation.RelationServiceCreateRelationArgs
	_args.Req = req
	var _result relation.RelationServiceCreateRelationResult
	if err = p.c.Call(ctx, "CreateRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest) (r *relation.BaseResp, err error) {
	var _args relation.RelationServiceDeleteRelationArgs
	_args.Req = req
	var _result relation.RelationServiceDeleteRelationResult
	if err = p.c.Call(ctx, "DeleteRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollow(ctx context.Context, req *relation.GetFollowListRequest) (r *relation.GetFollowListResponse, err error) {
	var _args relation.RelationServiceGetFollowArgs
	_args.Req = req
	var _result relation.RelationServiceGetFollowResult
	if err = p.c.Call(ctx, "GetFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollower(ctx context.Context, req *relation.GetFollowerListRequest) (r *relation.GetFollowerListResponse, err error) {
	var _args relation.RelationServiceGetFollowerArgs
	_args.Req = req
	var _result relation.RelationServiceGetFollowerResult
	if err = p.c.Call(ctx, "GetFollower", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriend(ctx context.Context, req *relation.GetFriendRequest) (r *relation.GetFriendResponse, err error) {
	var _args relation.RelationServiceGetFriendArgs
	_args.Req = req
	var _result relation.RelationServiceGetFriendResult
	if err = p.c.Call(ctx, "GetFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest) (r *relation.ValidIfFollowResponse, err error) {
	var _args relation.RelationServiceValidIfFollowRequestArgs
	_args.Req = req
	var _result relation.RelationServiceValidIfFollowRequestResult
	if err = p.c.Call(ctx, "ValidIfFollowRequest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
