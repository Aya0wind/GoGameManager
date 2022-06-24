package impl

import (
	"fmt"
	"github.com/james4k/rcon"
	"l4d2/service/l4d2/internal/config"
)

func ExecRconCommand(config *config.RconConfig, cmd string) (string, error) {
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
