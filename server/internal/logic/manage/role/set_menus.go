package role

import (
	"context"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/model/manage_role_menu"
	"server/internal/svc"
	types "server/internal/types/manage/role"
)

type SetMenus struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMenus(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenus {
	return &SetMenus{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenus) SetMenus(req *types.SetMenusRequest) (resp *types.SetMenusResponse, err error) {
	var datas []*manage_role_menu.ManageRoleMenu
	for _, v := range req.MenuIds {
		datas = append(datas, &manage_role_menu.ManageRoleMenu{
			RoleId:     int64(req.RoleId),
			MenuId:     int64(v),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
	}

	if len(datas) == 0 {
		return
	}

	if err = l.svcCtx.Model.ManageRoleMenu.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "role_id",
		Operator: condition.Equal,
		Value:    req.RoleId,
	}); err != nil {
		return
	}

	err = l.svcCtx.Model.ManageRoleMenu.BulkInsert(l.ctx, nil, datas)
	return
}