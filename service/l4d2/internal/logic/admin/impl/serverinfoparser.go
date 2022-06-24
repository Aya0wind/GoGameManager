package impl

import (
	"l4d2/service/l4d2/internal/types"
	"regexp"
)

func ParseServerInfo(state string) types.ServerInfo {
	regex, _ := regexp.Compile(`hostname: (?P<host>.*)\s*version : (?P<version>.*)\sudp/ip\s*: (?P<listen>.*)\s\[ public (?P<public>.*)\s]\sos\s*: (?P<os>.*)\smap\s*: (?P<map>.*)\splayers : (?P<player>.*max\))`)
	infos := (*regex).FindStringSubmatch(state)
	return types.ServerInfo{
		Host:    infos[1],
		Version: infos[2],
		Listen:  infos[3],
		Public:  infos[4],
		Os:      infos[5],
		Map:     infos[6],
		Player:  infos[7],
	}
}
