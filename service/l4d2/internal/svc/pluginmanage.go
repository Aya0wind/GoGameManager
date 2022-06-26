package svc

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"l4d2/service/l4d2/internal/config"
	"strings"
)

type LocalPlugin struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type PluginManager struct {
	Disabled []LocalPlugin `json:"disabled"`
	Enabled  []LocalPlugin `json:"enabled"`
}

func InitFromPluginDirectory(c *config.Config) (manager PluginManager, err error) {
	disabledFiles, err := readDirPlugins(c.Path.PluginPath + "/disabled")
	if err != nil {
		return
	}
	enabledFiles, err := readDirPlugins(c.Path.PluginPath)
	if err != nil {
		return
	}
	fileBytes, err := ioutil.ReadFile(c.Path.PluginDescriptionPath)
	if err != nil {
		return
	}
	var localPlugins map[string]LocalPlugin
	err = yaml.Unmarshal(fileBytes, &localPlugins)
	if err != nil {
		return
	}
	for _, file := range disabledFiles {
		name := strings.TrimSuffix(file, ".smx")
		if localPlugin, exist := localPlugins[name]; exist {
			manager.Disabled = append(manager.Disabled, LocalPlugin{
				Description: localPlugin.Description,
				Name:        name,
			})
		} else {
			manager.Disabled = append(manager.Disabled, LocalPlugin{
				Description: "暂无说明",
				Name:        name,
			})
		}
	}
	for _, file := range enabledFiles {
		name := strings.TrimSuffix(file, ".smx")
		if localPlugin, exist := localPlugins[name]; exist {
			manager.Enabled = append(manager.Enabled, LocalPlugin{
				Description: localPlugin.Description,
				Name:        name,
			})
		} else {
			manager.Enabled = append(manager.Enabled, LocalPlugin{
				Description: "暂无说明",
				Name:        name,
			})
		}
	}
	return
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
