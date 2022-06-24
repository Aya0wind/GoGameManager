package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"l4d2/service/l4d2/internal/config"
	"l4d2/service/l4d2/internal/types"
	"l4d2/service/l4d2/model"
)

type ServiceContext struct {
	Config       config.Config
	Db           model.L4D2Model
	LocalPlugins PluginManager
	Templates    []types.CommandTemplate
}

func NewServiceContext(c config.Config) *ServiceContext {
	Db, err := gorm.Open("mysql", c.DataBase.DataBaseUrl)
	Db.SingularTable(true)
	//Db.CreateTable(&model.MapFile{})
	//Db.CreateTable(&model.MapGroup{})
	//Db.Create(&model.User{
	//	Username: "admin",
	//	Password: "admin",
	//})
	if err != nil {
		logx.Errorf("database error:%s", err.Error())
		panic(err)
	}

	file, err := ioutil.ReadFile(c.Path.CommandTemplatePath)
	if err != nil {
		panic(err)
	}
	var templates []types.CommandTemplate
	err = yaml.Unmarshal(file, &templates)
	if err != nil {
		panic(err)
	}

	PluginManager := PluginManager{}

	PluginManager.readDirPlugins(c.Path.PluginPath)

	return &ServiceContext{
		Config: c,
		Db: model.L4D2Model{
			Db: Db,
		},
		Templates: templates,
	}
}
