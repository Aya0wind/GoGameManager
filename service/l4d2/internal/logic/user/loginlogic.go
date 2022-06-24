package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		resp = &types.LoginReply{
			Code: http.StatusBadRequest,
			Msg:  "username or password is empty",
		}
		return
	}
	user, err := l.svcCtx.Db.QueryUserByUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		err = nil
		resp = &types.LoginReply{
			Code: http.StatusBadRequest,
			Msg:  "username or password is wrong",
		}
		return
	}
	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.AdminAuth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.AdminAuth.AccessSecret, now, l.svcCtx.Config.AdminAuth.AccessExpire, user.Id)
	if err != nil {
		err = nil
		resp = &types.LoginReply{
			Code: http.StatusBadRequest,
			Msg:  "token generate failed",
		}
		return
	}
	// ---end---
	resp = &types.LoginReply{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: &types.LoginReplyData{
			Id:           user.Id,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}
	return
}
