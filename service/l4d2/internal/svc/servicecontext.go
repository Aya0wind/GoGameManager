package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
	"l4d2/service/l4d2/internal/config"
	"l4d2/service/l4d2/internal/types"
	"l4d2/service/l4d2/model"
	"os"
)

type ServiceContext struct {
	Config    config.Config
	Db        model.L4D2Model
	Plugins   PluginManager
	Templates []types.CommandTemplate
}

func NewServiceContext(c config.Config) *ServiceContext {
	Db, err := gorm.Open("mysql", c.DataBase.DataBaseUrl)
	Db.SingularTable(true)
	if err != nil {
		logx.Errorf("database error:%s", err.Error())
		panic(err)
	}

	file, err := os.ReadFile(c.Path.CommandTemplatePath)
	if err != nil {
		logx.Errorf("read file error:%s", err.Error())
		panic(err)
	}
	templates := make([]types.CommandTemplate, 0)
	err = yaml.Unmarshal(file, &templates)
	if err != nil {
		logx.Errorf("unmarshal file error:%s", err.Error())
		panic(err)
	}
	plugins, err := InitFromPluginDirectory(&c)
	if err != nil {
		logx.Errorf("read file error:%s", err.Error())
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Db: model.L4D2Model{
			Db: Db,
		},
		Templates: templates,
		Plugins:   plugins,
	}
}
