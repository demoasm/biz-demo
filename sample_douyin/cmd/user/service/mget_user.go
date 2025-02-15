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

package service

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/user/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/user/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new CheckUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{
		ctx: ctx,
	}
}

// CreateUser create user info.
func (s *MGetUserService) GetUsers(req *douyinuser.MGetUserRequest) ([]*douyinuser.User, error) {
	users, err := db.MGetUSers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(users), nil
}
