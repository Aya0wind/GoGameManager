package admin

import (
	"context"
	"fmt"
	"l4d2/common"
	"os"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnablePluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnablePluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnablePluginLogic {
	return EnablePluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnablePluginLogic) EnablePlugin(req types.EnablePluginRequest) (resp *types.EnablePluginResponse, err error) {
	var pluginNames []string
	enablePath := l.svcCtx.Config.Path.PluginPath
	disablePath := l.svcCtx.Config.Path.PluginPath + "/disabled"
	for _, plugin := range l.svcCtx.Plugins.Disabled {
		pluginNames = append(pluginNames, plugin.Name)
	}
	needEnable := common.Intersect(pluginNames, req.PluginNames)
	if needEnable == nil {
		l.Logger.Error(err)
		resp = &types.EnablePluginResponse{
			Code: 500,
			Msg:  "enable plugin failed,not exist",
		}
		return
	}

	for _, pluginName := range needEnable {
		enabledPluginPath := fmt.Sprintf("%s/%s.smx", enablePath, pluginName)
		disablePluginPath := fmt.Sprintf("%s/%s.smx", disablePath, pluginName)
		_ = os.Rename(disablePluginPath, enabledPluginPath)

		//move plugin file from slice Disabled to slice Enabled
		for i, plugin := range l.svcCtx.Plugins.Disabled {
			if plugin.Name == pluginName {
				l.svcCtx.Plugins.Enabled = append(l.svcCtx.Plugins.Enabled, plugin)
				l.svcCtx.Plugins.Disabled = append(l.svcCtx.Plugins.Disabled[:i], l.svcCtx.Plugins.Disabled[i+1:]...)
				break
			}
		}
	}
	resp = &types.EnablePluginResponse{
		Code: 200,
		Msg:  "ok",
	}
	return
}
