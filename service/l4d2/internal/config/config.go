package config

import (
	"github.com/zeromicro/go-zero/rest"
	"l4d2/common"
)

type DataBaseConfig struct {
	DataBaseUrl string `json:"dataBaseUrl"`
}

type PathConfig struct {
	VpkPath               string `json:"vpkPath"`
	PluginPath            string `json:"pluginPath"`
	CommandTemplatePath   string `json:"commandTemplatePath"`
	PluginDescriptionPath string `json:"pluginDescriptionPath"`
}

type Config struct {
	rest.RestConf
	Rcon      common.RconConfig `json:"rcon"`
	DataBase  DataBaseConfig    `json:"database"`
	Path      PathConfig        `json:"path"`
	AdminAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
