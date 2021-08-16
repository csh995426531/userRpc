package logic

import (
	"context"

	"subModule/userRpc/internal/svc"
	"subModule/userRpc/userRpc"

	"github.com/tal-tech/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *userRpc.InfoRequest) (*userRpc.InfoResponse, error) {
	// todo: add your logic here and delete this line

	return &userRpc.InfoResponse{
		Id: in.Id,
		Name: "test",
	}, nil
}
