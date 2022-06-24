package svc

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"l4d2/service/l4d2/internal/config"
	"strings"
)

type LocalPlugin struct {
	FileName    string `json:"fileName"`
	Description string `json:"description"`
}

type PluginManager struct {
	Disabled []LocalPlugin
	Enabled  []LocalPlugin
}

func (receiver *PluginManager) InitFromPluginDirectory(c *config.Config) error {
	var err error
	disabledFiles, err := readDirPlugins(c.Path.PluginPath + "/disabled")
	if err != nil {
		return err
	}
	enabledFiles, err := readDirPlugins(c.Path.PluginPath)
	if err != nil {
		return err
	}
	file, err := ioutil.ReadFile(c.Path.PluginDescriptionPath)
	if err != nil {
		return err
	}

	var pluginDescriptions []LocalPlugin
	err = yaml.Unmarshal(file, &pluginDescriptions)
	if err != nil {
		panic(err)
	}

	return nil
}

func readDirPlugins(pathPrefix string) ([]string, error) {
	files, err := ioutil.ReadDir(pathPrefix)
	if err != nil {
		return nil, err
	}
	var info []string
	for _, fi := range files {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), ".smx") {
			info = append(info, fi.Name())
		}
	}
	return info, nil
}
