package common

import (
	"fmt"
	"github.com/james4k/rcon"
)

type RconConfig struct {
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Password string `json:"password"`
}

func ExecRconCommand(config *RconConfig, cmd string) (string, error) {
	ip := config.Ip
	port := config.Port
	password := config.Password
	rc, err := rcon.Dial(fmt.Sprintf("%v:%v", ip, port), password)
	if err != nil {
		return "", err
	}
	_, err = rc.Write(cmd + "\n")
	if err != nil {
		return "", err
	}
	response, _, err2 := rc.Read()
	if err2 != nil {
		return "", err2
	}
	return response, nil
}
