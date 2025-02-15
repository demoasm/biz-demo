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

// Code generated by hertz generator.

package douyinapi

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/mw"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
	"go.opentelemetry.io/otel/trace"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.ServiceErr.ErrCode,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		// use requestid mw
		requestid.New(
			requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
				traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
				return traceID
			}),
		),
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression),
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getfavoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MyMiddlewareFunc(),
	}
}

func _getfeedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _message_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messagechatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action2Mw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _publishvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getpublishlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relation_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...

	return nil
}

func _list3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list4Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registuserMw() []app.HandlerFunc {
	// your code...
	return nil
}
