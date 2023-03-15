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

type DisablePluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisablePluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) DisablePluginLogic {
	return DisablePluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisablePluginLogic) DisablePlugin(req types.DisablePluginRequest) (resp *types.DisablePluginResponse, err error) {
	//collect all Name field in Plugins to a slice
	var pluginNames []string
	for _, plugin := range l.svcCtx.Plugins.Enabled {
		pluginNames = append(pluginNames, plugin.Name)
	}

	needDisable := common.Intersect(pluginNames, req.PluginNames)
	if needDisable == nil {
		l.Logger.Error(err)
		resp = &types.DisablePluginResponse{
			Code: 500,
			Msg:  "disable plugin failed,not exist",
		}
		return
	}
	enablePath := l.svcCtx.Config.Path.PluginPath
	disablePath := l.svcCtx.Config.Path.PluginPath + "/disabled/"
	for _, pluginName := range needDisable {
		enabledPluginPath := fmt.Sprintf("%s/%s.smx", enablePath, pluginName)
		disablePluginPath := fmt.Sprintf("%s/%s.smx", disablePath, pluginName)
		_ = os.Rename(enabledPluginPath, disablePluginPath)
		//move plugin file from slice Enabled to slice Disabled
		for i, plugin := range l.svcCtx.Plugins.Enabled {
			if plugin.Name == pluginName {
				l.svcCtx.Plugins.Disabled = append(l.svcCtx.Plugins.Disabled, plugin)
				l.svcCtx.Plugins.Enabled = append(l.svcCtx.Plugins.Enabled[:i], l.svcCtx.Plugins.Enabled[i+1:]...)
				break
			}
		}
	}
	resp = &types.DisablePluginResponse{
		Code: 200,
		Msg:  "ok",
	}
	return
}
